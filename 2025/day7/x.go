package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	aoc "github.com/w3dg/aocutils"
	g "github.com/zyedidia/generic"
	"github.com/zyedidia/generic/hashset"
)

func getColIndexOfSplitters(line []int) []int {
	indices := []int{}

	for i := range line {
		if line[i] == 1 {
			indices = append(indices, i)
		}
	}

	return indices
}

func p1(S int, grid [][]int) int {
	indicesA := hashset.New[int](uint64(len(grid[0])), g.Equals[int], g.HashInt)
	indicesB := hashset.New[int](uint64(len(grid[0])), g.Equals[int], g.HashInt)

	indicesA.Put(S)

	numSplit := 0

	for row := 1; row < len(grid); row++ {
		splitters := getColIndexOfSplitters(grid[row])
		// fmt.Println(row, splitters)
		if len(splitters) == 0 {
			// viz
			// for idx := 0; idx < len(grid[0]); idx++ {
			// 	if indicesA.Has(idx) {
			// 		fmt.Printf("|")
			// 	} else {
			// 		fmt.Printf("%d", grid[row][idx])
			// 	}
			// }
			// fmt.Println()
			// end viz
			continue
		}

		for _, v := range splitters {
			if indicesA.Has(v) {
				numSplit += 1
				indicesB.Put(v - 1)
				indicesB.Put(v + 1)
				indicesA.Remove(v)
			}
		}

		// continue all the beams that didnt split

		indicesA.Each(func(idx int) {
			if !indicesB.Has(idx) {
				indicesB.Put(idx)
			}
		})

		// viz
		// for idx := 0; idx < len(grid[0]); idx++ {
		// 	fmt.Printf("%d", grid[row][idx])
		// }
		// fmt.Println()
		// end viz

		// swap pointers to both maps
		tmp := indicesA
		indicesA = indicesB
		indicesB = tmp

		indicesB.Clear()
	}

	return numSplit
}

var memo = [][]int{}

func dfs(paths [][]int, row, col int) int {
	if row == len(paths) {
		return 1
	}

	if memo[row][col] != -1 {
		return memo[row][col]
	}

	curr := paths[row][col]
	val := 0
	if curr == 1 {
		val = dfs(paths, row+1, col-1) + dfs(paths, row+1, col+1)
	} else {
		val = dfs(paths, row+1, col)
	}
	memo[row][col] = val
	return val
}

func p2(S int, grid [][]int) int {
	paths := [][]int{}
	for i, v := range grid {
		if i%2 == 1 {
			paths = append(paths, v)
		}
	}

	for range len(paths) {
		l := []int{}
		for range len(grid[0]) {
			l = append(l, -1)
		}
		memo = append(memo, l)
	}

	return dfs(paths, 0, S)
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
	startline := lines[0]
	S := 0

	for i, v := range strings.Split(startline, "") {
		if v == "S" {
			S = i
			break
		}
	}

	lines = lines[1:]

	grid := [][]int{}

	for _, v := range lines {
		l := []int{}
		for _, val := range strings.Split(v, "") {
			if val == "^" {
				l = append(l, 1)
			} else {
				l = append(l, 0)
			}
		}
		grid = append(grid, l)
	}

	// fmt.Println(S)
	// for _, v := range grid {
	// 	fmt.Println(v)
	// }
	r1, r2 := p1(S, grid), p2(S, grid)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}
