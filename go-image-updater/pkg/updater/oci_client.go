package updater

import (
	"context"
	"errors"
	"net/http"
	"slices"
	"time"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/go-containerregistry/pkg/v1/remote/transport"
)

type Auth struct {
	Registry string
	User     string
	Pass     string
}

type OCIClient struct {
	auths []Auth
}

func NewOCIClient(auths []Auth) *OCIClient {
	return &OCIClient{
		auths: auths,
	}
}

func (client *OCIClient) addRemoteAuth(opts []remote.Option, registry string) []remote.Option {
	for _, auth := range client.auths {
		if auth.Registry == registry {
			opts = append(opts, remote.WithAuth(&authn.Basic{
				Username: auth.User,
				Password: auth.Pass,
			}))
		}
	}
	return opts
}

type ImageMetadata struct {
	Digest string
	Labels map[string]string
}

type getImageMetataOptions struct {
	GetLabels   bool
	OtherDigest bool
}

type GetImageMetadataOption func(opts *getImageMetataOptions)

func WithLabels() GetImageMetadataOption {
	return func(opts *getImageMetataOptions) {
		opts.GetLabels = true
	}
}

type errorConst string

const ErrImageNotFound errorConst = "image not found"

func (err errorConst) Error() string {
	return string(err)
}

func (client *OCIClient) GetImageMetadata(image string, arch string, os string, opts ...GetImageMetadataOption) (meta *ImageMetadata, otherDigests []string, err error) {
	metaOpts := &getImageMetataOptions{}
	for _, opt := range opts {
		opt(metaOpts)
	}

	ref, err := name.ParseReference(image)
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	remoteOpts := []remote.Option{remote.WithContext(ctx)}
	remoteOpts = client.addRemoteAuth(remoteOpts, ref.Context().RegistryStr())

	desc, err := remote.Get(ref, remoteOpts...)
	if err != nil {
		var httpErr *transport.Error
		if errors.As(err, &httpErr) && httpErr.StatusCode == http.StatusNotFound {
			return &ImageMetadata{}, []string{}, ErrImageNotFound
		}
		return nil, nil, err
	}

	switch {
	case desc.MediaType.IsIndex():
		return client.metaFromIndex(desc, arch, os, metaOpts)
	case desc.MediaType.IsImage():
		meta, err := client.metaFromImage(desc, arch, os, metaOpts)
		return meta, nil, err
	default:
		return nil, nil, errors.New("unsupported media type")
	}
}

// metaFromIndex resolves metadata from a multi-platform image index.
func (client *OCIClient) metaFromIndex(desc *remote.Descriptor, arch string, os string, opts *getImageMetataOptions) (meta *ImageMetadata, otherDigests []string, err error) {
	index, err := desc.ImageIndex()
	if err != nil {
		return nil, nil, err
	}

	manifest, err := index.IndexManifest()
	if err != nil {
		return nil, nil, err
	}

	digest, otherDigests, found := findPlatformDigest(manifest.Manifests, arch, os)
	if !found {
		return nil, []string{}, ErrImageNotFound
	}

	meta = &ImageMetadata{
		Digest: digest.String(),
	}
	if !opts.GetLabels {
		return meta, otherDigests, nil
	}

	img, err := index.Image(digest)
	if err != nil {
		return nil, nil, err
	}

	meta.Labels, err = configLabels(img)
	if err != nil {
		return nil, nil, err
	}
	return meta, otherDigests, nil
}

// metaFromImage resolves metadata from a single-platform image descriptor.
func (client *OCIClient) metaFromImage(desc *remote.Descriptor, arch string, os string, opts *getImageMetataOptions) (*ImageMetadata, error) {
	meta := &ImageMetadata{Digest: desc.Descriptor.Digest.String()}
	if !opts.GetLabels {
		return meta, nil
	}

	img, err := desc.Image()
	if err != nil {
		return nil, err
	}

	cfg, err := img.ConfigFile()
	if err != nil {
		return nil, err
	}

	if cfg.Architecture != arch || cfg.OS != os {
		return nil, ErrImageNotFound
	}

	meta.Labels = cfg.Config.Labels
	return meta, nil
}

// findPlatformDigest returns the digest for the first manifest matching arch+os and all other digests
func findPlatformDigest(manifests []v1.Descriptor, arch string, os string) (digest v1.Hash, otherDigests []string, found bool) {
	otherDigests = []string{}
	for _, manif := range manifests {
		if manif.Platform == nil || slices.Contains([]string{"", "unknown"}, manif.Platform.Architecture) || slices.Contains([]string{"", "unknown"}, manif.Platform.OS) || manif.Digest.String() == "" {
			continue
		}
		if manif.Platform.Architecture == arch && manif.Platform.OS == os {
			found = true
			digest = manif.Digest
		} else {
			otherDigests = append(otherDigests, manif.Digest.String())
		}
	}

	return digest, otherDigests, found
}

// configLabels returns the OCI config labels for the given image.
func configLabels(img v1.Image) (map[string]string, error) {
	cfg, err := img.ConfigFile()
	if err != nil {
		return nil, err
	}
	return cfg.Config.Labels, nil
}
