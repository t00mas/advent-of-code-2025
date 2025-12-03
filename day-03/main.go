package main

import (
	_ "embed"
	"flag"
	"fmt"
)

var Debug = flag.Bool("debug", false, "Debug")
var Example = flag.Bool("example", false, "Use example input")
var Explain = flag.Bool("explain", false, "Explain the solution")
var GenerateGIF = flag.Bool("generate-gif", false, "Generate VHS tape file for GIF creation")

func PartOne(input []*BatteryBank) int {
	sum := 0
	for _, bank := range input {
		sum += bank.MaxJoltage()
	}
	return sum
}

func PartTwo(input []*BatteryBank) int {
	sum := 0
	for _, bank := range input {
		sum += bank.SuperJoltage()
	}
	return sum
}

func main() {
	flag.Parse()

	BatteryBanks := ParseInput()

	if *GenerateGIF {
		if len(BatteryBanks) == 0 {
			fmt.Println("No battery banks found")
			return
		}
		err := GenerateVHSTape(BatteryBanks[0].Batteries, 12, "explanation.gif")
		if err != nil {
			fmt.Printf("Error generating VHS tape: %v\n", err)
		}
		return
	}

	if *Explain {
		SuperJoltageExplained(BatteryBanks[0].Batteries, 12)
		return
	}

	fmt.Println("Part 1: ")
	fmt.Println(PartOne(BatteryBanks))

	fmt.Println("Part 2: ")
	fmt.Println(PartTwo(BatteryBanks))
}
