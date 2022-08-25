package cmd

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/cli/go-gh"
	"github.com/spf13/cobra"

	"github.com/johnmanjiro13/gh-cmcm/comment"
)

func newListCmd() *cobra.Command {
	var (
		repository string
	)

	listCmd := &cobra.Command{
		Use:   "list <commit_sha>",
		Short: "List comments of a commit",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := &comment.Config{}
			if repository == "" {
				owner, repo, err := resolveRepository()
				if err != nil {
					return err
				}
				cfg.Owner = owner
				cfg.Repo = repo
			} else {
				s := strings.Split(repository, "/")
				if len(s) != 2 {
					return errors.New("invalid repository name")
				}
				cfg.Owner = s[0]
				cfg.Repo = s[1]
			}

			sha := args[0]
			ctx := context.Background()
			client, err := comment.NewClient(ctx, cfg)
			if err != nil {
				return err
			}
			cmt, err := client.List(ctx, sha)
			if err != nil {
				return err
			}
			for _, c := range cmt {
				fmt.Printf("Author: %s, Body: %s\n", c.Author, c.Body)
			}
			return nil
		},
	}

	listCmd.Flags().StringVarP(&repository, "repo", "R", "", "Select another repository using the OWNER/REPO format")
	return listCmd
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
