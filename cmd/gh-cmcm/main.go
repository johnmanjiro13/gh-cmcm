package main

import (
	"os"

	cmcm "github.com/johnmanjiro13/gh-cmcm"
)

type exitCode int

const (
	// exitStatusOK is status code zero
	exitStatusOK exitCode = iota
	// exitStatusError is status code non-zero
	exitStatusError
)

func main() {
	os.Exit(int(run()))
}

func run() exitCode {
	rootCmd, err := cmcm.NewRootCmd()
	if err != nil {
		return exitStatusError
	}
	if err := rootCmd.Execute(); err != nil {
		return exitStatusError
	}
	return exitStatusOK
}
