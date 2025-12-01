package main

import (
	_ "embed"
	"flag"
	"fmt"
)

var Debug = flag.Bool("debug", false, "Debug")
var Example = flag.Bool("example", false, "Use example input")

func PartOne(input []string) int {
	return 0
}

func PartTwo(input []string) int {
	return 0
}

func main() {
	flag.Parse()

	parsedInput := ParseInput()

	fmt.Println("Part 1: ")
	fmt.Println(PartOne(parsedInput))

	fmt.Println("Part 2: ")
	fmt.Println(PartTwo(parsedInput))
}
