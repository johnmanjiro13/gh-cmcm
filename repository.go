package cmcm

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/cli/safeexec"
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
	stdOut, _, err := gh(args...)
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

func gh(args ...string) (sout, eout *bytes.Buffer, err error) {
	sout = new(bytes.Buffer)
	eout = new(bytes.Buffer)
	bin, err := safeexec.LookPath("gh")
	if err != nil {
		err = fmt.Errorf("could not find gh. err: %w", err)
		return
	}

	cmd := exec.Command(bin, args...)
	cmd.Stdout = sout
	cmd.Stderr = eout

	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf("failed to run gh. err: %w, eout: %s", err, eout.String())
		return
	}
	return
}
