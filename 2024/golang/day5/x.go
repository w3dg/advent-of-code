package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var (
	edges = make(map[int][]int)
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

func num(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic("cannot convert to int")
	}

	return n
}

func parseRules(rules []string) {
	re := regexp.MustCompile(`(\d+)\|(\d+)`)
	for _, r := range rules {
		matches := re.FindAllStringSubmatch(r, -1)
		if matches == nil {
			continue
		}
		for _, m := range matches {
			before, after := num(m[1]), num(m[2])
			edges[before] = append(edges[before], after)
		}
	}
}

func parsePrints(prints []string) [][]int {
	p := [][]int{}

	for _, l := range prints {
		pages := []int{}
		for p := range strings.SplitSeq(l, ",") {
			pages = append(pages, num(p))
		}

		p = append(p, pages)
	}

	return p
}

func p1(updates [][]int) (int, [][]int) {
	badUpdates := [][]int{}
	goodUpdates := [][]int{}

	for _, u := range updates {
		// [1,2,3,4,5]
		earlyStop := false
		for i, v := range u {
			if earlyStop {
				break
			}

			prevEl := u[0:i]
			adjEl := edges[v]

			for _, pr := range prevEl {
				if slices.Contains(adjEl, pr) {
					badUpdates = append(badUpdates, u)
					earlyStop = true
					break
				}
			}
		}

		if !earlyStop {
			goodUpdates = append(goodUpdates, u)
		}
	}

	fmt.Println("bad", len(badUpdates))
	return sumMiddle(goodUpdates), badUpdates
}

func sumMiddle(goodUpdates [][]int) int {
	res := 0
	for _, upd := range goodUpdates {
		l := len(upd)
		res += upd[l/2]
	}
	return res
}

func p2(updates [][]int) int {
	fmt.Println(updates)
	fmt.Println(edges)
	fmt.Println()

	correctedOrders := [][]int{}

	for _, upd := range updates {
		localEdges := make(map[int][]int)
		inDegree := make(map[int]int)

		// build up a graph with only given nodes but the same ordering rules
		// upd [75 97 47 61 53]
		for i, u := range upd {
			if _, ok := inDegree[u]; !ok {
				inDegree[u] = 0
			}

			rest := []int{}
			rest = append(rest, upd[0:i]...)
			rest = append(rest, upd[i+1:]...)
			for _, r := range rest {
				if slices.Contains(edges[u], r) {
					localEdges[u] = append(localEdges[u], r)
					inDegree[r]++
				}
			}
		}

		correctOrder := []int{}
		// run toposort on the upd according to the localedges and inDegree

		// repeat while indegrees map has elements
		// // find from inDegree which node has the indegree of 0
		// // append node to the correct order list
		// // for all the adjacents to that node, reduce their indegree in the map by 1
		// // remove the node from the map

		for len(inDegree) > 0 {
			starting := -1
			for k, v := range inDegree {
				if v == 0 {
					starting = k
					break
				}
			}

			if starting == -1 {
				panic("no starting node found")
			}

			correctOrder = append(correctOrder, starting)
			adjacents := localEdges[starting]

			for _, v := range adjacents {
				inDegree[v]--
			}

			delete(inDegree, starting)
		}

		correctedOrders = append(correctedOrders, correctOrder)
		fmt.Println(upd, correctOrder)
		// fmt.Println("for update ", upd, "localEdges", localEdges, "indegrees", inDegree)
	}

	return sumMiddle(correctedOrders)
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

	rules := []string{}
	prints := []string{}

	hitEmptyLine := false

	for _, l := range lines {
		if !hitEmptyLine && len(l) == 0 {
			hitEmptyLine = true
			continue
		}

		if hitEmptyLine {
			prints = append(prints, l)
		} else {
			rules = append(rules, l)
		}
	}

	parseRules(rules)
	updates := parsePrints(prints)
	r1, badUpdates := p1(updates)
	r2 := p2(badUpdates)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}
