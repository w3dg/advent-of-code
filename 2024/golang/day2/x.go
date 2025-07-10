package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

func isSafe(report []int) bool {
	f, s := report[0], report[1]

	asc := true

	if f == s {
		return false
	} else if f > s {
		asc = false
	}

	for i := 1; i < len(report); i++ {
		if asc && (report[i]-report[i-1] < 0) {
			return false
		}
		if !asc && (report[i]-report[i-1] > 0) {
			return false
		}

		diff := math.Abs(float64(report[i]) - float64(report[i-1]))
		if diff > 3 || diff < 1 {
			return false
		}
	}

	return true
}

func p1(reports [][]int) (safe int, unsafeReports [][]int) {
	safe = 0
	unsafeReports = [][]int{}

	for _, report := range reports {
		if isSafe(report) {
			safe++
		} else {
			unsafeReports = append(unsafeReports, report)
		}
	}

	return
}

func p2(unsafeReports [][]int) (madesafe int) {
	madesafe = 0

	for _, report := range unsafeReports {
		for i := 0; i < len(report); i++ {
			tempReport := []int{}
			tempReport = append(tempReport, report[0:i]...)
			tempReport = append(tempReport, report[i+1:]...)

			if isSafe(tempReport) {
				madesafe++
				break
			}
		}
	}

	return
}

func main() {
	// lines, err := ReadInputLines("./sample.txt")
	lines, err := ReadInputLines("./input.txt")
	if err != nil {
		log.Fatal("cannot read")
	}

	reports := [][]int{}

	for _, l := range lines {
		r := []int{}

		for s := range strings.SplitSeq(l, " ") {
			if v, err := strconv.Atoi(s); err != nil {
				log.Fatal("Cannot parse")
			} else {
				r = append(r, v)
			}
		}
		reports = append(reports, r)
	}

	safe, unsafeReports := p1(reports)

	fmt.Println("Puzzle 1:", safe)
	fmt.Println("Puzzle 2:", safe+p2(unsafeReports))
}
