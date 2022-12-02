package day2

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2022/challenge"
)

func TestB(t *testing.T) {
	input := challenge.FromLiteral(exampleInput)

	result := partB(input)

	require.Equal(t, 12, result)
}
