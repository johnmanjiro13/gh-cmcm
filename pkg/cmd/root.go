package cmd

import (
	"github.com/spf13/cobra"
)

func New() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:   "cmcm",
		Short: "comments to git commit",
	}
	rootCmd.AddCommand(newListCmd())
	rootCmd.AddCommand(newGetCmd())
	updateCmd, err := newUpdateCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(updateCmd)

	return rootCmd, nil
}
