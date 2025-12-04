package main

import (
	_ "embed"
	"flag"
	"fmt"
)

var Debug = flag.Bool("debug", false, "Debug")
var Example = flag.Bool("example", false, "Use example input")

func PartOne(input *Grid) int {
	accessiblePaperRolls := 0
	for y := 0; y < input.Height; y++ {
		for x := 0; x < input.Width; x++ {
			if input.IsPaperRollAccessible(x, y) {
				accessiblePaperRolls++
			}
		}
	}

	return accessiblePaperRolls
}

func PartTwo(input *Grid) int {
	removed := 0
	for {
		removed_this_pass := 0
		for y := 0; y < input.Height; y++ {
			for x := 0; x < input.Width; x++ {
				if input.IsPaperRollAccessible(x, y) {
					input.RemovePaperRoll(x, y)
					removed++
					removed_this_pass++
				}
			}
		}

		if removed_this_pass == 0 {
			break
		}
	}

	return removed
}

func main() {
	flag.Parse()

	parsedInput := ParseInput()

	fmt.Println("Part 1: ")
	fmt.Println(PartOne(parsedInput))

	fmt.Println("Part 2: ")
	fmt.Println(PartTwo(parsedInput))
}
