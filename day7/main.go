package main

type step struct {
	name     rune
	parents  []*step
	children []*step
}

func main() {
	steps := readInput()

	exo2(steps)
}
