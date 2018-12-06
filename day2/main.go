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
	exo2(lines)
}

func readInput() []string {
	lines := []string{}

	fi, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

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

func exo2(lines []string) {

	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {

			matches, commonLetters := compare(lines[i], lines[j])

			if matches {
				for _, l := range commonLetters {
					fmt.Printf("%c", l)
					fmt.Printf("\n")
				}
				break
			}
		}
	}
}

func compare(lineA, lineB string) (bool, []byte) {
	commonLetters := []byte{}

	for i := 0; i < len(lineA); i++ {
		if lineA[i] == lineB[i] {
			commonLetters = append(commonLetters, lineA[i])
		}

	}

	if len(commonLetters) == len(lineA)-1 {
		return true, commonLetters
	}

	return false, commonLetters
}
