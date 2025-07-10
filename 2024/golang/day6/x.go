package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func ReadInputLines(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

type Guard struct {
	i           int
	j           int
	facingDir   string
	distinctPos map[int][]string
}

func (g Guard) haveSpaceAhead(area [][]string) bool {
	// also take into account that the space ahead could not have been
	// traversed, as then the guard would have walked over it and we cannot
	// put an obstacle there
	r, c := len(area), len(area[0])
	switch g.facingDir {
	case "UP":
		if _, ok := g.distinctPos[uniqueId(g.i-1, g.j)]; !ok {
			return g.i-1 >= 0
		} else {
			return false
		}
	case "DOWN":
		if _, ok := g.distinctPos[uniqueId(g.i+1, g.j)]; !ok {
			return g.i+1 < r
		} else {
			return false
		}
	case "LEFT":
		if _, ok := g.distinctPos[uniqueId(g.i, g.j-1)]; !ok {
			return g.j-1 >= 0
		} else {
			return false
		}
	case "RIGHT":
		if _, ok := g.distinctPos[uniqueId(g.i, g.j+1)]; !ok {
			return g.j+1 < c
		} else {
			return false
		}

	default:
		panic("Unreaachable")
	}
}

func isObstacle(area [][]string, i, j int) bool {
	r, c := len(area), len(area[0])
	if i >= 0 && i < r && j >= 0 && j < c && area[i][j] == "#" {
		return true
	}
	return false
}

func (g *Guard) DirIfObsPresent(area [][]string) string {
	switch g.facingDir {
	case "UP":
		return "RIGHT"
	case "DOWN":
		return "LEFT"
	case "LEFT":
		return "UP"
	case "RIGHT":
		return "DOWN"
	default:
		panic("Invalid direction")
	}
}

func (g *Guard) Advance(area [][]string) {
	switch g.facingDir {
	case "UP":
		if isObstacle(area, g.i-1, g.j) {
			g.facingDir = "RIGHT"
		} else {
			g.i = g.i - 1
		}
	case "DOWN":
		if isObstacle(area, g.i+1, g.j) {
			g.facingDir = "LEFT"
		} else {
			g.i = g.i + 1
		}
	case "LEFT":
		if isObstacle(area, g.i, g.j-1) {
			g.facingDir = "UP"
		} else {
			g.j = g.j - 1
		}

	case "RIGHT":
		if isObstacle(area, g.i, g.j+1) {
			g.facingDir = "DOWN"
		} else {
			g.j = g.j + 1
		}

	default:
		panic("Invalid direction")
	}
}

func (g *Guard) findTraversedCell(area [][]string, dir string) bool {
	i, j := g.i, g.j
	r, c := len(area), len(area[0])
	switch dir {
	case "UP":
		for i > 0 {
			if pastDirs, ok := g.distinctPos[uniqueId(i, j)]; !ok {
				i--
			} else {
				if slices.Contains(pastDirs, dir) {
					return true
				} else {
					i--
				}
			}
		}
		return false
	case "DOWN":
		for i < r {
			if pastDirs, ok := g.distinctPos[uniqueId(i, j)]; !ok {
				i++
			} else {
				if slices.Contains(pastDirs, dir) {
					return true
				} else {
					i++
				}
			}
		}
		return false
	case "LEFT":
		for j > 0 {
			if pastDirs, ok := g.distinctPos[uniqueId(i, j)]; !ok {
				j--
			} else {
				if slices.Contains(pastDirs, dir) {
					return true
				} else {
					j--
				}
			}
		}
		return false
	case "RIGHT":
		for j < c {
			if pastDirs, ok := g.distinctPos[uniqueId(i, j)]; !ok {
				j++
			} else {
				if slices.Contains(pastDirs, dir) {
					return true
				} else {
					j++
				}
			}
		}
		return false
	}

	return false
}

func solution(area [][]string, si, sj int) (int, int) {
	r, c := len(area), len(area[0])
	g := Guard{si, sj, "UP", make(map[int][]string)}
	obs := 0

	for g.i < r && g.j < c {
		// is this the first time at this location
		first := false
		if _, ok := g.distinctPos[uniqueId(g.i, g.j)]; !ok {
			first = true
		}

		// add current position and direction to the map
		g.distinctPos[uniqueId(g.i, g.j)] = append(g.distinctPos[uniqueId(g.i, g.j)], g.facingDir)

		if g.haveSpaceAhead(area) {
			nextDirIfObs := g.DirIfObsPresent(area)
			if !first {
				dirs := g.distinctPos[uniqueId(g.i, g.j)]
				if slices.Contains(dirs, nextDirIfObs) {
					obs++
				}
			} else {
				// search in that direction, until we hit a cell that is already traversed
				// if it has the same direction, then we can intersect that
				if g.findTraversedCell(area, nextDirIfObs) {
					obs++
				}
			}
		}

		g.Advance(area)
		if g.i == r || g.j == c || g.i < 0 || g.j < 0 {
			fmt.Println("off the map")
			break
		}
		// fmt.Println(g.i, g.j, g.facingDir)
	}

	// fmt.Println(g.distinctPos)
	// fmt.Println(obs)
	return len(g.distinctPos), obs
}

// part 2
// when we reach a point we have been before in,
// we check if we have space available in front of our direction on the map
// if we turn here, are we in a direction that we have been through this spot before
// and it we are back to the same orientation that we were in when we first saw it
// then we will create a loop
// hence we can place a obstacle in front of us
// if we turn here and somewhere beyond in this direction, we intercept a trodden path before, then we should come back here

func printMap(guardMap []string) {
	for _, s := range guardMap {
		fmt.Println(s)
	}
}

// func num(s string) int {
// 	n, err := strconv.Atoi(s)
// 	if err != nil {
// 		panic("cannot convert to int")
// 	}

// 	return n
// }

func parseMap(lines []string) ([][]string, int, int) {
	area := [][]string{}
	startI := 0
	startJ := 0

	for i, s := range lines {
		row := []string{}
		for j, n := range strings.Split(s, "") {
			row = append(row, n)
			if n == "^" {
				startI, startJ = i, j
			}
		}

		area = append(area, row)
	}

	return area, startI, startJ
}

func uniqueId(i, j int) int {
	return 1000*i + j
}

func main() {
	infile := "./sample.txt"
	arg := os.Args[1:]
	if len(arg) != 0 {
		infile = arg[0]
	}
	lines, err := ReadInputLines(infile)
	if err != nil {
		log.Fatal("cannot read")
	}

	printMap(lines)
	area, si, sj := parseMap(lines)
	fmt.Println(area, si, sj)
	r1, r2 := solution(area, si, sj)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}
