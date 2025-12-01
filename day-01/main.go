package main

import (
	"flag"
	"fmt"
)

var Debug = flag.Bool("debug", false, "Debug")
var Example = flag.Bool("example", false, "Use example input")

func PartOne(instructions []Instruction) int {
	cur, zeroes := 50, 0
	for _, i := range instructions {
		cur = Rotate(cur, i)
		if cur == 0 {
			zeroes++
		}
	}
	return zeroes
}

func PartTwo(instructions []Instruction) int {
	cur, zeroes := 50, 0
	for _, i := range instructions {
		c, z := RotateCountingZeroes(cur, i)
		cur = c
		zeroes += z
	}
	return zeroes
}

func main() {
	flag.Parse()

	parsedInput := ParseInput()
	instructions := ParseInstructions(parsedInput)

	fmt.Println("Part 1: ")
	fmt.Println(PartOne(instructions))

	fmt.Println("Part 2: ")
	fmt.Println(PartTwo(instructions))
}
