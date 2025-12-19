package main

import (
	_ "embed"
	"flag"
	"fmt"
)

var Debug = flag.Bool("debug", false, "Debug")
var Example = flag.Bool("example", false, "Use example input")

func PartOne(input []Problem) int {
	sum := 0
	for _, problem := range input {
		sum += Solve(problem)
	}
	return sum
}

func PartTwo(input []Problem) int {
	sum := 0
	for _, problem := range input {
		sum += Solve(problem)
	}
	return sum
}

func main() {
	flag.Parse()

	parsedInput := ParseInput()
	fmt.Println("Part 1: ")
	fmt.Println(PartOne(parsedInput))

	parsedInputCephalopod := ParseInputCephalopod()
	fmt.Println("Part 2: ")
	fmt.Println(PartTwo(parsedInputCephalopod))
}
