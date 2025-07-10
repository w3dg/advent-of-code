package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

type Equation struct {
	target   int
	operands []int
}

func Satisfier(t int, operands []int) bool {
	if len(operands) == 0 {
		panic("length 0")
	}

	if len(operands) == 1 {
		return t == operands[0]
	}

	sumSlice := []int{}
	sumSlice = append(sumSlice, operands[0]+operands[1])
	sumSlice = append(sumSlice, operands[2:]...)

	prodSlice := []int{}
	prodSlice = append(prodSlice, operands[0]*operands[1])
	prodSlice = append(prodSlice, operands[2:]...)

	return Satisfier(t, sumSlice) || Satisfier(t, prodSlice)
}

func ConcatSatisfier(t int, operands []int) bool {
	if len(operands) == 0 {
		panic("length 0")
	}

	if len(operands) == 1 {
		return t == operands[0]
	}

	sumSlice := []int{}
	sumSlice = append(sumSlice, operands[0]+operands[1])
	sumSlice = append(sumSlice, operands[2:]...)

	prodSlice := []int{}
	prodSlice = append(prodSlice, operands[0]*operands[1])
	prodSlice = append(prodSlice, operands[2:]...)

	concatSlice := []int{}
	concatResult := strconv.Itoa(operands[0]) + strconv.Itoa(operands[1])
	concatSlice = append(concatSlice, num(concatResult))
	concatSlice = append(concatSlice, operands[2:]...)

	return ConcatSatisfier(t, sumSlice) || ConcatSatisfier(t, prodSlice) || ConcatSatisfier(t, concatSlice)
}

func solution(eqns []Equation) (int, int) {
	calibrationRes, secondaryCalibrationResult := 0, 0

	for _, eq := range eqns {
		if Satisfier(eq.target, eq.operands) {
			calibrationRes += eq.target
		} else if ConcatSatisfier(eq.target, eq.operands) {
			secondaryCalibrationResult += eq.target
		}
	}

	return calibrationRes, secondaryCalibrationResult + calibrationRes
}

func num(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic("cannot convert to int")
	}

	return n
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

	eqns := []Equation{}

	for _, l := range lines {
		p := strings.Split(l, " ")
		target, opds := num(strings.TrimSuffix(p[0], ":")), p[1:]

		operands := []int{}
		for _, o := range opds {
			operands = append(operands, num(o))
		}

		eqns = append(eqns, Equation{target, operands})
	}
	r1, r2 := solution(eqns)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}
