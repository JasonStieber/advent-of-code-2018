package main

import (
	"log"
	"strings"
)

type task struct {
	step   string
	preReq []string
	status state
}

type state int

const (
	unassigned state = 0
	working    state = 1
	finished   state = 2
)

type worker struct {
	start, end int
	status     state
	task       string
}

func (t task) ready() bool {
	return len(t.preReq) == 0 && t.status == unassigned
}

func main() {
	steps := parse(input)
	// test := parse(testInput)
	t := buildWithHelp(steps, 5)
	steps = parse(input)
	o := order(steps)
	log.Printf("The answer to part A is that the steps should go in: %v order \n", o)
	log.Printf("The answer to part B TEST is: %v time in seconds \n", t)

}

func buildWithHelp(steps []task, w int) int {
	workers := buildWorkForce(w)
	time := 0
	bwhOrder := ""
	for len(bwhOrder) < len(steps) {
		for j := 0; j < len(steps); j++ {
			s := steps[j]
			if s.ready() {
				for i := 0; i < len(workers); i++ {
					if workers[i].status != working {
						workers[i].start = time
						workers[i].end = time + int((s.step[0] - 'A' + 61))
						workers[i].status = working
						workers[i].task = s.step
						steps[j].status = working
						break
					}
				}
			}
		}
		// update worker and time
		nextFree := soonishWorkerFinish(workers, time)
		time = workers[nextFree].end
		steps = completeStep(steps, workers[nextFree].task)
		workers[nextFree].status = finished
		if len(bwhOrder) == 0 || bwhOrder[len(bwhOrder)-1:] != workers[nextFree].task {
			bwhOrder += workers[nextFree].task
		}

		//	log.Printf("task order so far %v \n\n\n At time :%03d\n our workers are doing %v\n\n\n this is our tasks%v\n\n", bwhOrder, time, workers, steps)
	}
	for i := 0; i < len(workers); i++ {
		if time < workers[i].end {
			time = workers[i].end
		}
	}
	return time
}

func completeStep(steps []task, step string) []task {
	updatedSteps := []task{}
	for _, v := range steps {
		if v.step == step {
			v.status = finished
		} else {
			v.preReq = genNewRecs(v, step)
		}
		updatedSteps = append(updatedSteps, v)
	}
	// log.Printf("These are the steps %v", updatedSteps)
	return updatedSteps
}

func genNewRecs(t task, r string) []string {
	newRecs := []string{}
	for j := 0; j < len(t.preReq); j++ {
		if t.preReq[j] != r {
			newRecs = append(newRecs, t.preReq[j])
		}
	}
	return newRecs
}

func soonishWorkerFinish(w []worker, t int) int {
	soonest := -1
	for i := 0; i < len(w); i++ {
		if soonest == -1 {
			if w[i].status == working && w[i].end > t {
				soonest = i
			}
			continue
		} else if w[i].end < w[soonest].end && w[i].end != 0 && w[i].end > t {
			soonest = i
		}
	}
	return soonest
}

func buildWorkForce(w int) (workers []worker) {
	for i := 0; i < w; i++ {
		w := worker{0, 0, 0, ""}
		workers = append(workers, w)
	}
	return workers
}

func order(steps []task) (ordered string) {
	for len(ordered) < len(steps) {
		for i, s := range steps {
			if s.status != finished {
				newRecs := []string{}
				for j := 0; j < len(s.preReq); j++ {
					if steps[s.preReq[j][0]-65].status != finished {
						newRecs = append(newRecs, s.preReq[j])
					}
				}
				steps[i].preReq = newRecs
				if steps[i].ready() {
					ordered += s.step
					steps[i].status = finished
					break
				}
			}
		}
	}
	return ordered
}

func buildEmptySteps() (tasks []task) {
	for r := 'A'; r <= 'Z'; r++ {
		tasks = append(tasks, task{step: string(r)})
	}
	return tasks
}

// Step E must be finished before step H can begin.
func parse(s string) []task {
	steps := buildEmptySteps()
	for _, line := range strings.Split(s, "\n") {
		p := line[5:6] // 'E'
		st := line[36] // 'H'
		steps[st-65].preReq = append(steps[st-65].preReq, p)
	}
	return steps
}
