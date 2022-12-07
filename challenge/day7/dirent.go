package day7

import (
	"strings"

	"github.com/nlowe/aoc2022/challenge"
	"github.com/nlowe/aoc2022/util"
)

func parse(challenge *challenge.Input) *dirent {
	fs := &dirent{
		children: map[string]*dirent{},
	}

	var stack []string
	for line := range challenge.Lines() {
		parts := strings.Fields(line)

		if parts[0] == "$" {
			// We can just ignore the line for the 'ls' command
			if parts[1] != "cd" {
				continue
			}

			switch parts[2] {
			case "..":
				stack = stack[:len(stack)-1]
			case "/":
				stack = nil
			default:
				stack = append(stack, parts[2])
			}

			continue
		}

		if parts[0] == "dir" {
			continue
		}

		p := fs
		for _, path := range stack {
			if _, ok := p.children[path]; !ok {
				p.children[path] = &dirent{children: map[string]*dirent{}}
			}

			p = p.children[path]
		}

		p.files += util.MustAtoI(parts[0])
	}

	return fs
}

type dirent struct {
	children map[string]*dirent

	files int

	nestedSize int
}

func (d *dirent) Size() int {
	if d.nestedSize != 0 {
		return d.nestedSize
	}

	sum := d.files
	for _, c := range d.children {
		sum += c.Size()
	}

	d.nestedSize = sum
	return sum
}

func (d *dirent) ThresholdSum(max int) int {
	var sum int
	if sz := d.Size(); sz <= max {
		sum += sz
	}

	for _, c := range d.children {
		sum += c.ThresholdSum(max)
	}

	return sum
}

func (d *dirent) Walk(visit func(*dirent)) {
	visit(d)
	for _, c := range d.children {
		c.Walk(visit)
	}
}
