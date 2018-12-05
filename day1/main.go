package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	numbers := readInput()

	exo1(numbers)
	exo2(numbers)
}

func readInput() []int {
	numbers := []int{}

	fi, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(fi)

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, val)
	}

	return numbers
}

func exo2(numbers []int) {
	result := 0
	freqMap := map[int]bool{0: true}

	for i := 0; ; i++ {
		result += numbers[i%len(numbers)]
		if freqMap[result] {
			fmt.Println(result)
			break
		}
		freqMap[result] = true
	}

}

func exo1(numbers []int) {
	result := 0

	for _, val := range numbers {
		result += val
	}

	fmt.Println(result)
}
