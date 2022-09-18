package cmcm

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func newGetCmd() *cobra.Command {
	var (
		repository string
		output     string
	)

	getCmd := &cobra.Command{
		Use:   "get <comment_id> [flags]",
		Short: "Get a commit comment by id",
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
			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse arg to integer: %s", args[0])
			}

			ctx := context.Background()
			commenter, err := newCommenter(ctx, &config{
				owner: owner,
				repo:  repo,
			})
			if err != nil {
				return err
			}

			cmt, err := commenter.Get(ctx, id)
			if err != nil {
				return err
			}
			if outputFormat == JSON {
				if err := printJSON(cmd.OutOrStdout(), cmt); err != nil {
					return err
				}
			} else {
				if err := printPlain(cmd.OutOrStdout(), cmt); err != nil {
					return err
				}
			}
			return nil
		},
	}

	getCmd.Flags().StringVarP(&repository, "repo", "R", "", "Select another repository using the OWNER/REPO format")
	getCmd.Flags().StringVarP(&output, "output", "o", string(Plain), "Output format. plain or json")

	return getCmd
}
