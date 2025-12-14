package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	aoc "github.com/w3dg/aocutils"
	_ "github.com/zyedidia/generic"
)

type Set map[string]int

type Game struct {
	ID   int
	Sets []Set
}

func solve(games []Game) (int, int) {
	// initially loaded with 12 red cubes, 13 green cubes, and 14 blue cubes
	sumpossible := 0
	sumpowerset := 0

	for _, g := range games {
		valid := true
		mr, mg, mb := -1, -1, -1
		for _, s := range g.Sets {
			for k, v := range s {
				switch k {
				case "blue":
					if v > 14 {
						valid = false
					}
					mb = max(mb, v)
				case "green":
					if v > 13 {
						valid = false
					}
					mg = max(mg, v)
				case "red":
					if v > 12 {
						valid = false
					}
					mr = max(mr, v)
				default:
					panic("Unknown key")
				}
			}
		}

		if valid {
			sumpossible += g.ID
		}

		powerset := mr * mb * mg
		sumpowerset += powerset
	}
	return sumpossible, sumpowerset
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

	games := []Game{}

	for _, l := range lines {
		g := strings.Split(l, ": ")
		gameid := aoc.ParseNumOrPanic(strings.TrimPrefix(g[0], "Game "))
		sets := []Set{}
		setsstring := strings.Split(g[1], "; ")
		for _, set := range setsstring {
			balls := strings.Split(set, ", ")
			s := make(Set)
			for _, b := range balls {
				p := strings.Split(b, " ")
				count, color := aoc.ParseNumOrPanic(p[0]), p[1]
				s[color] = count
			}
			sets = append(sets, s)
		}
		games = append(games, Game{
			ID:   gameid,
			Sets: sets,
		})
	}

	r1, r2 := solve(games)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}
