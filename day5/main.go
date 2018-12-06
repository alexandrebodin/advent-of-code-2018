package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	polymer := readPolymer()

	exo1(polymer)
	exo2(polymer)
}

func readPolymer() []byte {
	fi, err := os.Open("./input.txt")
	check(err)
	defer fi.Close()

	scanner := bufio.NewScanner(fi)

	scanner.Scan()
	polymer := scanner.Text()
	return []byte(polymer)
}

func exo2(polymer []byte) {

	values := map[rune]int{}
	for i := 'a'; i <= 'z'; i++ {

		a := bytes.Replace(polymer, []byte(string(i)), []byte(""), -1)
		b := bytes.Replace(a, []byte(strings.ToUpper(string(i))), []byte(""), -1)

		r := reactPolymer(b)

		values[i] = len(r)
	}

	min := struct {
		c      rune
		length int
	}{}

	for k, val := range values {
		if min.c == 0 {
			min.c = k
			min.length = val
			continue
		}

		if val < min.length {
			min.c = k
			min.length = val
		}
	}

	fmt.Printf("%c - %d\n", min.c, min.length)
}

func exo1(polymer []byte) {
	reactedPolymer := reactPolymer(polymer)
	fmt.Println("Result:", len(reactedPolymer))
}

func reactPolymer(polymer []byte) []byte {
	result := polymer
	for {
		tmp := []byte{}
		opCount := 0
		for i := 0; i < len(result)-1; {
			if result[i] != result[i+1] && (result[i] == bytes.ToLower(result[i+1 : i+2])[0] || bytes.ToLower(result[i : i+1])[0] == result[i+1]) {
				opCount++
				i += 2
				continue
			}

			tmp = append(tmp, result[i])
			i++
		}

		// add remaining char
		tmp = append(tmp, result[len(result)-1])

		if opCount == 0 {
			break
		}
		result = tmp
	}
	return result
}
