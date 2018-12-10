package main

import (
	"fmt"
)

type unit struct {
	name     rune
	duration int
	timer    int
}

func (u *unit) work() {
	if u.timer > 0 {
		u.timer = u.timer - 1
	}
}

func (u *unit) done() bool {
	return u.timer == 0
}

func (u unit) String() string {
	return fmt.Sprintf("%s - %d(%d)", u.name, u.duration, u.timer)
}

func newUnit(name rune, duration int) *unit {
	return &unit{
		name,
		duration,
		duration,
	}
}
