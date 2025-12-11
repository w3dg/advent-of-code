package main

import (
	"fmt"
	"strings"

	aoc "github.com/w3dg/aocutils"
)

const (
	CD_COMMAND string = "cd"
	LS_COMMAND string = "ls"
)

type CommandOutput struct {
	Command  string
	Argument *string   // ls command does not have any args, cd has
	Output   *[]string // cd commands have no output, ls have them
}

func (c CommandOutput) String() string {
	s := fmt.Sprint(c.Command)

	if c.Command == CD_COMMAND {
		s += fmt.Sprintln(" ", "TARGET_DIR:", *c.Argument)
	}

	if c.Command == LS_COMMAND {
		s += fmt.Sprintln(" output:")
		for _, outputline := range *(c.Output) {
			s += fmt.Sprintln(outputline)
		}
	}

	return s
}

func ParseIntoCommands(lines []string) []CommandOutput {
	commands := []CommandOutput{}

	currentCmd := ""
	l := len(lines)
	for i := 0; i < l; i++ {
		if after, ok := strings.CutPrefix(lines[i], "$ "); ok {
			// new command
			currentCmd = after
			parts := strings.Split(currentCmd, " ")
			currentCmd = parts[0]
			if currentCmd == CD_COMMAND {
				arg := parts[1]
				commands = append(commands, CommandOutput{
					Command:  currentCmd,
					Argument: &arg,
					Output:   nil,
				})
			} else if currentCmd == LS_COMMAND {
				// also go throuugh the ls output
				output := []string{}
				for k := i + 1; k < l; k++ {
					// reached another command, bail out and let outer parser parse
					// another reason can be the end of the output ends with ls output
					if k == l-1 || strings.HasPrefix(lines[k], "$ ") {
						// handle the end of input carefully
						if k == l-1 {
							output = append(output, lines[k])
						}
						commands = append(commands, CommandOutput{
							Command:  currentCmd,
							Argument: nil,
							Output:   &output,
						})
						i = k - 1 // as it will update i to i++ next
						break
					}
					output = append(output, lines[k])
				}
			} else {
				panic("UNKNOWN COMMAND" + currentCmd)
			}
		}
	}

	return commands
}

func MakeFSFromCommands(commands []CommandOutput) *Node {
	r := MakeRootDir()
	curr := r

	commands = commands[1:]

	for _, c := range commands {
		switch c.Command {
		case LS_COMMAND:
			if c.Output == nil {
				panic("LS COMMAND HAS NIL OUTPUT INSTEAD OF EMPTY")
			}
			output := *c.Output
			for _, l := range output {
				parts := strings.Split(l, " ")
				if parts[0] == "dir" {
					curr.MakeChildDir(parts[1])
				} else {
					filesize := parts[0]
					filename := parts[1]
					curr.MakeChildFile(filename, aoc.ParseNumOrPanic(filesize))
				}
			}
		case CD_COMMAND:
			curr = curr.MoveIntoDir(*c.Argument)
		default:
			panic("UNKNOWN COMMAND " + c.String())
		}
	}

	return r
}
