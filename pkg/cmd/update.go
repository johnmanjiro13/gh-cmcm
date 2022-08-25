package cmd

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/johnmanjiro13/gh-cmcm/pkg/api"
	"github.com/johnmanjiro13/gh-cmcm/pkg/comment"
)

func newUpdateCmd() (*cobra.Command, error) {
	var (
		repository string
		body       string
	)

	updateCmd := &cobra.Command{
		Use:   "update <comment_id> [flags]",
		Short: "Update a commit comment",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if body == "" {
				return errors.New("body must not be blank")
			}
			owner, repo, err := parseRepository(repository)
			if err != nil {
				return err
			}
			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse arg to integer: %s", args[0])
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

			url, err := commenter.Update(ctx, id, body)
			if err != nil {
				return err
			}
			fmt.Println("Comment updated.")
			fmt.Println("URL: ", url)
			return nil
		},
	}

	updateCmd.Flags().StringVarP(&repository, "repo", "R", "", "Select another repository using the OWNER/REPO format")
	updateCmd.Flags().StringVarP(&body, "body", "b", "", "Content of the commit comment")

	if err := updateCmd.MarkFlagRequired("body"); err != nil {
		return nil, err
	}
	return updateCmd, nil
}
