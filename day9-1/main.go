package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type node struct {
	children []node
	metadata []int
}

func main() {
	fi, err := os.Open("./input.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer fi.Close()

	var players int
	var marble int
	s := bufio.NewScanner(fi)
	for s.Scan() {
		str := s.Text()
		fmt.Sscanf(str, "%d players; last marble is worth %d points", &players, &marble)
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	scores := []int{0}

	currentIndex := 0
	currentPlayer := 1

	pointsPerPlayer := map[int]int{}

	i := 1
	for {
		if i > marble {
			break
		}

		if i%23 == 0 {
			pointsPerPlayer[currentPlayer] += i
			indexToRemove := (currentIndex - 7 + len(scores)) % len(scores)

			pointsPerPlayer[currentPlayer] += scores[indexToRemove]
			copy(scores[indexToRemove:], scores[indexToRemove+1:])
			scores[len(scores)-1] = 0
			scores = scores[:len(scores)-1]

			currentIndex = indexToRemove

		} else {
			insertIndex := (currentIndex+1)%len(scores) + 1
			scores = append(scores, 0)
			copy(scores[insertIndex+1:], scores[insertIndex:])
			scores[insertIndex] = i
			currentIndex = insertIndex
		}
		i++
		currentPlayer = currentPlayer%players + 1
	}

	max := 0
	maxPlayer := 0
	for i, val := range pointsPerPlayer {
		if val > max {
			max = val
			maxPlayer = i
		}
	}

	fmt.Printf("Max player: %d, width value: %d\n", maxPlayer, max)
}

func printValues(scores []int, currentPlayer, currentIndex int) {
	fmt.Printf("[%d] ", currentPlayer)

	for i, v := range scores {
		if i == currentIndex {
			fmt.Printf("(%d)", v)
		} else {
			fmt.Printf("%d", v)
		}

		fmt.Print(" ")
	}
	fmt.Printf("\n")
}
