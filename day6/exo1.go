package main

import (
	"fmt"
	"sort"
)

type gridPoint struct {
	x            int
	y            int
	closestPoint point
}

type grid []gridPoint

func exo1(points []point) {
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

	fmt.Println("Bounds")
	fmt.Println(left, top, right, bottom)
	fmt.Println("-----------------")

	// build a grid

	gridPoints := grid{}

	for x := left; x < right; x++ {

		for y := top; y < bottom; y++ {

			distances := []struct {
				p        point
				distance int
			}{}

			for _, p := range points {
				d := distanceToPoint(x, y, p)
				distances = append(distances, struct {
					p        point
					distance int
				}{p, d})
			}

			sort.Sort(byDistance(distances))

			// check only on closest
			if distances[0].distance < distances[1].distance {
				gridPoints = append(gridPoints, gridPoint{
					x,
					y,
					distances[0].p,
				})
			}
		}
	}

	max := 0
	for _, p := range points {
		if p.isInfinite(left, top, right, bottom) {
			continue
		}

		count := 0
		for _, gP := range gridPoints {
			if gP.closestPoint == p {
				count++
			}
		}

		if count > max {
			max = count
		}
	}

	fmt.Println(max)
}
