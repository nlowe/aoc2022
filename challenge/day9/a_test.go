package day9

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2022/challenge"
)

const example = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(example)

	result := partA(input)

	require.Equal(t, 13, result)
}
