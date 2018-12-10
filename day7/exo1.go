package main

import (
	"fmt"
	"sort"
)

func exo1(steps []step) {
	finished := []rune{}

	for {
		if len(finished) == len(steps) {
			break
		}

		toFinsih := []rune{}
		for _, step := range steps {
			if includes(finished, step.name) {
				continue
			}

			if len(step.parents) == 0 || finishedParents(finished, step.parents) {
				toFinsih = append(toFinsih, step.name)
			}
		}

		sort.Sort(byRune(toFinsih))

		fmt.Println(toFinsih)
		finished = append(finished, toFinsih[0])
	}
	fmt.Println(finished)
}

type byRune []rune

func (p byRune) Len() int           { return len(p) }
func (p byRune) Less(i, j int) bool { return p[i] < p[j] }
func (p byRune) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
