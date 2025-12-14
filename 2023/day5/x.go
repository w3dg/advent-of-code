package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	aoc "github.com/w3dg/aocutils"
	_ "github.com/zyedidia/generic"
)

type FarmMapEntry struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

func (fme *FarmMapEntry) String() string {
	return fmt.Sprintf("dest_range_start: %v source_range_start: %v range_length: %v", fme.DestinationRangeStart, fme.SourceRangeStart, fme.RangeLength)
}

type FarmMap []FarmMapEntry

func (fm *FarmMap) String() string {
	s := ""
	for _, v := range *fm {
		s += fmt.Sprintln(v)
	}
	return s
}

var Name2Map map[string]*FarmMap = make(map[string]*FarmMap)

func makeFarmMap(name string, entries []string) {
	fm := make(FarmMap, len(entries))
	for i, v := range entries {
		parts := strings.Fields(v)
		dstart, sstart, rnglen := parts[0], parts[1], parts[2]
		fm[i] = FarmMapEntry{
			aoc.ParseNumOrPanic(dstart),
			aoc.ParseNumOrPanic(sstart),
			aoc.ParseNumOrPanic(rnglen),
		}
	}
	Name2Map[name] = &fm
}

func passThroughMap(input []int, fmap *FarmMap) []int {
	res := []int{}
	_map := *fmap
	for _, source := range input {
		hasPassed := false

		for _, m := range _map {
			rngstart, rngend := m.SourceRangeStart, m.SourceRangeStart+m.RangeLength
			if rngstart <= source && source < rngend {
				hasPassed = true
				offset := source - m.SourceRangeStart
				dst := m.DestinationRangeStart + offset
				res = append(res, dst)
				break
			}

		}
		if !hasPassed {
			res = append(res, source)
		}
	}
	return res
}

func p1(seeds []int) int {
	out := passThroughMap(seeds, Name2Map["seed-to-soil"])

	out = passThroughMap(out, Name2Map["soil-to-fertilizer"])
	out = passThroughMap(out, Name2Map["fertilizer-to-water"])
	out = passThroughMap(out, Name2Map["water-to-light"])
	out = passThroughMap(out, Name2Map["light-to-temperature"])
	out = passThroughMap(out, Name2Map["temperature-to-humidity"])
	finalLocations := passThroughMap(out, Name2Map["humidity-to-location"])

	minLoc := finalLocations[0]

	for _, v := range finalLocations {
		minLoc = min(v, minLoc)
	}

	return minLoc
}

func p2(seeds []int) int {
	// bigger ranges...
	// maybe merge ranges and then try again
	return 0
}

func main() {
	infile := "./sample.txt"
	arg := os.Args[1:]
	if len(arg) != 0 {
		infile = arg[0]
	}
	lines, err := aoc.ReadFileBlocks(infile)
	if err != nil {
		log.Fatal("cannot read")
	}

	seedline := strings.Fields(lines[0][0])
	seedsstr := seedline[1:]
	seeds := []int{}

	for _, v := range seedsstr {
		seeds = append(seeds, aoc.ParseNumOrPanic(v))
	}

	mapBlocks := lines[1:]

	for _, block := range mapBlocks {
		mapname, entries := strings.Fields(block[0])[0], block[1:]
		makeFarmMap(mapname, entries)
	}

	// fmt.Println(Name2Map)
	// fmt.Println(seeds)

	r1, r2 := p1(seeds), p2(seeds)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}
