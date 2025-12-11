package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	aoc "github.com/w3dg/aocutils"
	_ "github.com/zyedidia/generic"
)

type Point struct {
	X, Y int
}

// Finds area of a rectangle that can be formed with these two points
// at the corners. Also considers them to be indices so they are adjusted by 1.
func AreaRecWithTwoPoints(a, b Point) int {

	ydiff := math.Abs(float64(a.Y-b.Y)) + 1
	if a.X == b.X {
		return int(ydiff)
	}

	xdiff := math.Abs(float64(a.X-b.X)) + 1
	if a.Y == b.Y {
		return int(xdiff)
	}

	return int(xdiff * ydiff)
}

func p1(points []Point) int {
	maxArea := 0

	for i, v1 := range points {
		for j, v2 := range points {
			if i == j || j < i {
				continue
			}
			area := AreaRecWithTwoPoints(v1, v2)
			fmt.Println("Points: ", v1, v2, "Area : ", area)

			maxArea = max(maxArea, area)
		}
	}

	return maxArea
}

func p2(points []Point) int {
// 	maxThin := 0
//
// 	for i := range len(points) {
//
// 		// straight line borders
// 		if i > 0 {
// 			pa := points[i]
// 			pb := points[i-1]
// 			maxThin = max(maxThin, AreaRecWithTwoPoints(pa, pb))
// 			// fmt.Println(maxThin)
// 		}
//
// 		altPoints := []Point{}
// 		// choose every point at index 0, 2, 4 etc
// 		if i%2 == 0 {
// 			altPoints = append(altPoints, points[i])
// 		}
// 	}
//
// 	fmt.Println(maxThin)
	return 0
}

func main() {
	infile := "./sample.txt"
	arg := os.Args[1:]
	if len(arg) != 0 {
		infile = arg[0]
	}
	lines, err := aoc.ReadFileLines(infile)
	if err != nil {
		log.Fatal("cannot read")
	}

	points := make([]Point, 0, len(lines))

	for _, s := range lines {
		p := strings.Split(s, ",")
		X, Y := aoc.ParseNumOrPanic(p[0]), aoc.ParseNumOrPanic(p[1])
		points = append(points, Point{X, Y})
	}

	r1, r2 := p1(points), p2(points)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}
