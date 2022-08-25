package main

import (
	"os"

	"github.com/johnmanjiro13/gh-cmcm/pkg/cmd"
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
	rootCmd, err := cmd.New()
	if err != nil {
		return exitStatusError
	}
	if err := rootCmd.Execute(); err != nil {
		return exitStatusError
	}
	return exitStatusOK
}
