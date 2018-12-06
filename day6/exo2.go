package main

import (
	"fmt"
)

func exo2(points []point) {

	// find boundaries
	left := points[0].x
	right := points[0].x
	top := points[0].y
	bottom := points[0].y

	for _, p := range points {
		if p.x < left {
			left = p.x
		}

		if p.x > right {
			right = p.x
		}

		if p.y < top {
			top = p.y
		}

		if p.y > bottom {
			bottom = p.y
		}
	}

	width := right - left
	height := bottom - top

	// build a grid
	grid := make([][]int, width)

	for x := 0; x < width; x++ {

		grid[x] = make([]int, height)

		for y := 0; y < height; y++ {
			grid[x][y] = 0

			for _, p := range points {
				d := distanceToPoint(x+left, y+top, p)
				grid[x][y] += d
			}
		}
	}

	total := 0
	for _, row := range grid {
		for _, c := range row {
			if c < 10000 {
				total++
			}
		}

	}
	fmt.Println(total)
}
