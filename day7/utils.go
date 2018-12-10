package main

import "log"

func includes(arr []rune, s rune) bool {
	for _, val := range arr {
		if s == val {
			return true
		}
	}

	return false
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func finishedParents(finished []rune, parents []*step) bool {
	f := true
	for _, p := range parents {
		if !includes(finished, p.name) {
			f = false
		}
	}

	return f
}
