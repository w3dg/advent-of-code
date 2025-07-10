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

func num(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic("cannot convert to int")
	}

	return n
}

func cumulativeAddToMap(m *map[int]int, key, val int) {
	_m := *m
	if _, ok := _m[key]; !ok {
		_m[key] = val
	} else {
		_m[key] += val
	}
}

func totalValueCount(m map[int]int) int {
	t := 0
	for _, v := range m {
		t += v
	}
	return t
}

func countDigits(x int) int {
	if x < 0 {
		panic("less than 0")
	}
	c := 0
	for x > 0 {
		x = x / 10
		c++
	}

	return c
}

func splitStone(x int) (int, int) {
	s := strconv.Itoa(x)
	l := len(s)
	return num(string(s[:l/2])), num(string(s[l/2:]))
}

func solution(stones []int) (int, int) {
	// 25, 75
	blinks := 0
	after25 := 0
	total := 0

	m1, m2 := make(map[int]int), make(map[int]int)
	side := 0

	for _, v := range stones {
		if _, ok := m1[v]; !ok {
			m1[v] = 1
		} else {
			m1[v]++
		}
	}

	current := []map[int]int{m1, m2}

	for blinks < 75 {
		nextSide := 1 - side

		for key, value := range current[side] {
			switch {
			case key == 0:
				cumulativeAddToMap(&current[nextSide], 1, value)
			case countDigits(key)%2 == 0:
				st1, st2 := splitStone(key)
				cumulativeAddToMap(&current[nextSide], st1, value)
				cumulativeAddToMap(&current[nextSide], st2, value)
			default:
				cumulativeAddToMap(&current[nextSide], key*2024, value)
			}
		}
		blinks++
		// fmt.Println("after", blinks, "days state", current[nextSide], totalValueCount(current[nextSide]))
		if blinks == 25 {
			after25 = totalValueCount(current[nextSide])
		}
		current[side] = make(map[int]int)
		side = nextSide
	}
	total = totalValueCount(current[side])

	return after25, total
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

	stones := []int{}
	for v := range strings.SplitSeq(lines[0], " ") {
		stones = append(stones, num(v))
	}

	r1, r2 := solution(stones)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}
