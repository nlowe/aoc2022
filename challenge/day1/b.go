package day1

import (
	"fmt"
	"sort"

	"github.com/nlowe/aoc2022/util"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2022/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 1, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	var topThree []int

	currentElf := 0
	for line := range challenge.Lines() {
		if line == "" {
			topThree = populatePodium(currentElf, topThree)

			currentElf = 0
			continue
		}

		v := util.MustAtoI(line)
		currentElf += v
	}

	// And once more in case there's no trailing newline
	if currentElf != 0 {
		topThree = populatePodium(currentElf, topThree)
	}

	if len(topThree) != 3 {
		panic("no solution")
	}

	return topThree[0] + topThree[1] + topThree[2]
}

func populatePodium(candidate int, best []int) []int {
	if len(best) < 3 {
		return append(best, candidate)
	}

	best = append(best, candidate)
	sort.Ints(best)

	return best[1:]
}
