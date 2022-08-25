package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/johnmanjiro13/gh-cmcm/pkg/api"
	"github.com/johnmanjiro13/gh-cmcm/pkg/comment"
)

func newListCmd() *cobra.Command {
	var (
		repository string
		output     string
		perPage    int
	)

	listCmd := &cobra.Command{
		Use:   "list <commit_sha> [flags]",
		Short: "List comments of a commit",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			outputFormat, err := ParseOutputFormat(output)
			if err != nil {
				return err
			}
			owner, repo, err := parseRepository(repository)
			if err != nil {
				return err
			}

			ctx := context.Background()
			client, err := api.NewClient(ctx, &api.Config{
				Owner: owner,
				Repo:  repo,
			})
			if err != nil {
				return err
			}
			commenter := comment.NewCommenter(client)

			sha := args[0]
			comments, err := commenter.List(ctx, sha, perPage)
			if err != nil {
				return err
			}

			if outputFormat == JSON {
				if err := printJSON(cmd.OutOrStdout(), comments...); err != nil {
					return err
				}
			} else {
				if err := printPlain(cmd.OutOrStdout(), comments...); err != nil {
					return err
				}
			}
			return nil
		},
	}

	listCmd.Flags().StringVarP(&repository, "repo", "R", "", "Select another repository using the OWNER/REPO format")
	listCmd.Flags().IntVarP(&perPage, "per-page", "", 30, "The number of results per page (max 100)")
	listCmd.Flags().StringVarP(&output, "output", "o", string(Plain), "Output format. plain or json")

	return listCmd
}
