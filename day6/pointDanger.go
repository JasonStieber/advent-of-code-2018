package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x, y    int
	unBound bool
	close   int
}

var locKey = map[int]string{
	0:  "aa",
	1:  "ab",
	2:  "ac",
	3:  "ad",
	4:  "ae",
	5:  "af",
	6:  "ag",
	7:  "ah",
	8:  "ai",
	9:  "aj",
	10: "ak",
	11: "al",
	12: "am",
	13: "an",
	14: "ao",
	15: "ap",
	16: "aq",
	17: "ar",
	18: "as",
	19: "at",
	20: "au",
	21: "av",
	22: "aw",
	23: "ax",
	24: "ay",
	25: "az",
	26: "ba",
	27: "bb",
	28: "bc",
	29: "bd",
	30: "be",
	31: "bf",
	32: "bg",
	33: "bh",
	34: "bi",
	35: "bj",
	36: "bk",
	37: "bl",
	38: "bm",
	39: "bn",
	40: "bo",
	41: "bp",
	42: "bq",
	43: "br",
	44: "bs",
	45: "bt",
	46: "bu",
	47: "bv",
	48: "bw",
	49: "bx",
}

type state int

const (
	unused state = 0
	tied   state = 1
	found  state = 2
)

type entry struct {
	cordNum       int
	distToClosest int
	state         state
	isDest        bool
}

func main() {
	points := parse(input)
	//log.Printf("this is the unNormalized points: %v\n", points)
	_, minX, _, minY := bounds(points)
	nomPoints := normalize(points, minX, minY)
	maxX, minX, maxY, minY := bounds(nomPoints)
	// log.Printf("This is the normalized input: %v\n", points)
	// log.Printf("Our new bounds are (%v,%v) (%v,%v)", minX, minY, maxX, maxY)
	matrix := makeMatrix(maxX+1, maxY+1)
	locMatrix := enterLocs(matrix, nomPoints)
	filledMatrix := fillMatrix(locMatrix, nomPoints)
	// log.Printf("this is our matrix %v", filledMatrix)
	nomPoints = checkForUnbound(filledMatrix, nomPoints)
	printGrid(filledMatrix)
	fCount := countCosest(filledMatrix, nomPoints)
	printMap(fCount)
	p, biggest := findBiggest(fCount)

	log.Printf("The answer to part A: is point %v, with %v closest squares\n", p, biggest)
	log.Panicf("The answer to part B: there are %v points that are under 10,000 unites away from all other points\n", t)
}

func printMap(m map[int]int) {
	for k, v := range m {
		log.Printf("The location %v has %v nearest points\n", k, v)
	}
}

func colorStr(s string, n int) string {
	//\033[48;5;57m      #That is, \033[48;5;<BG COLOR>m
	ns := ""
	ns = fmt.Sprintf("\033[48;5;%vm%v\033[0m", n, s)
	return ns

}

func findBiggest(m map[int]int) (int, int) {
	biggest := 0
	point := 0
	for k, v := range m {
		if v > biggest {
			biggest = v
			point = k
		}
	}
	return point, biggest
}

func countCosest(m [][]entry, p []point) map[int]int {
	near := make(map[int]int)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			e := m[i][j]
			if !e.isDest && !p[e.cordNum].unBound && e.state != tied {
				near[e.cordNum]++
			}
		}
	}
	return near
}

func printGrid(m [][]entry) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			key := ""
			if m[i][j].isDest {
				key = strings.ToUpper(locKey[m[i][j].cordNum])
				log.Printf("We have a cordinate")
			} else if m[i][j].state == tied {
				key = "*"
			} else {
				//log.Printf("The cordinate number is%v\n", m[i][j].cordNum)
				key = locKey[m[i][j].cordNum]
			}
			if key == "*" {
				fmt.Printf("%3s", colorStr(key, 0))
			} else {
				fmt.Printf("%3s", colorStr(key, 41+m[i][j].cordNum))
			}

			if j == len(m[0])-1 {
				fmt.Printf("\n")
			}
		}
	}
}

func checkForUnbound(m [][]entry, p []point) (f []point) {
	for i := 0; i < len(m); i++ { // bottom row
		p[m[i][0].cordNum].unBound = true
		// log.Printf("The cordniate closest to edge is %v \n", (m[i][0].cordNum))

	}
	for i := 0; i < len(m); i++ { // top row
		p[m[i][len(m[i])-1].cordNum].unBound = true
		log.Printf("The cordniate closest to edge is %v \n", (m[i][len(m[0])-1].cordNum))

	}
	for i := 0; i < len(m[0]); i++ { // left edge
		p[m[0][i].cordNum].unBound = true

	}
	for i := 0; i < len(m[0]); i++ { // check right edge
		p[m[len(m)-1][i].cordNum].unBound = true

	}
	return p
}

func fillMatrix(matrix [][]entry, locs []point) (filledMatrix [][]entry) {
	filledMatrix = makeMatrix(len(matrix), len(matrix[0]))
	for i := 0; i < len(filledMatrix); i++ {
		for j := 0; j < len(filledMatrix[0]); j++ {
			uRHere := filledMatrix[i][j]
			if !uRHere.isDest {
				for k := 0; k < len(locs); k++ {
					uRHere = filledMatrix[i][j]
					dist := abs(locs[k].x-i) + abs(locs[k].y-j)
					if uRHere.state == unused || dist < uRHere.distToClosest {
						//log.Printf("The distance is %v and the closest point dis is %v\n", dist, uRHere.distToClosest)
						//log.Printf("The state is %v", uRHere.state)
						filledMatrix[i][j].state = found
						filledMatrix[i][j].distToClosest = dist
						filledMatrix[i][j].cordNum = k
					} else if uRHere.distToClosest == dist {
						filledMatrix[i][j].state = tied
					}
				}
			}
		}
	}
	return filledMatrix
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func enterLocs(matrix [][]entry, cord []point) (locMapped [][]entry) {
	locMapped = makeMatrix(len(matrix), len(matrix[0]))
	for n, v := range cord {
		locMapped[v.x][v.y].isDest = true
		locMapped[v.x][v.y].cordNum = n
	}
	return locMapped
}

func makeMatrix(x, y int) (matrix [][]entry) {
	matrix = make([][]entry, x)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]entry, y)
	}
	return matrix
}

func normalize(points []point, minX int, minY int) (normalized []point) {
	for _, v := range points {
		normalized = append(normalized, point{v.x - minX, v.y - minY, false, 0})
	}
	return normalized
}

func bounds(points []point) (maxX, minX, maxY, minY int) {
	maxX, minX = math.MinInt, math.MaxInt
	maxY, minY = math.MinInt, math.MaxInt
	for _, p := range points {
		if p.x > maxX {
			maxX = p.x
		} else if p.x < minX {
			minX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		} else if p.y < minY {
			minY = p.y
		}
	}
	return maxX, minX, maxY, minY
}

func parse(s string) (points []point) {
	xys := strings.Split(s, "\n")
	for _, v := range xys {
		xy := strings.Split(v, ", ")
		var p point
		x, err := strconv.Atoi(xy[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(xy[1])
		if err != nil {
			panic(err)
		}
		p.x, p.y = x, y
		points = append(points, p)
	}
	return points
}
