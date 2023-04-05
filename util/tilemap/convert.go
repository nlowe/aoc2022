package tilemap

import "aoc-go-22/util"

func ToRunes(v rune) rune {
	return v
}

func ToInts(v rune) int {
	return util.MustAtoI(string(v))
}
