package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type fabricMap map[int]map[int]int

func main() {

	m := readInput()

	fmt.Println(m)
}

func readInput() fabricMap {
	fi, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	m := fabricMap{}

	scanner := bufio.NewScanner(fi)

	for scanner.Scan() {
		line := scanner.Text()
		// parse line
		s := strings.Split(line, "@")

		id := strings.TrimSpace(s[0])
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

		fmt.Println(id, left, top, width, height)
	}

	return m
}
