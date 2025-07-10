package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func p1(lines []string) int {
	return 0
}

func p2(lines []string) int {
	return 0
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

	r1, r2 := p1(lines), p2(lines)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}
