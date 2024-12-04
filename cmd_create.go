package cmcm

import (
	"errors"

	"github.com/spf13/cobra"
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

			commenter, err := newCommenter(&config{
				owner: owner,
				repo:  repo,
			})
			if err != nil {
				return err
			}

			sha := args[0]
			url, err := commenter.Create(sha, body, path, position)
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
