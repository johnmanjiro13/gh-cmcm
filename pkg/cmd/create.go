package cmd

import (
	"context"
	"errors"

	"github.com/spf13/cobra"

	"github.com/johnmanjiro13/gh-cmcm/pkg/api"
	"github.com/johnmanjiro13/gh-cmcm/pkg/comment"
)

func newCreateCmd() (*cobra.Command, error) {
	var (
		repository string
		body       string
		path       string
		position   int
	)

	createCmd := &cobra.Command{
		Use:   "create <commit_sha> [flags]",
		Short: "Create a commit comment",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if body == "" {
				return errors.New("body must not be blank")
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
			opt := &comment.CreateOption{
				Path:     path,
				Position: position,
			}
			url, err := commenter.Create(ctx, sha, body, opt)
			if err != nil {
				return err
			}
			cmd.Println("Comment created.")
			cmd.Println("URL: ", url)
			return nil
		},
	}

	createCmd.Flags().StringVarP(&repository, "repo", "R", "", "Select another repository using the OWNER/REPO format")
	createCmd.Flags().StringVarP(&body, "body", "b", "", "Content of the commit comment")
	createCmd.Flags().StringVarP(&path, "path", "", "", "Relative path of the file to comment on")
	createCmd.Flags().IntVarP(&position, "position", "", 0, "Line index in the diff to comment on")

	if err := createCmd.MarkFlagRequired("body"); err != nil {
		return nil, err
	}
	return createCmd, nil
}
