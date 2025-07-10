package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

func handleMul(inst string) int {
	inst = strings.TrimPrefix(inst, "mul(")
	inst = strings.TrimSuffix(inst, ")")
	digits := strings.Split(inst, ",")
	x, err := strconv.Atoi(digits[0])
	if err != nil {
		log.Fatal("Cannot parse X")
	}
	y, err := strconv.Atoi(digits[1])
	if err != nil {
		log.Fatal("Cannot parse Y")
	}
	return x * y
}

func answer(lines []string) (int, int) {
	r1, r2 := 0, 0
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)

	enabled := true
	for _, l := range lines {
		matches := re.FindAllString(l, -1) // get all matches
		for _, m := range matches {
			switch m {
			case "don't()":
				enabled = false
			case "do()":
				enabled = true
			default:
				calc := handleMul(m)
				if enabled {
					r2 += calc
				}
				r1 += calc
			}
		}
	}
	return r1, r2
}

func main() {
	// lines, err := ReadInputLines("./sample.txt")
	lines, err := ReadInputLines("./input.txt")
	if err != nil {
		log.Fatal("cannot read")
	}

	r1, r2 := answer(lines)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}
