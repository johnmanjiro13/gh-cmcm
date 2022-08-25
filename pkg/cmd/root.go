package cmd

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "cmcm",
		Short: "comments to git commit",
	}
	rootCmd.AddCommand(newListCmd())
	rootCmd.AddCommand(newGetCmd())
	return rootCmd
}
