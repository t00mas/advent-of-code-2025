package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
)

var Debug = flag.Bool("debug", false, "Debug")
var Example = flag.Bool("example", false, "Use example input")

func PartOne(ranges []Range) int {
	sum := 0
	for _, r := range ranges {
		mss := r._2Sequences()
		for _, ms := range mss {
			ms_i, _ := strconv.Atoi(ms)
			sum += ms_i
		}
	}
	return sum
}

func PartTwo(ranges []Range) int {
	sum := 0
	for _, r := range ranges {
		nss := r._NSequences()
		for _, ns := range nss {
			ns_i, _ := strconv.Atoi(ns)
			sum += ns_i
		}
	}
	return sum
}

func main() {
	flag.Parse()

	ranges := ParseInput()

	fmt.Println("Part 1: ")
	fmt.Println(PartOne(ranges))

	fmt.Println("Part 2: ")
	fmt.Println(PartTwo(ranges))
}
