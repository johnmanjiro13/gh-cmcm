/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
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
	if err := cmd.New().Execute(); err != nil {
		return exitStatusError
	}
	return exitStatusOK
}
