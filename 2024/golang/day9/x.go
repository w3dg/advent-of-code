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

func calculateChecksum(diskMap []int) int {
	sum := 0
	for i, v := range diskMap {
		if v == -1 {
			continue
		}
		sum += i * v
	}

	return sum
}

// func calculateChecksumAlt(altDiskMap []FileBlock) int {
// 	sum := 0
// 	i := 0

// 	for _, block := range altDiskMap {
// 		if block.isFree {
// 			i += block.len
// 		} else {
// 			for range block.len {
// 				sum += block.index * i
// 				i++
// 			}
// 		}
// 	}

// 	return sum
// }

func p1(diskMap []int) int {
	l, r := 0, len(diskMap)-1

	for l < r {
		// find free space on left
		for diskMap[l] != -1 {
			l++
		}
		// find a block from the right
		for diskMap[r] == -1 {
			r--
		}
		// if the invariant still holds, move the block over from right to left
		if l < r {
			diskMap[l] = diskMap[r]
			diskMap[r] = -1
			l++
			r--
		}
	}

	return calculateChecksum(diskMap)
}

type FileBlock struct {
	len         int
	blockNumber int
	startIndex  int
	endIndex    int
}

func parseAltDiskMap(diskMap []int) []FileBlock {
	altDiskMap := []FileBlock{}
	blockCounter := 0
	for i := 0; i < len(diskMap); {
		for i < len(diskMap) && diskMap[i] == -1 {
			i++
		}
		last := diskMap[i]
		blockStart := i
		blockLen := 0
		j := i
		for ; j < len(diskMap) && diskMap[j] == last; j++ {
			blockLen++
		}
		i = j
		altDiskMap = append(altDiskMap, FileBlock{blockLen, blockCounter, blockStart, blockStart + blockLen - 1})
		blockCounter++
	}
	return altDiskMap
}

func p2(altDiskMap []FileBlock) int {
	// newDiskMap := []FileBlock{}
	// l, r := 0, len(altDiskMap)-1

	// for r > 0 && l < r {
	// 	// rightBlock := altDiskMap[r]

	// }

	// if we get some space consecutive in between blocks endindex of one block and start index of next block,
	// // we insert a block after the block with endIndex+1 to endIndex+len
	// slices.Insert()
	// slices.Delete() the end block from [endblockindex:endblockindex+1]
	// fmt.Println(altDiskMap)
	// PrettyPrintAltDisk(altDiskMap)
	return 0
	// return calculateChecksumAlt(newDiskMap)
}

func PrettyPrintDisk(diskMap []int) {
	for _, v := range diskMap {
		if v == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(v)
		}
	}
	fmt.Println()
}

func PrettyPrintAltDisk(altDiskMap []FileBlock) {
	for i := 1; i < len(altDiskMap); i++ {
		currBlock, nextBlock := altDiskMap[i-1], altDiskMap[i]

		// print current block
		for range currBlock.len {
			fmt.Print(currBlock.blockNumber)
		}
		// see if space is there between this block and the next block
		if currBlock.endIndex+1 != nextBlock.startIndex {
			for range nextBlock.startIndex - currBlock.endIndex - 1 {
				fmt.Print(".")
			}
		}
		// advance forward
	}
	// print the last block if there is
	if len(altDiskMap)-1 >= 0 {
		lastBlock := altDiskMap[len(altDiskMap)-1]
		for range lastBlock.len {
			fmt.Print(lastBlock.blockNumber)
		}
	}
	// if theres any free space ahead after compaction, we dont care, its gonna be multiplied as 0 anyway
	fmt.Println()
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

	diskMapStructure := lines[0]
	diskMap := []int{}

	isFile := true
	idx := 0
	for v := range strings.SplitSeq(diskMapStructure, "") {
		n := num(v)
		if isFile {
			for range n {
				diskMap = append(diskMap, idx)
			}
			idx++
		} else {
			for range n {
				diskMap = append(diskMap, -1)
			}
		}
		isFile = !isFile
	}

	// PrettyPrintDisk(diskMap)
	altDiskMap := parseAltDiskMap(diskMap)
	r1, r2 := p1(diskMap), p2(altDiskMap)
	fmt.Println("Puzzle 1:", r1, "Puzzle 2:", r2)
}
