package main

import "strconv"

type Instruction struct {
	Direction string
	Steps     int
}

func ParseInstructions(instruction []string) []Instruction {
	instructions := make([]Instruction, len(instruction))
	for n, i := range instruction {
		direction := string(i[0])
		steps, _ := strconv.Atoi(i[1:])
		instructions[n] = Instruction{Direction: direction, Steps: steps}
	}
	return instructions
}
