package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	aoc "github.com/w3dg/aocutils"
	"github.com/zyedidia/generic/mapset"
)

type Range struct {
	firstId int
	lastId  int
}

func numDigits(x int) int {
	return int(math.Log10(float64(x))) + 1
}

func p1(ranges []Range) int {
	sum := 0
	for _, r := range ranges {
		for i := r.firstId; i <= r.lastId; i++ {
			n := numDigits(i)
			if n%2 == 0 {
				// even length, let us check for invalids
				p := n / 2
				t := int(math.Pow10(p))
				if i/t == i%t {
					// fmt.Println("Invalid", i)
					sum += i
				}
			}
		}
	}
	return sum
}

func splitIntoParts(x int, partlen int) []string {
	s := fmt.Sprint(x)
	n := numDigits(x)
	parts := make([]string, 0, n/partlen)

	for i := 0; i+partlen <= n; i += partlen {
		parts = append(parts, s[i:i+partlen])
	}

	return parts
}

func p2(ranges []Range) int {
	set := mapset.New[int]()
	for _, r := range ranges {
		// for each id in the range
		for id := r.firstId; id <= r.lastId; id++ {
			n := numDigits(id)
			// iterate over all valid subpart lengths (1-len(num)/2)
			for l := 1; l <= n/2; l++ {
				// 	// it is valid only if len(num) % len(subpart) == 0
				if n%l != 0 {
					continue
				}
				parts := splitIntoParts(id, l)
				// fmt.Println(id, parts, "into parts of length", l)
				// // keep a map of count of subparts
				counts := make(map[string]int)
				for _, p := range parts {
					if v, ok := counts[p]; ok {
						counts[p] = v + 1
					} else {
						counts[p] = 1
					}
				}
				// // if map contains 1 val, and len(val) * count == len(num), it is repeated
				if len(counts) == 1 {
					s := 0
					repeatedId := 0
					for k, v := range counts {
						s += len(k) * v
						repeatedId = id
					}
					if s == n {
						if !set.Has(repeatedId) {
							set.Put(repeatedId)
						}
					}
				}
				// // otherwise there isnt a way to divide them into repeated blocks, so it is not repeated
			}
		}

	}
	res := 0

	set.Each(func(key int) {
		res += key
	})

	return res
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

	line := lines[0]
	ids := strings.Split(line, ",")
	var ranges []Range

	for _, v := range ids {
		parts := strings.Split(v, "-")
		r := Range{aoc.ParseNumOrPanic(parts[0]), aoc.ParseNumOrPanic(parts[1])}
		ranges = append(ranges, r)
	}

	r1, r2 := p1(ranges), p2(ranges)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}
