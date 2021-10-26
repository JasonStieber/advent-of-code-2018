package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	r := convert(input)
	s := count(r)
	v := twiceReached(r)
	fmt.Printf("Solution to part 1: sum the total: %v \n", s)
	fmt.Printf("Solution to part 2: find the duplicate sum: %v\n", v)
}

func count(i []int) int {
	c := 0
	for _, v := range i {
		c += v
	}
	return c
}

func twiceReached(i []int) int {
	c := 0
	m := make(map[int]bool)
	m[c] = true
	for {
		for j := 0; j < len(i); j++ {
			c += i[j]
			if m[c] {
				return c
			} else {
				m[c] = true
			}

		}
	}
}

func convert(s string) []int {
	sl := strings.Split(s, "\n")
	var intSlice []int
	for _, v := range sl {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, i)
	}
	return intSlice
}
