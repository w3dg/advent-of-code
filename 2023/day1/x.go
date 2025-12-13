package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	aoc "github.com/w3dg/aocutils"
	_ "github.com/zyedidia/generic"
)

func p1(lines []string) int {
	sum := 0
	re := regexp.MustCompile(`\d`)

	for _, l := range lines {
		digits := re.FindAll([]byte(l), -1)
		ns := ""
		if len(digits) == 0 {
			panic("No digits!")
		} else if len(digits) == 1 {
			ns = string(digits[0][0]) + string(digits[0][0])
		} else {
			ld := len(digits)
			ns = string(digits[0][0]) + string(digits[ld-1][0])
		}
		s := aoc.ParseNumOrPanic(ns)
		sum += s
	}

	return sum
}

func strtonum(s string) int {
	switch s {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	}
	panic("None of the digits matched against" + s)
}

type Digit struct {
	Value int
	Index int
}

func p2(lines []string) int {
	sum := 0
	regexes := []string{`\d`, `one`, `two`, `three`, `four`, `five`, `six`, `seven`, `eight`, `nine`}

	for _, l := range lines {
		first := Digit{Value: -1, Index: 1000}
		last := Digit{Value: -1, Index: -1}

		for _, reg := range regexes {
			re := regexp.MustCompile(reg)
			d := re.FindAllStringIndex(l, -1)
			if d == nil {
				continue
			}

			ld := len(d)
			fidx, lidx := d[0], d[ld-1]

			fn, fi, ln, li := GetDigit(l, fidx, lidx)
			if first.Index > fi {
				first.Value = fn
				first.Index = fi
			}
			if last.Index < li {
				last.Value = ln
				last.Index = li
			}
		}

		v := first.Value*10 + last.Value
		sum += v
	}

	return sum
}

func GetDigit(l string, fd, ld []int) (int, int, int, int) {
	fdiff := fd[1] - fd[0]
	ldiff := ld[1] - ld[0]
	first := 0
	last := 0

	i1, i2 := fd[0], fd[1]
	// handle number case
	n := l[i1:i2]
	if fdiff == 1 {
		first = aoc.ParseNumOrPanic(n)
	} else {
		// handle string case
		first = strtonum(n)
	}

	i1, i2 = ld[0], ld[1]
	// handle number case
	n = l[i1:i2]
	if ldiff == 1 {
		last = aoc.ParseNumOrPanic(l[i1:i2])
	} else {
		// handle string case
		numl := strtonum(l[i1:i2])
		last = numl
	}

	return first, fd[0], last, ld[0]
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

	fmt.Println("Puzzle 1:", p1(lines), "Puzzle 2:", p2(lines))
}
