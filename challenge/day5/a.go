package day5

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2022/util/gmath"

	"github.com/nlowe/aoc2022/util"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2022/challenge"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 5, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %s\n", partA(challenge.FromFile()))
		},
	}
}

type instruction struct {
	from int
	to   int
	move int
}

type harbor struct {
	crates map[int]string

	instructions []instruction

	maxCrate int
}

func partA(challenge *challenge.Input) string {
	h := parse(challenge)
	h.simulate9000()

	return h.topOfStacks()
}

func parse(challenge *challenge.Input) *harbor {
	h := &harbor{
		crates: map[int]string{},
	}

	parsingInstructions := false
	for line := range challenge.Lines() {
		// If this is the line separating crate stacks and instructions, or the line identifying crate stack numbers,
		// skip it and start parsing instructions.
		if line == "" || line[1] == '1' {
			parsingInstructions = true
			continue
		}

		// First, try parsing instructions if it is time
		if parsingInstructions {
			fields := strings.Fields(line)
			h.instructions = append(h.instructions, instruction{
				from: util.MustAtoI(fields[3]),
				to:   util.MustAtoI(fields[5]),
				move: util.MustAtoI(fields[1]),
			})

			continue
		}

		// Otherwise, parse stacks
		//
		// Including the separator, crates are 4 characters wide
		// The start '[', the crate contents, the end ']', and the separator ' '
		for i := 0; i < len(line); i += 4 {
			crateID := (i / 4) + 1

			h.maxCrate = gmath.Max(h.maxCrate, crateID)

			if line[i+1] != ' ' {
				// there's a crate here
				h.crates[crateID] += string(line[i+1])
			}
		}
	}

	return h
}

func (h *harbor) simulate9000() {
	for _, instr := range h.instructions {
		for i := 0; i < instr.move; i++ {
			h.crates[instr.to], h.crates[instr.from] = string(h.crates[instr.from][0])+h.crates[instr.to], h.crates[instr.from][1:]
		}
	}
}

func (h *harbor) topOfStacks() string {
	answer := strings.Builder{}
	for i := 1; i <= h.maxCrate; i++ {
		answer.WriteRune(rune(h.crates[i][0]))
	}

	return answer.String()
}
