package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nlowe/aoc2022/challenge"
)

func TestB(t *testing.T) {
	for _, tt := range []struct {
		stream   string
		expected int
	}{
		{stream: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", expected: 19},
		{stream: "bvwbjplbgvbhsrlpgdmjqwftvncz", expected: 23},
		{stream: "nppdvjthqldpwncqszvftbrmjlhg", expected: 23},
		{stream: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", expected: 29},
		{stream: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", expected: 26},
	} {
		t.Run(tt.stream, func(t *testing.T) {
			input := challenge.FromLiteral(tt.stream)

			result := partB(input)

			assert.Equal(t, tt.expected, result)
		})
	}
}
