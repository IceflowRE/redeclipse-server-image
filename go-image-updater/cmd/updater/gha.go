package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/IceflowRE/redeclipse-server-image/pkg/updater"
)

type ghaCli struct {
	noGHOut    bool
	printGhOut bool

	registryUrls []string
	registryUser []string
	registryPass []string

	arch       string
	auths      []updater.Auth
	baseImage  string
	dockerfile string
	gitRef     string
	gitRepo    string
	images     []string
	os         string
}

func (cli *ghaCli) PreRunE(cmd *cobra.Command, args []string) error {
	if len(cli.registryUrls) != len(cli.registryUser) || len(cli.registryUrls) != len(cli.registryPass) {
		return fmt.Errorf("registry-url, registry-user and registry-pass must have the same length")
	}

	for idx := range len(cli.registryUrls) {
		cli.auths = append(cli.auths, updater.Auth{
			Registry: cli.registryUrls[idx],
			User:     cli.registryUser[idx],
			Pass:     cli.registryPass[idx],
		})
	}
	cli.images = args

	if cli.arch == "" {
		return fmt.Errorf("arch is required")
	}
	if cli.baseImage == "" {
		return fmt.Errorf("base-image is required")
	}
	if cli.dockerfile == "" {
		return fmt.Errorf("dockerfile is required")
	}
	if cli.gitRef == "" {
		return fmt.Errorf("git-ref is required")
	}
	if cli.gitRepo == "" {
		return fmt.Errorf("git-repository is required")
	}
	if cli.os == "" {
		return fmt.Errorf("os is required")
	}

	return nil
}

func (cli *ghaCli) RunE(cmd *cobra.Command, args []string) error {
	newFp, err := updater.GetNewFingerprint(
		cli.baseImage,
		updater.ContainerOption{
			Auths: cli.auths,
			Arch:  cli.arch,
			Os:    cli.os,
		},
		cli.dockerfile,
		updater.GitOption{
			Repository: cli.gitRepo,
			Reference:  cli.gitRef,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to get new fingerprint: %w", err)
	}

	otherDigestsMap := map[string]string{}
	isUpdateRequired := false
	for idx, img := range cli.images {
		curFp, otherDigests, err := updater.GetCurrentFingerprint(img, cli.arch, cli.os, cli.auths)
		if err != nil {
			return fmt.Errorf("failed to get current fingerprint for image %s: %w", img, err)
		}
		isUpdateRequired = isUpdateRequired || updater.IsUpdateRequired(*curFp, *newFp)
		var parts []string
		for _, digest := range otherDigests {
			parts = append(parts, fmt.Sprintf("%s@%s", img, digest))
		}
		otherDigestsMap[img] = strings.Join(parts, " ")

		if idx > 0 {
			fmt.Println()
		}
		fmt.Printf("%s - %s\n", img, cli.arch)
		fmt.Println(updater.FingerprintCmpString(*curFp, *newFp))
	}

	if !cli.noGHOut || cli.printGhOut {
		otherDigestsJson, err := json.Marshal(otherDigestsMap)
		if err != nil {
			return fmt.Errorf("failed to marshal other digests: %w", err)
		}
		values := map[string]any{
			"update-required": isUpdateRequired,
			"base-digest":     newFp.BaseDigest,
			"re-commit":       newFp.ReCommit,
			"dockerfile-hash": newFp.Dockerfile,
			"other-digests":   string(otherDigestsJson),
		}
		if cli.printGhOut {
			writeGhValue(os.Stdout, values)
		}
		if !cli.noGHOut {
			err = writeToGhOutput(values)
			if err != nil {
				return fmt.Errorf("failed to write GITHUB_OUTPUT: %w", err)
			}
		}
	}

	return nil
}

func writeGhValue(writer io.Writer, values map[string]any) {
	for key, value := range values {
		fmt.Fprintf(writer, "%s=%v\n", key, value)
	}
}

func writeToGhOutput(values map[string]any) error {
	ghOutput, err := os.OpenFile(os.Getenv("GITHUB_OUTPUT"), os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("failed to open GITHUB_OUTPUT: %w", err)
	}
	defer ghOutput.Close()

	writeGhValue(ghOutput, values)

	return nil
}
