package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	id int
	x  int
	y  int
}

func (p point) isInfinite(left, top, right, bottom int) bool {
	if p.x >= right || p.x <= left || p.y <= top || p.y >= bottom {
		return true
	}

	return false
}

func readInput() []point {
	fi, err := os.Open("./input.txt")
	check(err)
	defer fi.Close()

	scanner := bufio.NewScanner(fi)

	points := []point{}
	i := 1
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		check(err)
		y, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		check(err)

		points = append(points, point{i, x, y})
		i++
	}

	return points
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func distanceToPoint(x, y int, p point) int {
	return Abs(p.x-x) + Abs(p.y-y)
}

type byDistance []struct {
	p        point
	distance int
}

func (s byDistance) Len() int {
	return len(s)
}
func (s byDistance) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byDistance) Less(i, j int) bool {
	return s[i].distance < s[j].distance
}

// Abs azd
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
