package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x, y int
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
	log.Printf("this is the unNormalized points: %v\n", points)
	_, minX, _, minY := bounds(points)
	nomPoints := normalize(points, minX, minY)
	maxX, minX, maxY, minY := bounds(nomPoints)
	log.Printf("This is the normalized input: %v\n", points)
	log.Printf("Our new bounds are (%v,%v) (%v,%v)", minX, minY, maxX, maxY)
	matrix := makeMatrix(maxX+1, maxY+1)
	locMatrix := enterLocs(matrix, nomPoints)
	filledMatrix := fillMatrix(locMatrix, nomPoints)
	log.Printf("this is our matrix %v", filledMatrix)
	//	log.Printf("The answer to part A: is point %v, with %v closest squares\n", "Penis", 69)
	log.Println(bounds(parse(input)))
}

func fillMatrix(matrix [][]entry, locs []point) (filledMatrix [][]entry) {
	filledMatrix = makeMatrix(len(matrix), len(matrix[0]))
	for i := 0; i < len(filledMatrix); i++ {
		for j := 0; j < len(filledMatrix[0]); j++ {
			uRHere := filledMatrix[i][j]
			if !uRHere.isDest {
				for k := 0; k < len(locs); k++ {
					if uRHere.state == unused || dist < uRHere.distToClosest {
						uRHere
					}
					dist := abs(locs[k].x-i) + abs(locs[k].y-j)
					if uRHere.distToClosest == dist {
						filledMatrix[i][j].tie = true
					} else if dist < uRHere.distToClosest {
						filledMatrix[i][j].tie = false
						filledMatrix[i][j].distToClosest = dist
						filledMatrix[i][j].cordNum = k
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
	for _, v := range cord {
		locMapped[v.x][v.y].isDest = true
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
		normalized = append(normalized, point{v.x - minX, v.y - minY})
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
