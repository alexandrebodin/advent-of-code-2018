package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type claim struct {
	id     string
	left   int
	top    int
	width  int
	height int
}

func (a *claim) intersectsWith(b claim) bool {

	if a.left >= b.left+b.width || b.left > a.left+a.width {
		return false
	}

	if a.top >= b.top+b.height || b.top >= a.top+a.height {
		return false
	}

	return true
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	claims := readInput()

	exo1(claims)
	exo2(claims)
}

func readInput() []claim {
	fi, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	claims := []claim{}

	scanner := bufio.NewScanner(fi)

	for scanner.Scan() {
		line := scanner.Text()
		// parse line
		s := strings.Split(line, "@")

		id := s[0]
		bounds := s[1]

		boundsParts := strings.Split(bounds, ":")
		pos := strings.TrimSpace(boundsParts[0])
		size := strings.TrimSpace(boundsParts[1])

		posParts := strings.Split(pos, ",")
		left, _ := strconv.Atoi(posParts[0])
		top, _ := strconv.Atoi(posParts[1])

		sizeParts := strings.Split(size, "x")
		width, _ := strconv.Atoi(sizeParts[0])
		height, _ := strconv.Atoi(sizeParts[1])

		claims = append(claims, claim{
			id,
			left,
			top,
			width,
			height,
		})
	}

	return claims
}

func exo1(claims []claim) {
	m := [1000][1000]int{}
	for _, claim := range claims {
		for i := claim.left; i < claim.left+claim.width; i++ {
			for j := claim.top; j < claim.top+claim.height; j++ {
				m[i][j]++
			}
		}
	}

	total := 0
	for _, row := range m {
		for _, val := range row {
			if val > 1 {
				total++
			}
		}
	}

	fmt.Println(total)
}

func exo2(claims []claim) {

	for i := 0; i < len(claims); i++ {
		a := claims[i]
		intersected := false

		for j := 0; j < len(claims); j++ {
			if i == j {
				continue
			}

			b := claims[j]
			if a.intersectsWith(b) {
				intersected = true
				break
			}
		}

		if !intersected {
			fmt.Println(a.id)
			break
		}
	}
}
