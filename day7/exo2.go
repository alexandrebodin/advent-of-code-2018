package main

import (
	"fmt"
)

const workerCount = 5

type worker struct {
	id          string
	currentWork *unit
}

func (w *worker) doWork() bool {
	fmt.Println(w.currentWork)
	if w.currentWork == nil {
		return true
	}

	w.currentWork.work()
	if w.currentWork.done() {
		return true
	}

	return false
}

func exo2(steps []step) {

	// init workers
	workers := []*worker{}
	for i := 1; i <= workerCount; i++ {
		workers = append(workers, &worker{
			id: "Worker " + string(i),
		})
	}

	finished := []rune{}
	running := []rune{}
	i := 0
	for {
		fmt.Println(finished, running)
		if len(finished) == len(steps) {
			break
		}

		available := []rune{}
		var x rune
		for _, step := range steps {
			if includes(finished, step.name) {
				continue
			}

			if includes(running, step.name) {
				continue
			}

			if len(step.parents) == 0 || finishedParents(finished, step.parents) {
				available = append(available, step.name)
			}
		}

		fmt.Println(available)

		for _, w := range workers {

			if w.doWork() { // when worker is free
				if w.currentWork != nil {
					finished = append(finished, w.currentWork.name)
					w.currentWork = nil
				}

				if len(available) > 0 {
					fmt.Println("Add work to worker")
					x, available = available[0], available[1:]
					w.currentWork = newUnit(x, charToDuration(x))
					running = append(running, x)
				}
			}
		}
		i++
	}

	fmt.Printf("Done in %d steps\n", i)

}

func charToDuration(r rune) int {
	return 60 + int(r) - 65
}
