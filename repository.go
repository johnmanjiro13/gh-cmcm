package cmcm

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cli/go-gh"
)

func parseRepository(repository string) (owner, repo string, err error) {
	if repository == "" {
		owner, repo, err = resolveRepository()
		if err != nil {
			return "", "", err
		}
	} else {
		s := strings.Split(repository, "/")
		if len(s) != 2 {
			return "", "", errors.New("invalid repository name")
		}
		owner = s[0]
		repo = s[1]
	}
	return owner, repo, nil
}

func resolveRepository() (owner, repo string, err error) {
	args := []string{"repo", "view"}
	stdOut, _, err := gh.Exec(args...)
	if err != nil {
		return "", "", fmt.Errorf("failed to view repo: %w", err)
	}
	viewOut := strings.Split(stdOut.String(), "\n")[0]
	ownerRepo := strings.Split(strings.TrimSpace(strings.Split(viewOut, ":")[1]), "/")
	if len(ownerRepo) != 2 {
		return "", "", errors.New("failed to parse repository")
	}
	owner = ownerRepo[0]
	repo = ownerRepo[1]
	return owner, repo, nil
}
