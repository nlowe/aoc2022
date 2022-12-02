package day2

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2022/challenge"
)

const exampleInput = `A Y
B X
C Z`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(exampleInput)

	result := partA(input)

	require.Equal(t, 15, result)
}
