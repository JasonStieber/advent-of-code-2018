package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type log struct {
	id   int
	info string
	date time.Time
}

type logs []log

func (s logs) Less(i, j int) bool { return s[i].date.Before(s[j].date) }
func (s logs) Swap(i, j int)      { s[i].date, s[j].date = s[j].date, s[i].date }
func (s logs) Len() int           { return len(s) }

const layout = "2006-01-02 15:04"

func main() {
	strLog := strings.Split(input, "\n")
	var sheet logs
	for _, v := range strLog {
		l := parse(v)
		sheet = append(sheet, l)
		//fmt.Printf("Agent #:%v %v on %v \n", l.id, l.info, l.date)
	}
	sort.Sort(sheet)
	currentID := 0
	for _, v := range sheet {
		if v.id == 0 {
			v.id = currentID
		} else {
			currentID = v.id
		}
		fmt.Printf("%v %v %v\n", v.date.Format("[2006-01-02 15:04]"), v.id, v.info)
		//fmt.Printf("Agent #%v %v on %v \n", v.id, v.info, v.date)
	}

}

func parse(s string) log {
	l := log{}
	//	fmt.Printf("%#+v\n", s)
	re := regexp.MustCompile(`^\[(\d\d\d\d-\d\d-\d\d \d\d:\d\d)\] Guard #?(\d*) (.+)|^\[(\d\d\d\d-\d\d-\d\d \d\d:\d\d)\] (.*)$`)
	m := re.FindStringSubmatch(s)
	//	fmt.Printf("This is the match %#+v \n", m[1:])
	//fmt.Printf("Our matches looks like this: %#+v\n", m[1:])
	if m[1] != "" {
		t, err := time.Parse(layout, m[1])
		if err != nil {
			panic(err)
		}
		i, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		l.date = t
		l.id = i
		l.info = m[3]
	} else {
		t, err := time.Parse(layout, m[4])
		if err != nil {
			panic(err)
		}
		l.date = t
		l.info = m[5]
	}
	//	fmt.Printf("%+#v \n", l.info)
	return l

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
