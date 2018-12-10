package main

import (
	"bufio"
	"os"
)

func readInput() []step {
	stepMap := map[rune]*step{}

	fi, err := os.Open("./input.txt")
	check(err)
	defer fi.Close()

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := scanner.Text()

		name := rune(line[5])
		childName := rune(line[36])

		var s *step
		s, ok := stepMap[name]
		if !ok {
			s = &step{
				name:     name,
				parents:  []*step{},
				children: []*step{},
			}
			stepMap[name] = s
		}

		var child *step
		child, ok = stepMap[childName]
		if !ok {
			child = &step{
				name:     childName,
				parents:  []*step{},
				children: []*step{},
			}

			stepMap[childName] = child
		}

		s.children = append(s.children, child)
		child.parents = append(child.parents, s)
	}

	steps := []step{}
	for _, s := range stepMap {
		steps = append(steps, *s)
	}

	return steps
}
