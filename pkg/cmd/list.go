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
		perPage    int
	)

	listCmd := &cobra.Command{
		Use:   "list <commit_sha> [flags]",
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
			comments, err := client.List(ctx, sha, perPage)
			if err != nil {
				return err
			}
			fmt.Println(comments)
			return nil
		},
	}

	listCmd.Flags().StringVarP(&repository, "repo", "R", "", "Select another repository using the OWNER/REPO format")
	listCmd.Flags().IntVarP(&perPage, "per-page", "", 30, "The number of results per page (default 30, max 100)")
	return listCmd
}
