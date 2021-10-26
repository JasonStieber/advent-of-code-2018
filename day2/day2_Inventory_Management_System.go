package main

import (
	"fmt"
	"strings"
)

func main() {
	s := sliceIt(input)
	r := checksum(s)
	dif := findDifference(s)
	fmt.Printf("The ansewr to part A: what is our checksum total is %v \n", r)
	fmt.Printf("The ansewr to part B: what is the close SN of boxes %v \n", dif)

}

func findDifference(s []string) string {
	v := ""
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			b, str := isOneDifferent(s[i], s[j])
			if b {
				return str
			}
		}
	}
	return v
}

func checksum(s []string) int {
	doub := 0
	trip := 0
	for _, v := range s {
		d, t := dubOrTrip(v)
		if d {
			doub++
		}
		if t {
			trip++
		}
	}
	return doub * trip
}

func isOneDifferent(s1, s2 string) (bool, string) {
	splt1 := strings.Split(s1, "")
	splt2 := strings.Split(s2, "")
	d := 0
	loc := 0
	for i := 0; i < len(splt1); i++ {
		if splt1[i] != splt2[i] {
			d++
			loc = i
		}
	}
	if d == 1 {
		nslice := append(splt1[:loc], splt1[loc+1:]...)
		return true, strings.Join(nslice, "")
	}
	return false, ""
}

func sliceIt(s string) []string {
	v := strings.Split(s, "\n")
	return v
}

func dubOrTrip(s string) (bool, bool) {
	d, t := false, false
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	for _, v := range m {
		if v == 2 {
			d = true
		} else if v == 3 {
			t = true
		}
	}
	return d, t
}
