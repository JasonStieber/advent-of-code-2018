package main

import (
	"log"
	"strings"
)

type task struct {
	step   string
	preRec []string
	fin    bool
}

func (t task) ready() bool {
	return len(t.preRec) == 0
}

func main() {
	steps := parse(input)
	o := order(steps)
	log.Printf("The answer to part A is that the steps should go in: %v order \n", o)
}

func order(steps []task) string {
	o := ""
	for i := 0; i < len(steps); i++ {
		s := steps[i]
		if !s.fin {
			newRecs := []string{}
			for j := 0; j < len(s.preRec); j++ {
				if !steps[s.preRec[j][0]-65].fin {
					newRecs = append(newRecs, s.preRec[j])
				}
			}
			steps[i].preRec = newRecs
			if steps[i].ready() {
				o += s.step
				steps[i].fin = true
				i = -1

			}
		}
	}
	return o
}

func buildEmptySteps() []task {
	instra := []task{}
	for r := 'A'; r <= 'Z'; r++ {
		t := task{}
		t.step = string(r)
		instra = append(instra, t)
	}
	return instra
}

// Step E must be finished before step H can begin.
func parse(s string) []task {
	steps := buildEmptySteps()
	for _, line := range strings.Split(s, "\n") {
		p := line[5:6]
		st := line[36]
		steps[st-65].preRec = append(steps[st-65].preRec, p)
	}
	return steps
}
