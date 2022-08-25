package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"

	"github.com/johnmanjiro13/gh-cmcm/pkg/comment"
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
			owner, repo, err := parseRepository(repository)
			if err != nil {
				return err
			}
			sha := args[0]

			ctx := context.Background()
			client, err := comment.NewClient(ctx, &comment.Config{
				Owner: owner,
				Repo:  repo,
			})
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
