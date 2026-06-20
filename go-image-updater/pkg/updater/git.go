package updater

import (
	"fmt"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/config"
	"github.com/go-git/go-git/v6/storage/memory"
)

func getGitHash(repo string, ref string) (string, error) {
	if isHash(ref) {
		return ref, nil
	}

	rem := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{repo},
	})

	refs, err := rem.List(&git.ListOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to list remote references: %w", err)
	}

	possibleRefs := []string{
		ref,
		"refs/heads/" + ref,
		"refs/tags/" + ref,
		"refs/remotes/origin/" + ref,
	}

	for _, curRef := range refs {
		refName := curRef.Name().String()
		for _, possible := range possibleRefs {
			if refName == possible {
				return curRef.Hash().String(), nil
			}
		}
	}

	return "", fmt.Errorf("reference %q not found", ref)
}

func isHash(str string) bool {
	return (len(str) == 40 || len(str) == 64) && isHex(str)
}

func isHex(str string) bool {
	for idx := 0; idx < len(str); idx++ {
		chr := str[idx]
		if !((chr >= '0' && chr <= '9') || (chr >= 'a' && chr <= 'f') || (chr >= 'A' && chr <= 'F')) {
			return false
		}
	}
	return true
}
