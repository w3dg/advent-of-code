package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	aoc "github.com/w3dg/aocutils"
	_ "github.com/zyedidia/generic"
)

func p1(numbergrid [][]int, operators []string) int {
	results := make([]int, len(numbergrid[0]))
	copy(results, numbergrid[0])

	for _, numbers := range numbergrid[1:] {
		for i, v := range numbers {
			results[i] = calculate(results[i], v, operators[i])
		}
	}

	// fmt.Println(results)
	s := 0
	for _, v := range results {
		s += v
	}
	return s
}

func p2(numbersGroupedByColumns [][]int, operators []string) int {
	results := make([]int, len(numbersGroupedByColumns))

	for i := range len(results) {
		results[i] = numbersGroupedByColumns[i][0]
	}

	for i, numbers := range numbersGroupedByColumns {
		for _, v := range numbers[1:] {
			results[i] = calculate(results[i], v, operators[i])
		}
	}

	// fmt.Println(results)
	s := 0
	for _, v := range results {
		s += v
	}
	return s
}

func calculate(op1, op2 int, operator string) int {
	switch operator {
	case "+":
		return op1 + op2
	case "*":
		return op1 * op2
	default:
		panic("UNKNOWN Operator! Only + or * allowed")
	}
}

func parseNumberGrid(lines []string) [][]int {
	numbers := make([][]int, len(lines))

	for i, l := range lines {
		scanner := bufio.NewScanner(strings.NewReader(l))
		scanner.Split(bufio.ScanWords)

		numsOnLine := []int{}
		for scanner.Scan() {
			s := scanner.Text()
			numsOnLine = append(numsOnLine, aoc.ParseNumOrPanic(s))
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
		numbers[i] = numsOnLine
	}

	return numbers
}

func parseOperators(line string) []string {
	ops := []string{}
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		s := scanner.Text()
		ops = append(ops, s)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return ops
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

	problemgrid := lines[:len(lines)-1]

	operators := parseOperators(lines[len(lines)-1])
	numbergrid := parseNumberGrid(problemgrid)
	numbersGroupedByCols := makeNumberGrid(problemgrid)

	if len(numbersGroupedByCols) != len(operators) {
		panic("number of columns didnt match number or operators")
	}

	// fmt.Println(numbergrid)
	// fmt.Println(operators)
	r1 := p1(numbergrid, operators)
	r2 := p2(numbersGroupedByCols, operators)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}

func makeNumberGrid(problemgrid []string) [][]int {
	maxLenLine := func() int {
		m := len(problemgrid[0])
		for _, v := range problemgrid {
			m = max(m, len(v))
		}
		return m
	}()

	grid := [][]string{}
	for _, line := range problemgrid {
		charsPerLine := []string{}

		for _, v := range strings.Split(line, "") {
			charsPerLine = append(charsPerLine, v)
		}

		diffInLen := maxLenLine - len(charsPerLine)
		if diffInLen != 0 {
			charsPerLine = append(charsPerLine, " ")
		}

		grid = append(grid, charsPerLine)
	}

	// modify the delim lines
	isAllSpaces := func(grid [][]string, colIndex int) bool {
		if colIndex < 0 || colIndex >= len(grid[0]) {
			panic("COLUMN INDEX OUT OF BOUNDS")
		}

		for i := 0; i < len(grid); i++ {
			if grid[i][colIndex] != " " {
				return false
			}
		}

		return true
	}

	numbers := [][]int{}

	numPerOperator := []int{}
	for col := 0; col < len(grid[0]); col++ {
		if isAllSpaces(grid, col) {
			numbers = append(numbers, numPerOperator)
			numPerOperator = []int{}
		} else {
			n := ""
			for row := 0; row < len(grid); row++ {
				if grid[row][col] != " " {
					n += grid[row][col]
				}
			}
			numPerOperator = append(numPerOperator, aoc.ParseNumOrPanic(n))
		}
	}
	// append the last column which will not have an all spaces column after it
	numbers = append(numbers, numPerOperator)
	return numbers
}
