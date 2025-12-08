package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"

	aoc "github.com/w3dg/aocutils"
	"github.com/zyedidia/generic/heap"
)

type Position struct {
	X, Y, Z int
}

type Pair struct {
	A, B Position
}

type Distance struct {
	Value  float64
	Points Pair
}

func straightLineDistance(pair Pair) float64 {
	val := math.Sqrt(math.Pow(float64(pair.A.X-pair.B.X), 2) +
		math.Pow(float64(pair.A.Y-pair.B.Y), 2) +
		math.Pow(float64(pair.A.Z-pair.B.Z), 2))
	return val
}

func solve(positions []Position) {
	distances := []Distance{}
	uf := NewUnionFind[Position]()

	farFromWall := 0

	for i, v1 := range positions {
		uf.Add(v1)
		for j, v2 := range positions {
			if j <= i {
				continue
			}
			pair := Pair{v1, v2}
			dist := straightLineDistance(pair)
			distances = append(distances, Distance{dist, pair})
		}
	}

	h := heap.From(func(a, b Distance) bool { return a.Value < b.Value }, distances...)

	// iter := 10
	iter := 0

	for h.Size() > 0 {

		if iter == 1000 {
			m := uf.GetGroupSizes()
			s := []int{}
			for k, v := range m {
				for range v {
					s = append(s, k)
				}
			}
			slices.Sort(s)
			// fmt.Printf("s: %v\n", s)

			l := len(s)
			threeLargest := s[l-1] * s[l-2] * s[l-3]
			fmt.Println("Part 1:", threeLargest)
		}

		dist, ok := h.Pop()
		if !ok {
			break
		}

		// fmt.Println("Next chosen: ", dist.Points.A, dist.Points.B)

		aFind := uf.Find(dist.Points.A)
		bFind := uf.Find(dist.Points.B)

		if aFind == bFind {
			// fmt.Printf("Comparing %v,  %v - they are already connected\n", dist.Points.A, dist.Points.B)
		} else {
			uf.Union(aFind, bFind)
			if len(uf.GetGroups()) == 1 {
				fmt.Printf("Just joined %v and %v to get one single circuit\n", dist.Points.A, dist.Points.B)
				farFromWall = dist.Points.A.X * dist.Points.B.X
				fmt.Println("Part 2:", farFromWall)
				return
			}
		}

		// fmt.Println(uf)
		iter++
	}

}

func p2(positions []Position) int {
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

	positions := make([]Position, len(lines))

	for i, v := range lines {
		coords := strings.Split(v, ",")
		positions[i] = Position{
			X: aoc.ParseNumOrPanic(coords[0]),
			Y: aoc.ParseNumOrPanic(coords[1]),
			Z: aoc.ParseNumOrPanic(coords[2]),
		}
	}

	solve(positions)
}
