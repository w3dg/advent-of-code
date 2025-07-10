package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
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

func findStrInDiag(matrix [][]string) (string, string, error) {
	if len(matrix) != len(matrix[0]) {
		return "", "", errors.New("Matrix must be square")
	}

	p, o := "", ""
	for i, j := 0, len(matrix)-1; i < len(matrix) && j >= 0; i, j = i+1, j-1 {
		p += matrix[i][i]
		o += matrix[i][j]
	}

	return p, o, nil
}

func p1(lines []string, matrix [][]string) int {
	xmasre := regexp.MustCompile("XMAS")
	samxre := regexp.MustCompile("SAMX")
	count := 0

	// line wise
	for _, l := range lines {
		matchesxmas, matchessamx := xmasre.FindAllString(l, -1), samxre.FindAllString(l, -1)
		count += len(matchesxmas) + len(matchessamx)
	}

	// col wise
	cols, rows := len(matrix[0]), len(matrix)

	for j := 0; j < cols; j++ {
		for i := 0; i <= rows-4; i++ {
			word := matrix[i][j] + matrix[i+1][j] + matrix[i+2][j] + matrix[i+3][j]
			if word == "XMAS" || word == "SAMX" {
				count++
			}
		}
	}

	for i := 0; i <= rows-4; i++ {
		for j := 0; j <= cols-4; j++ {
			// deep copy of the patch, not a ref to the underlying array
			patch := [][]string{}
			for idx := i; idx < i+4; idx++ {
				patch = append(patch, matrix[idx])
			}

			for idx, r := range patch {
				patch[idx] = r[j : j+4]
			}
			pdiagWord, odiagWord, err := findStrInDiag(patch)
			if err != nil {
				log.Fatal(err)
			}

			if pdiagWord == "XMAS" || pdiagWord == "SAMX" {
				count++
			}

			if odiagWord == "XMAS" || odiagWord == "SAMX" {
				count++
			}
		}
	}

	return count
}

func p2(matrix [][]string) int {
	count := 0
	cols, rows := len(matrix[0]), len(matrix)
	wordLen := 3

	for i := 0; i <= rows-wordLen; i++ {
		for j := 0; j <= cols-wordLen; j++ {
			// deep copy of the patch, not a ref to the underlying array
			patch := [][]string{}
			for idx := i; idx < i+wordLen; idx++ {
				patch = append(patch, matrix[idx])
			}

			for idx, r := range patch {
				patch[idx] = r[j : j+wordLen]
			}
			pdiagWord, odiagWord, err := findStrInDiag(patch)
			if err != nil {
				log.Fatal(err)
			}

			if (pdiagWord == "MAS" || pdiagWord == "SAM") && (odiagWord == "MAS" || odiagWord == "SAM") {
				count++
			}
		}
	}

	return count
}

func main() {
	// lines, err := ReadInputLines("./sample.txt")
	lines, err := ReadInputLines("./input.txt")
	if err != nil {
		log.Fatal("cannot read")
	}

	matrix := [][]string{}
	for _, l := range lines {
		matrix = append(matrix, strings.Split(l, ""))
	}

	r1, r2 := p1(lines, matrix), p2(matrix)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}
