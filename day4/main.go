package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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

	sleepingTimes := map[int][]int{}

	for i := 0; i < len(eventsList); {
		event := eventsList[i]

		if strings.HasPrefix(event.desc, "Guard") {
			guardId, err := strconv.Atoi(strings.Split(event.desc, " ")[1][1:])
			check(err)

			if _, ok := sleepingTimes[guardId]; !ok {
				sleepingTimes[guardId] = make([]int, 59)
			}
			currentTime := 0

			for j := i + 1; ; j++ {
				if j == len(eventsList) {
					i = j
					break
				}

				e := eventsList[j]

				if strings.HasPrefix(e.desc, "Guard") {
					i = j
					break
				}

				switch e.desc {
				case "falls asleep":
					{
						for i := currentTime; i < e.timestamp.Minute(); i++ {
							sleepingTimes[guardId][i] += 0
						}
						currentTime = e.timestamp.Minute()
					}

				case "wakes up":
					{
						for i := currentTime; i < e.timestamp.Minute(); i++ {
							sleepingTimes[guardId][i] += 1
						}
						currentTime = e.timestamp.Minute()
					}
				}

			}

			for i := currentTime; i < 59; i++ {
				sleepingTimes[guardId][i] += 0
			}
		}
	}

	var maxId int
	max := 0
	for guardId, guard := range sleepingTimes {
		total := 0
		for _, v := range guard {
			total += v
		}

		if total > max {
			max = total
			maxId = guardId
		}
	}

	maxMinute := 0
	maxMinuteVal := 0
	for i, v := range sleepingTimes[maxId] {
		if v > maxMinuteVal {
			maxMinuteVal = v
			maxMinute = i
		}
	}

	fmt.Println(maxId, maxMinute, maxId*maxMinute)
}
