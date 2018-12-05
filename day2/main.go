package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	lines := readInput()

	exo1(lines)

}

func readInput() []string {
	lines := []string{}

	fi, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(fi)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func exo1(lines []string) {
	idsWithTwoLetters := 0
	idsWithThreeLetters := 0

	for _, line := range lines {

		freqMap := map[rune]int{}

		// build frequency map
		for _, letter := range line {
			if _, ok := freqMap[letter]; !ok {
				freqMap[letter] = 1
			} else {
				freqMap[letter]++
			}
		}

		// check if any char has a freq of 3
		for _, value := range freqMap {
			if value == 3 {
				idsWithThreeLetters++
				break
			}
		}

		// check if any char has a freq of 2
		for _, value := range freqMap {
			if value == 2 {
				idsWithTwoLetters++
				break
			}
		}

	}

	fmt.Println(idsWithTwoLetters * idsWithThreeLetters)
}
