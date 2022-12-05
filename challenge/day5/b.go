package day5

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2022/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 5, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %s\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) string {
	h := parse(challenge)
	h.simulate9001()

	return h.topOfStacks()
}

func (h *harbor) simulate9001() {
	for _, instr := range h.instructions {
		h.crates[instr.to], h.crates[instr.from] = h.crates[instr.from][:instr.move]+h.crates[instr.to], h.crates[instr.from][instr.move:]
	}
}
