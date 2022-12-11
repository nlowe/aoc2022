package day11

import (
	"fmt"
	"sort"
	"strings"

	"github.com/nlowe/aoc2022/challenge"
	"github.com/nlowe/aoc2022/util"
)

type monkey struct {
	items []int

	op   func(int) int
	test int

	pass int
	fail int

	inspections int
}

func parse(challenge *challenge.Input) []*monkey {
	sections := challenge.SectionSlice()

	monkeys := make([]*monkey, len(sections))
	for i, spec := range sections {
		m := &monkey{}

		lines := strings.Split(spec, "\n")
		rawItems := strings.Split(strings.TrimPrefix(lines[1], "  Starting items: "), ", ")
		for _, item := range rawItems {
			m.items = append(m.items, util.MustAtoI(item))
		}

		rawOp := strings.Fields(lines[2])
		switch rawOp[4] {
		case "+":
			if rawOp[5] == "old" {
				m.op = makeSelfAddOp()
			} else {
				m.op = makeAddOp(util.MustAtoI(rawOp[5]))
			}
		case "*":
			if rawOp[5] == "old" {
				m.op = makeSelfMulOp()
			} else {
				m.op = makeMulOp(util.MustAtoI(rawOp[5]))
			}
		default:
			panic(fmt.Errorf("unknown op '%s' for func: %s", rawOp[4], lines[2]))
		}

		m.test = util.MustAtoI(strings.Fields(lines[3])[3])
		m.pass = util.MustAtoI(strings.Fields(lines[4])[5])
		m.fail = util.MustAtoI(strings.Fields(lines[5])[5])

		monkeys[i] = m
	}

	return monkeys
}

func simulateAndSort(monkeys []*monkey, turns, shrink int) {
	for i := 0; i < turns; i++ {
		for _, m := range monkeys {
			var item int
			for len(m.items) > 0 {
				m.inspections++
				item, m.items = m.items[0], m.items[1:]

				item = m.op(item)
				if shrink == 0 {
					item /= 3
				} else {
					item %= shrink
				}

				if item%m.test == 0 {
					monkeys[m.pass].items = append(monkeys[m.pass].items, item)
				} else {
					monkeys[m.fail].items = append(monkeys[m.fail].items, item)
				}
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})
}

func makeSelfAddOp() func(int) int {
	return func(old int) int {
		return old + old
	}
}

func makeAddOp(off int) func(int) int {
	return func(old int) int {
		return old + off
	}
}

func makeSelfMulOp() func(int) int {
	return func(old int) int {
		return old * old
	}
}

func makeMulOp(off int) func(int) int {
	return func(old int) int {
		return old * off
	}
}
