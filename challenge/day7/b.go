package day7

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2022/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 7, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	fs := parse(challenge)

	smallest := fs.Size()

	need := 30000000 - (70000000 - smallest)
	fs.Walk(func(d *dirent) {
		if d.Size() >= need && d.Size() <= smallest {
			smallest = d.Size()
		}
	})

	return smallest
}
