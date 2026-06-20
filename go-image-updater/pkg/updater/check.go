package updater

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

type Fingerprint struct {
	// the current digest of the image
	Digest     string `json:"digest"`
	BaseDigest string `json:"alpine"`
	Dockerfile string `json:"dockerfile"`
	ReCommit   string `json:"reCommit"`
}

func IsUpdateRequired(curFp Fingerprint, newFp Fingerprint) bool {
	return curFp != newFp
}

func FingerprintCmpString(curFp Fingerprint, newFp Fingerprint) string {
	return valueStatusString("Base Digest", curFp.BaseDigest, newFp.BaseDigest) + "\n" +
		valueStatusString("Dockerfile", curFp.Dockerfile, newFp.Dockerfile) + "\n" +
		valueStatusString("RE-commit", curFp.ReCommit, newFp.ReCommit)
}

type ContainerOption struct {
	Auths []Auth
	Arch  string
	Os    string
}

type GitOption struct {
	Repository string
	Reference  string
}

func GetNewFingerprint(baseImage string, containerOpt ContainerOption, dockerfile string, gitOpt GitOption) (fingerprint *Fingerprint, err error) {
	client := NewOCIClient(containerOpt.Auths)
	meta, _, err := client.GetImageMetadata(baseImage, containerOpt.Arch, containerOpt.Os)
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata: %w", err)
	}
	fingerprint = &Fingerprint{
		BaseDigest: meta.Digest,
	}

	fingerprint.Dockerfile, err = getFileHash(dockerfile, sha256.New())
	if err != nil {
		return nil, err
	}
	fingerprint.Dockerfile = "sha256:" + fingerprint.Dockerfile

	fingerprint.ReCommit, err = getGitHash(gitOpt.Repository, gitOpt.Reference)
	if err != nil {
		return nil, err
	}

	return fingerprint, nil
}

// othersDigests is the digests of the image with different platforms
func GetCurrentFingerprint(image string, arch string, os string, auths []Auth) (res *Fingerprint, otherDigests []string, err error) {
	client := NewOCIClient(auths)
	meta, otherDigests, err := client.GetImageMetadata(image, arch, os, WithLabels())
	if errors.Is(err, ErrImageNotFound) {
		return &Fingerprint{}, otherDigests, nil
	}
	if err != nil {
		return nil, nil, err
	}
	return &Fingerprint{
		Digest:     meta.Digest,
		BaseDigest: meta.Labels["org.opencontainers.image.base.digest"],
		Dockerfile: meta.Labels["dockerfile.hash"],
		ReCommit:   meta.Labels["org.opencontainers.image.revision"],
	}, otherDigests, nil
}

func valueStatusString(key string, oldVal string, newVal string) string {
	status := "    Status:  SAME"
	if oldVal != newVal {
		status = "    Status:  CHANGED"
	}
	return fmt.Sprintf("%s:\n    Current: %s\n    New:     %s\n%s", key, oldVal, newVal, status)
}
