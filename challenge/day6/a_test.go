package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nlowe/aoc2022/challenge"
)

func TestA(t *testing.T) {
	for _, tt := range []struct {
		stream   string
		expected int
	}{
		{stream: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", expected: 7},
		{stream: "bvwbjplbgvbhsrlpgdmjqwftvncz", expected: 5},
		{stream: "nppdvjthqldpwncqszvftbrmjlhg", expected: 6},
		{stream: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", expected: 10},
		{stream: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", expected: 11},
	} {
		t.Run(tt.stream, func(t *testing.T) {
			input := challenge.FromLiteral(tt.stream)

			result := partA(input)

			assert.Equal(t, tt.expected, result)
		})
	}
}
