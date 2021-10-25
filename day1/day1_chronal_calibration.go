package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("we did get started")
	r := convert(input)
	v := count(r)
	fmt.Print("The total value is %v", v)
}

func count(i []int) int {
	c := 0
	for _, v := range i {
		c += v
	}
	return c
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
