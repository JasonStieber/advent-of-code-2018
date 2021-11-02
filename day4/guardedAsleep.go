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
	for i := range sheet {

		if sheet[i].id == 0 {
			sheet[i].id = currentID
		} else {
			currentID = sheet[i].id
		}
		//fmt.Printf("%v %v %v\n", v.date.Format("[2006-01-02 15:04]"), v.id, v.info)
		//fmt.Printf("Agent #%v %v on %v \n", v.id, v.info, v.date)
	}

	sleepCount := howLongSleeping(sheet)
	t, i := 0.0, 0
	for k, v := range sleepCount {
		if t < v {
			t, i = v, k
		}
	}
	times := populateSleepTime(i, sheet)
	fmt.Printf("The Sleepest Guard is Guard #%v \n", i)
	sleepestMin := ""
	count := 0
	for k, v := range times {
		fmt.Printf("Guard #1777 was asleep at %v %v times\n", k, v)
		if v > count {
			count = v
			sleepestMin = k
		}
	}
	//fmt.Printf("The sleepy guard sleeps %v", times)
	fmt.Printf("The sleepest min is %v\n", sleepestMin)
}

func populateSleepTime(id int, sheet logs) map[string]int {
	sleepLog := make(map[string]int)
	for i := 0; i < len(sheet); i++ {
		if sheet[i].id == id && sheet[i].info == "falls asleep" {
			//fmt.Printf("Sleep time: %v wakeup time: %v\n", sheet[i].date.Format("3:04PM"), sheet[i+1].date.Format("3:04PM"))
			for t := sheet[i].date; t != sheet[i+1].date; t = t.Add(time.Minute) {
				//	fmt.Printf("time tracking is: %v\n", t.Format("3:04pm"))
				sleepLog[t.Format("15:04")]++
			}
		}
	}
	return sleepLog
}

func howLongSleeping(s logs) map[int]float64 {
	sum := make(map[int]float64)
	for i := 0; i < len(s); i++ {
		//fmt.Printf("Agent #%v %v on %v \n", s[i].id, s[i].info, s[i].date)
		if s[i].info == "falls asleep" {
			//	fmt.Printf("Guard #%v fell asleep", s[i].id)
			sum[s[i].id] += (s[i+1].date.Sub(s[i].date)).Minutes()
			//	fmt.Printf("Here is the time map: %v", sum)
			//	fmt.Printf("total time asleep is %v \n", sum[s[i].id])

		}
	}
	return sum
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
