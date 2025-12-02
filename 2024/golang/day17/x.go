package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	aoc "github.com/w3dg/aocutils"
	_ "github.com/zyedidia/generic"
)

// https://en.wiktionary.org/wiki/quine

var registerA, registerB, registerC int

func getValueFromComboOperand(operand int) int {
	switch operand {
	case 7:
		log.Fatal("Combo Operand 7!")
	case 0, 1, 2, 3:
		return operand
	case 4:
		return registerA
	case 5:
		return registerB
	case 6:
		return registerC
	}

	return 0
}

func performOperation(opcode, op int) (int, bool) {
	comboOperand := getValueFromComboOperand(op)
	switch opcode {
	case 0:
		// adv
		// numerator := registerA
		// denominator := int(math.Pow(2, float64(comboOperand)))
		// registerA = numerator / denominator
		// note: this is the same as right shifting by comboOperand number of places
		registerA = registerA >> comboOperand
	case 1:
		// bxl
		registerB = registerB ^ op
	case 2:
		// bst
		registerB = comboOperand % 8
	case 3:
		// jnz
		if registerA != 0 {
			// set ip to be op
			// do not inc ip by 2 after jump!
			return op, true // returns next ip and whether it is jump
		}
	case 4:
		// bxc
		registerB = registerB ^ registerC
	case 5:
		// out
		mod8 := comboOperand % 8
		return mod8, false // return value as out and it was not a jump
	case 6:
		// bdv
		// numerator := registerA
		// denominator := int(math.Pow(2, float64(comboOperand)))
		// registerB = numerator / denominator
		registerB = registerA >> comboOperand
	case 7:
		// cdv
		// numerator := registerA
		// denominator := int(math.Pow(2, float64(comboOperand)))
		// registerC = numerator / denominator
		registerC = registerA >> comboOperand
	}
	return -1, false // if -1 is returned, its not to be appended to out
}

func p1(initA, initB, initC int, program []string) string {
	registerA = initA
	registerB = initB
	registerC = initC

	// Instruction pointer starts at index 0 for the program string input
	// increases by 2 normally except jumps
	// line 1: opcode
	// line 2: operand
	// When ip reaches outside of the program string length, it halts
	ip := 0
	output := ""

	for ip < len(program) {
		opcode := aoc.ParseNumOrPanic(program[ip])
		operand := aoc.ParseNumOrPanic(program[ip+1])
		// fmt.Println(opcode, operand)
		val, wasJump := performOperation(opcode, operand)
		if wasJump {
			ip = val // jump ip to the operand
		} else {
			if val != -1 {
				output += fmt.Sprintf("%d,", val)
			}
			ip += 2
		}
	}

	return strings.TrimSuffix(output, ",")
}

func main() {
	infile := "./input.txt"
	arg := os.Args[1:]
	if len(arg) != 0 {
		infile = arg[0]
	}
	lines, err := aoc.ReadFileLines(infile)
	if err != nil {
		log.Fatal("cannot read")
	}

	ra := aoc.ParseNumOrPanic(strings.Split(lines[0], ": ")[1])
	rb := aoc.ParseNumOrPanic(strings.Split(lines[1], ": ")[1])
	rc := aoc.ParseNumOrPanic(strings.Split(lines[2], ": ")[1])

	programstr := strings.Split(lines[4], ": ")[1]
	program := strings.Split(programstr, ",")

	r1 := p1(ra, rb, rc, program)
	fmt.Println("part 1:", r1)
}

func p2(programstr string) {

}
