package main

import (
	"fmt"
	"log"
	"os"

	aoc "github.com/w3dg/aocutils"
	_ "github.com/zyedidia/generic"
)

func p1(lines []string) int {
	return 0
}

func p2(lines []string) int {
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

	r1, r2 := p1(lines), p2(lines)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}
