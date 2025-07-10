package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

func p1(l1, l2 []int) int {
	if len(l1) != len(l2) {
		log.Fatal("Lists of unequal length")
	}

	result := 0

	for i := 0; i < len(l1); i++ {
		v1, v2 := l1[i], l2[i]
		if v1 > v2 {
			result += v1 - v2
		} else {
			result += v2 - v1
		}
	}

	return result
}

func p2(l1, l2 []int) int {
	result := 0
	ptr1, ptr2 := 0, 0
	countMap := make(map[int]int)

	for ptr1 < len(l1) && ptr2 < len(l2) {
		if l2[ptr2] > l1[ptr1] {
			ptr1++
		} else if l1[ptr1] > l2[ptr2] {
			ptr2++
		} else {
			// we are at equal values and start comparing forward from here
			break
		}
	}

	for i, val := range l1 {
		if i < ptr1 {
			continue
		}

		count, ok := countMap[val]

		if ok {
			result += count * val
		} else {
			for ptr2 < len(l2) && l2[ptr2] < val { // if the vals in second list are smaller, seek ahead
				ptr2++
			}

			if ptr2 >= len(l2) {
				break
			}

			if l2[ptr2] > val {
				continue
			}

			c := 0
			for ptr2 < len(l2) && l2[ptr2] == val {
				c++
				ptr2++
			}
			countMap[val] = c
			result += c * val
		}
	}

	return result
}

func main() {
	// lines, err := ReadInputLines("./sample.txt")
	lines, err := ReadInputLines("./input.txt")
	if err != nil {
		log.Fatal("cannot read")
	}

	list1, list2 := make([]int, 0, len(lines)), make([]int, 0, len(lines))

	for _, line := range lines {
		vals := strings.Split(line, "   ")
		i, err := strconv.Atoi(vals[0])
		if err != nil {
			log.Fatal("cant parse as number")
		}
		list1 = append(list1, i)

		i, err = strconv.Atoi(vals[1])
		if err != nil {
			log.Fatal("cant parse as number")
		}
		list2 = append(list2, i)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	fmt.Println("Puzzle 1:", p1(list1, list2))
	fmt.Println("Puzzle 2:", p2(list1, list2))
}
