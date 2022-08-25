package cmd

import (
	"github.com/spf13/cobra"
)

func New() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:   "cmcm",
		Short: "comments to git commit",
	}

	// create
	createCmd, err := newCreateCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(createCmd)
	// update
	updateCmd, err := newUpdateCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(updateCmd)
	// list
	rootCmd.AddCommand(newListCmd())
	// get
	rootCmd.AddCommand(newGetCmd())
	// delete
	rootCmd.AddCommand(newDeleteCmd())

	return rootCmd, nil
}
