package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	aoc "github.com/w3dg/aocutils"
	_ "github.com/zyedidia/generic"
)

func p1(grid [][]int) (int, [][]int) {
	rolls := 0
	newgrid := [][]int{}
	DIRS := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for r := 0; r < len(grid); r++ {
		newrow := []int{}
		for c := 0; c < len(grid[0]); c++ {
			// check if this is a roll
			if grid[r][c] != 1 {
				newrow = append(newrow, 0)
				continue
			}
			count := 0
			// look around 8 adjacent positions for less than 4 rolls
			for _, dir := range DIRS {
				row, col := dir[0]+r, dir[1]+c
				if row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0]) && grid[row][col] == 1 {
					count += 1
				}
			}

			if count < 4 {
				newrow = append(newrow, 0)
				rolls += 1
			} else {
				newrow = append(newrow, 1)
			}
		}
		newgrid = append(newgrid, newrow)
	}

	return rolls, newgrid
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

	grid := make([][]int, 0, len(lines))

	for _, l := range lines {
		vals := []int{}
		for v := range strings.SplitSeq(l, "") {
			x := 0
			if v == "@" {
				x = 1
			}
			vals = append(vals, x)
		}

		grid = append(grid, vals)
	}

	count := 0
	r1, newgrid := p1(grid)
	removedRolls := r1
	count += removedRolls
	for removedRolls > 0 {
		removedRolls, newgrid = p1(newgrid)
		count += removedRolls
	}
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", count)
}
