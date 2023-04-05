package example

import (
	"fmt"

	"github.com/spf13/cobra"

	"aoc-go-22/challenge"
	"aoc-go-22/util"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Example Day, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", a(challenge.FromFile()))
		},
	}
}

func a(input *challenge.Input) (result int) {
	return util.MustAtoI(<-input.Lines())
}
