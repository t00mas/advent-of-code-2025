package main

import (
	_ "embed"
	"flag"
	"fmt"
)

var Debug = flag.Bool("debug", false, "Debug")
var Example = flag.Bool("example", false, "Use example input")

func PartOne(input *TachyonManifold) int {
	simulation := NewBeamSimulation()
	for i := 0; i < len(input.Lines)-1; i++ {
		simulation.CurrentStep = i
		simulation.SimulateBeamStep(input)
	}
	return simulation.TachyonSplits
}

func PartTwo(input *TachyonManifold) int {
	simulation := NewQuantumBeamSimulation()
	return simulation.CalculateTimelines(input)
}

func main() {
	flag.Parse()

	parsedInput := ParseInput()

	fmt.Println("Part 1: ")
	fmt.Println(PartOne(parsedInput))

	fmt.Println("Part 2: ")
	fmt.Println(PartTwo(parsedInput))
}
