package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	aoc "github.com/w3dg/aocutils"
	_ "github.com/zyedidia/generic"
)

func getMaxJoltageTwoDigits(bank []int) int {
	leftBatteries := bank[:len(bank)-1]
	maxLeft := leftBatteries[0]
	leftIndex := 0
	for i, v := range leftBatteries {
		if maxLeft < v {
			maxLeft = v
			leftIndex = i
		}

	}

	rightBatteries := bank[leftIndex+1:]
	maxRight := bank[leftIndex+1]
	for _, v := range rightBatteries {
		maxRight = max(v, maxRight)
	}

	// fmt.Printf("made highest volatage using %d index and %d index to be %d%d\n", leftIndex, rightIndex, maxLeft, maxRight)
	j := aoc.ParseNumOrPanic(fmt.Sprintf("%d%d", maxLeft, maxRight))
	return j
}

func p1(banks [][]int) int {
	sum := 0
	for _, bank := range banks {
		sum += getMaxJoltageTwoDigits(bank)
	}
	return sum
}

func p2(banks [][]int) int {
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

	banks := make([][]int, len(lines))
	for i, v := range lines {
		banks[i] = make([]int, 0, len(lines[i]))
		for ch := range strings.SplitSeq(v, "") {
			banks[i] = append(banks[i], aoc.ParseNumOrPanic(ch))
		}
	}

	r1, r2 := p1(banks), p2(banks)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}
