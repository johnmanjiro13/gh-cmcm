package cmcm

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func newDeleteCmd() *cobra.Command {
	var (
		repository string
	)

	deleteCmd := &cobra.Command{
		Use:   "delete <commit_id> [flags]",
		Short: "Delete a commit comment",
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

			commenter, err := newCommenter(&config{
				owner: owner,
				repo:  repo,
			})
			if err != nil {
				return err
			}

			if err := commenter.Delete(id); err != nil {
				return err
			}
			cmd.Println("Comment deleted.")
			return nil
		},
	}

	deleteCmd.Flags().StringVarP(&repository, "repo", "R", "", "Select another repository using the OWNER/REPO format")
	return deleteCmd
}
