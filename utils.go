package cmcm

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func printPlain(w io.Writer, cmt ...*Comment) error {
	bw := bufio.NewWriter(w)
	for i, c := range cmt {
		var s string
		if i > 0 {
			s = "----------------------------\n"
		}
		s = s + fmt.Sprintln("ID:\t", c.ID)
		s = s + fmt.Sprintln("Author:\t", c.Author)
		s = s + fmt.Sprintln("URL:\t", c.HTMLURL)
		s = s + fmt.Sprintln("")
		s = s + fmt.Sprintln(c.Body)
		if _, err := bw.Write([]byte(s)); err != nil {
			return err
		}
	}

	return bw.Flush()
}

func printJSON(w io.Writer, cmt ...*Comment) error {
	var s []byte
	var err error
	if len(cmt) == 1 {
		s, err = json.Marshal(cmt[0])
	} else {
		s, err = json.Marshal(cmt)
	}
	if err != nil {
		return fmt.Errorf("failed to marshal: %w", err)
	}
	bw := bufio.NewWriter(w)
	if _, err := bw.Write(s); err != nil {
		return err
	}
	return bw.Flush()
}
