package day8

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2022/challenge"
)

const example = `30373
25512
65332
33549
35390`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(example)

	result := partA(input)

	require.Equal(t, 21, result)
}
