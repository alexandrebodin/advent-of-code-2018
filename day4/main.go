package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type event struct {
	timestamp time.Time
	desc      string
}

type events []event

func (a events) Len() int {
	return len(a)
}

func (a events) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a events) Less(i, j int) bool {
	return a[i].timestamp.Before(a[j].timestamp)
}

func main() {
	fi, err := os.Open("./input.txt")
	check(err)
	defer fi.Close()

	scanner := bufio.NewScanner(fi)

	eventsList := events{}

	for scanner.Scan() {
		line := scanner.Text()

		timestamp := line[6:17]
		desc := line[19:]

		t, err := time.Parse("01-02 15:04", timestamp)
		check(err)

		eventsList = append(eventsList, event{t, desc})
	}

	sort.Sort(eventsList)

	for _, event := range eventsList {
		fmt.Println(event.timestamp.Format("01-02 15:04"), event.desc)
	}
}
