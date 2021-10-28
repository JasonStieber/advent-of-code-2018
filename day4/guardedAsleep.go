package main

import (
	"fmt"
	"time"
)

type log struct {
	id   int
	info string
	date time.Time
}

func main() {
	fmt.Printf("we be rocking")
}

func parse(s string) log {
	`^\[(\d\d\d\d-\d\d-\d\d \d\d:\d\d)\] Guard #?(\d*) (.+)|\[(\d\d\d\d-\d\d-\d\d \d\d:\d\d)\] (.*)$`gm
}

// [1518-04-15 23:56] Guard #2213 begins shift

/*

layout := "2006-01-02T15:04:05.000Z"
str := "2014-11-12T11:45:26.371Z"
t, err := time.Parse(layout, str)

if err != nil {
    fmt.Println(err)
}
fmt.Println(t)

*/
