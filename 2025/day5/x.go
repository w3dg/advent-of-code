package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	aoc "github.com/w3dg/aocutils"
)

type IngredientID int

type Interval struct {
	start IngredientID
	end   IngredientID
}

func insertIntoIntervals(intervals []Interval, newInterval Interval) []Interval {
	var result []Interval
	i := 0
	n := len(intervals)

	// 1. Add all intervals that end before newInterval starts
	for i < n && intervals[i].end < newInterval.start {
		result = append(result, intervals[i])
		i++
	}

	// 2. Merge all intervals that overlap with newInterval
	for i < n && intervals[i].start <= newInterval.end {
		if intervals[i].start < newInterval.start {
			newInterval.start = intervals[i].start
		}
		if intervals[i].end > newInterval.end {
			newInterval.end = intervals[i].end
		}
		i++
	}

	// Add the merged newInterval
	result = append(result, newInterval)

	// 3. Add the remaining intervals (completely to the right)
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}

func mergeIntervals(intervals []Interval) []Interval {
	resultIntervals := []Interval{}

	slices.SortFunc(intervals, func(a, b Interval) int {
		return int(a.start) - int(b.start)
	})

	resultIntervals = append(resultIntervals, intervals[0])

	for i := 1; i < len(intervals); i++ {
		resultIntervals = insertIntoIntervals(resultIntervals, intervals[i])
	}

	return resultIntervals
}

func findInIntervals(intervals []Interval, val IngredientID) bool {
	for _, i := range intervals {
		if i.start <= val && val <= i.end {
			return true
		}
	}

	return false
}

func p1(intervals []Interval, ids []IngredientID) (int, []Interval) {
	fresh := 0
	mergedIntervals := mergeIntervals(intervals)
	for _, v := range ids {
		if findInIntervals(mergedIntervals, v) {
			fresh += 1
		}
	}
	return fresh, mergedIntervals
}

func p2(mergedIntervals []Interval) int {
	c := 0

	for _, in := range mergedIntervals {
		c += int(in.end-in.start) + 1
	}

	return c
}

func main() {
	infile := "./sample.txt"
	arg := os.Args[1:]
	if len(arg) != 0 {
		infile = arg[0]
	}
	blocks, err := aoc.ReadFileBlocks(infile)
	if err != nil {
		log.Fatal("cannot read")
	}

	b1, b2 := blocks[0], blocks[1]

	intervals := make([]Interval, len(b1))
	ingredients := make([]IngredientID, len(b2))

	for i, v := range b1 {
		parts := strings.Split(v, "-")
		s, e := aoc.ParseNumOrPanic(parts[0]), aoc.ParseNumOrPanic(parts[1])
		intervals[i] = Interval{IngredientID(s), IngredientID(e)}
	}

	for i, v := range b2 {
		ingredients[i] = IngredientID(aoc.ParseNumOrPanic(v))
	}

	r1, mergedIntervals := p1(intervals, ingredients)
	r2 := p2(mergedIntervals)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}
