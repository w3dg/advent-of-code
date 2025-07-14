package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	aoc "github.com/w3dg/aocutils"
	_ "github.com/zyedidia/generic"
)

func cost(a, b int) int {
	return 3*a + b
}

func solve(eqn SimultaneousEqn) (int, int) {
	a := (eqn.y2*eqn.targetX - eqn.x2*eqn.targetY) / (eqn.y2*eqn.x1 - eqn.x2*eqn.y1)
	b := (eqn.targetY - eqn.y1*a) / eqn.y2

	return a, b
}

func p1(eqns []SimultaneousEqn) int {

	totalCost := 0

	for _, eqn := range eqns {
		a, b := solve(eqn)

		if (eqn.x1*a+eqn.x2*b == eqn.targetX) && (eqn.y1*a+eqn.y2*b == eqn.targetY) {
			totalCost += cost(a, b)
		}
	}

	return totalCost
}

func p2(eqns []SimultaneousEqn) int {

	totalCost := 0
	for _, eqn := range eqns {
		eqn.targetX += 10000000000000
		eqn.targetY += 10000000000000
		a, b := solve(eqn)

		if (eqn.x1*a+eqn.x2*b == eqn.targetX) && (eqn.y1*a+eqn.y2*b == eqn.targetY) {
			totalCost += cost(a, b)
		}
	}

	return totalCost
}

type SimultaneousEqn struct {
	x1      int
	x2      int
	targetX int
	y1      int
	y2      int
	targetY int
}

func parseXY(s, strPrefix, variableSeparator string) (int, int) {
	s = strings.TrimPrefix(s, strPrefix)
	args := strings.Split(s, ", ")
	x := aoc.ParseNumOrPanic(strings.TrimPrefix(args[0], "X"+variableSeparator))
	y := aoc.ParseNumOrPanic(strings.TrimPrefix(args[1], "Y"+variableSeparator))

	return x, y
}

func main() {
	infile := "./sample.txt"
	arg := os.Args[1:]
	if len(arg) != 0 {
		infile = arg[0]
	}
	blocks, err := aoc.ReadFileBlocks(infile)
	if err != nil {
		log.Fatal("cannot read")
	}

	eqns := []SimultaneousEqn{}
	for _, block := range blocks {
		x1, y1 := parseXY(block[0], "Button A: ", "+")
		x2, y2 := parseXY(block[1], "Button B: ", "+")
		targetX, targetY := parseXY(block[2], "Prize: ", "=")

		eqns = append(eqns, SimultaneousEqn{x1: x1, y1: y1, targetX: targetX, x2: x2, y2: y2, targetY: targetY})
	}
	// fmt.Println(eqns)
	r1, r2 := p1(eqns), p2(eqns)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}
