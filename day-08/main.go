package main

import (
	_ "embed"
	"flag"
	"fmt"
)

var Debug = flag.Bool("debug", false, "Debug")
var Example = flag.Bool("example", false, "Use example input")

func PartOne(input []JunctionBox) int {
	pairs := AllPairsByDistance(input)
	uf := NewUnionFind(len(input))

	targetPairs := 1000
	if *Example {
		targetPairs = 10
	}

	for i := 0; i < targetPairs && i < len(pairs); i++ {
		p := pairs[i]
		connected := uf.Union(p.I, p.J)
		if *Debug {
			status := "connects"
			if !connected {
				status = "skipped"
			}
			fmt.Printf("Pair %d: %v - %v (dist %.2f) - %s\n", i+1, input[p.I], input[p.J], p.Dist, status)
		}
	}

	sizes := uf.CircuitSizes()
	if *Debug {
		fmt.Println("Circuit sizes:", sizes)
	}
	return sizes[0] * sizes[1] * sizes[2]
}

func PartTwo(input []JunctionBox) int {
	pairs := AllPairsByDistance(input)
	uf := NewUnionFind(len(input))

	circuitsRemaining := len(input)
	for _, p := range pairs {
		if uf.Union(p.I, p.J) {
			circuitsRemaining--
			if circuitsRemaining == 1 {
				return input[p.I].PosX * input[p.J].PosX
			}
		}
	}
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
