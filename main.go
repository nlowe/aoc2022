package main

import (
	"os"

	"github.com/nlowe/aoc2022/challenge/cmd"
)

//go:generate go run ./gen
func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
