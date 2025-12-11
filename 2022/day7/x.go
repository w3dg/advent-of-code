package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	aoc "github.com/w3dg/aocutils"
	_ "github.com/zyedidia/generic"
)

const TOTAL_FS_SIZE = 70000000
const REQUIRED_FREE_FS_SIZE = 30000000

var CURRENT_FREE int = 0

var DirSizesUnderLimit = make(map[*Node]int)

type DirSize struct {
	Dir  *Node
	Size int
}

var DirSizes = []DirSize{}

func calcSizeOfDir(node *Node) int {
	size := 0
	ch := *node.Children

	if len(ch) == 0 {
		return 0
	}
	dirs := []*Node{}

	for _, v := range ch {
		if v.Type == NODE_TYPE_FILE {
			size += v.Size
			// fmt.Println(v)
		} else {
			// it is a dir
			dirs = append(dirs, v)
		}
	}

	for _, d := range dirs {
		size += calcSizeOfDir(d)
	}

	// fmt.Println("dir", node.Name, "total size", size)
	DirSizes = append(DirSizes, DirSize{node, size})

	// keep track of dirs at most 100000
	if size <= 100_000 {
		DirSizesUnderLimit[node] = size
	}

	return size
}

func p1(root *Node) int {
	rootsize := calcSizeOfDir(root)

	// calculate metric for part 2
	CURRENT_FREE = TOTAL_FS_SIZE - rootsize

	result := 0
	for _, v := range DirSizesUnderLimit {
		result += v
	}

	return result
}

func p2() int {
	slices.SortFunc(DirSizes, func(a, b DirSize) int {
		return a.Size - b.Size
	})

	for _, v := range DirSizes {
		if CURRENT_FREE+v.Size >= REQUIRED_FREE_FS_SIZE {
			return v.Size
		}
	}

	return -1
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

	commands := ParseIntoCommands(lines)
	// for _, c := range commands {
	// 	fmt.Println(c)
	// }

	root := MakeFSFromCommands(commands)

	r1, r2 := p1(root), p2()
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}
