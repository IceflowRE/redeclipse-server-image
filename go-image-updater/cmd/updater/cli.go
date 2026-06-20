package main

import (
	"os"

	"github.com/spf13/cobra"
)

type cliCmd struct {
	rootCmd *cobra.Command
}

func newCli() *cliCmd {
	c := cliCmd{}
	c.rootCmd = &cobra.Command{}
	c.rootCmd.CompletionOptions.DisableDefaultCmd = true

	{
		cli := ghaCli{}
		cmd := &cobra.Command{
			Use:     "gha image arch os",
			Short:   "GitHub Actions run check.",
			Args:    cobra.MinimumNArgs(1),
			PreRunE: cli.PreRunE,
			RunE:    cli.RunE,
		}
		cmd.Flags().BoolVarP(&cli.noGHOut, "no-gh-out", "", false, "do not write to GITHUB_OUTPUT")
		cmd.Flags().BoolVarP(&cli.printGhOut, "print-gh-out", "", false, "print GITHUB_OUTPUT to stdout")
		cmd.Flags().StringArrayVar(&cli.registryUrls, "registry-url", []string{}, "registry url")
		cmd.Flags().StringArrayVar(&cli.registryUser, "registry-user", []string{}, "registry user")
		cmd.Flags().StringArrayVar(&cli.registryPass, "registry-pass", []string{}, "registry password")
		cmd.Flags().StringVarP(&cli.arch, "arch", "", "", "target architecture")
		cmd.Flags().StringVarP(&cli.baseImage, "base-image", "", "", "base image")
		cmd.Flags().StringVarP(&cli.dockerfile, "dockerfile", "", "Dockerfile", "path to Dockerfile")
		cmd.Flags().StringVarP(&cli.gitRef, "git-ref", "", "", "git reference")
		cmd.Flags().StringVarP(&cli.gitRepo, "git-repository", "", "", "git repository url")
		cmd.Flags().StringVarP(&cli.os, "os", "", "linux", "target os")
		cmd.MarkFlagRequired("arch")
		cmd.MarkFlagRequired("base-image")
		cmd.MarkFlagRequired("dockerfile")
		cmd.MarkFlagRequired("git-ref")
		cmd.MarkFlagRequired("git-repo")
		c.rootCmd.AddCommand(cmd)
	}
	return &c
}

func (c *cliCmd) Execute() {
	err := c.rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
