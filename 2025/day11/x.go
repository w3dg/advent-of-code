package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	aoc "github.com/w3dg/aocutils"
	_ "github.com/zyedidia/generic"
)

var result1 int = 0
var result2 int = 0

func dfs(g *Graph, start *Node) {
	if start.Val == "out" {
		result1 += 1
		return
	}

	for _, adjNode := range start.edges {
		dfs(g, adjNode)
		// fmt.Print(k + " ")
	}
}

func p1(g *Graph) int {
	start, ok := g.Vertices["you"]
	if !ok {
		panic("No you vertex to start from")
	}

	dfs(g, start)
	return result1
}

// -----------------------------------------

var memo map[*Node]bool

func dfsWithChecks(g *Graph, start *Node, visitedDAC, visitedFFT bool) bool {
	// if we can reach `out` return true and store in memo

	// cache with more info

	return false
}

func p2(g *Graph) int {
	svr, ok := g.Vertices["svr"]
	if !ok {
		panic("No you vertex to start from")
	}

	dfs(g, svr)
	return result1
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

	g := NewGraph()

	for _, line := range lines {
		parts := strings.Split(line, " ")
		n, adj := strings.TrimSuffix(parts[0], ":"), parts[1:]
		g.addVertex(n)
		g.addOrUpdateEdges(n, adj)
	}

	// fmt.Println("Puzzle 1:", p1(g))
	fmt.Println("Puzzle 2:", p2(g))
}
