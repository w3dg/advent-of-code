package main

import (
	"fmt"
	"log"
	"os"

	aoc "github.com/w3dg/aocutils"
	_ "github.com/zyedidia/generic"
)

type Direction struct {
	d string
	v int
}

func p1(directions []Direction) int {
	START := 50
	password := 0

	for _, direction := range directions {
		switch direction.d {
		case "R":
			START = (START + direction.v) % 100
		case "L":
			val := START - (direction.v % 100)
			if val < 0 {
				START = 100 + val // val is negative
			} else {
				START = val
			}
		default:
			panic("Unknown direction")
		}
		if START == 0 {
			password += 1
		}
		// fmt.Println("Current Dial", START)
		// fmt.Println("Password", password)
	}

	return password
}

func p2(directions []Direction) int {
	password := 0
	START := 50

	for _, direction := range directions {
		current := START
		for i := 0; i < direction.v; i++ {
			if direction.d == "L" {
				if current-1 < 0 {
					current = 99
				} else {
					current -= 1
				}
				if current == 0 {
					password += 1
				}
			} else {
				current = current + 1
				if current == 100 {
					current = 0
				}
				if current == 0 {
					password += 1
				}
			}
		}
		START = current
	}

	return password
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
	var directions []Direction

	for _, line := range lines {
		d := line[:1]
		val := aoc.ParseNumOrPanic(line[1:])
		dir := Direction{d, val}
		directions = append(directions, dir)
	}

	r1, r2 := p1(directions), p2(directions)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}
