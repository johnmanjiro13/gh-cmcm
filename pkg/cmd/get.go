package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/johnmanjiro13/gh-cmcm/pkg/comment"
)

func newGetCmd() *cobra.Command {
	var (
		repository string
	)

	getCmd := &cobra.Command{
		Use:   "get <comment_id>",
		Short: "Get a commit comment by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			owner, repo, err := parseRepository(repository)
			if err != nil {
				return err
			}
			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse arg to integer: %s", args[0])
			}

			ctx := context.Background()
			client, err := comment.NewClient(ctx, &comment.Config{
				Owner: owner,
				Repo:  repo,
			})
			if err != nil {
				return err
			}
			cmt, err := client.Get(ctx, id)
			if err != nil {
				return err
			}
			fmt.Printf("Author: %s, Body: %s\n", cmt.Author, cmt.Body)
			return nil
		},
	}

	getCmd.Flags().StringVarP(&repository, "repo", "R", "", "Select another repository using the OWNER/REPO format")
	return getCmd
}
