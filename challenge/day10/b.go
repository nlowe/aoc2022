package day10

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2022/util"
	"github.com/nlowe/aoc2022/util/tilemap"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2022/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 10, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %s\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) string {
	screen := tilemap.Of[rune](40, 6)
	cycles := 0
	x := 1
	for instr := range challenge.Lines() {
		cycles++
		if x >= ((cycles-1)%40)-1 && x <= ((cycles-1)%40)+1 {
			screen.SetTile((cycles-1)%40, (cycles-1)/40, '\u2588')
		} else {
			screen.SetTile((cycles-1)%40, (cycles-1)/40, '\u2800')
		}

		parts := strings.Fields(instr)

		switch parts[0] {
		case "noop":
		case "addx":
			cycles++
			if x >= ((cycles-1)%40)-1 && x <= ((cycles-1)%40)+1 {
				screen.SetTile((cycles-1)%40, (cycles-1)/40, '\u2588')
			} else {
				screen.SetTile((cycles-1)%40, (cycles-1)/40, '\u2800')
			}
			x += util.MustAtoI(parts[1])
		}
	}

	printout := strings.Builder{}
	printout.WriteRune('\n')
	for sy := 0; sy < 6; sy++ {
		for sx := 0; sx < 40; sx++ {
			r, ok := screen.TileAt(sx, sy)
			if !ok {
				panic(fmt.Errorf("%d,%d not on screen", sx, sy))
			}

			printout.WriteRune(r)
		}

		printout.WriteRune('\n')
	}

	return printout.String()
}
