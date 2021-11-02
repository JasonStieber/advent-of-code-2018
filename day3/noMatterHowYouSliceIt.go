package main

import (
	"fmt"
	"strconv"
	"strings"
)

type rect struct {
	x      int
	y      int
	width  int
	height int
}

func main() {
	count := 0
	open := 0
	m := make(map[string]int)
	rects := parse(input)
	for _, r := range rects {
		count += fill(r, m)
	}
	for i, r := range rects {
		if chkForNoDups(r, m) {
			open = i + 1
		}
	}
	fmt.Printf("The answer to part A: the number of double used squares is: %v\n", count)
	fmt.Printf("The answer to part B: the only open pattern is %v \n", open)
}

func chkForNoDups(r rect, m map[string]int) bool {
	for i := r.x; i < r.x+r.width; i++ {
		for j := r.y; j < r.y+r.height; j++ {
			x := strconv.Itoa(i)
			y := strconv.Itoa(j)
			key := x + "," + y
			if m[key] == 2 {
				return false
			}
		}
	}
	return true
}

func fill(r rect, m map[string]int) int {
	c := 0
	for i := r.x; i < r.x+r.width; i++ {
		for j := r.y; j < r.y+r.height; j++ {
			x := strconv.Itoa(i)
			y := strconv.Itoa(j)
			key := x + "," + y
			if m[key] == 0 {
				m[key] = 1
			} else if m[key] == 1 {
				m[key]++
				c++
			}
		}
	}
	return c
}

func parse(s string) []rect {
	r := []rect{}
	sl := strings.Split(s, "\n")
	for _, v := range sl {
		entry := rect{}
		t1 := strings.Index(v, "@") // finds where we begin to trim our string
		t2 := strings.Index(v, ":")
		trim := v[t1+2 : t2]
		xy := strings.Split(trim, ",") // a slice of the xy cordinates
		x, err := strconv.Atoi(xy[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(xy[1])
		if err != nil {
			panic(err)
		}
		entry.x, entry.y = x, y // first two cordinates assigned
		trim = v[t2+2:]
		wh := strings.Split(trim, "x")
		w, err := strconv.Atoi(wh[0])
		if err != nil {
			panic(err)
		}
		h, err := strconv.Atoi(wh[1])
		if err != nil {
			panic(err)
		}
		entry.width = w
		entry.height = h
		r = append(r, entry)
	}
	return r
}
