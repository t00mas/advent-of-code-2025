package main

import (
	_ "embed"
	"flag"
	"fmt"
)

var Debug = flag.Bool("debug", false, "Debug")
var Example = flag.Bool("example", false, "Use example input")

func PartOne(db *IngredientDB, ingredients []int) int {
	freshIngredients := 0
	for _, ingredient := range ingredients {
		if db.IsFresh(ingredient) {
			freshIngredients++
		}
	}
	return freshIngredients
}

func PartTwo(db *IngredientDB, ingredients []int) int {
	return db.GetTotalFreshRangeCoverage()
}

func main() {
	flag.Parse()

	db, ingredients := ParseInput()

	fmt.Println("Part 1: ")
	fmt.Println(PartOne(db, ingredients))

	fmt.Println("Part 2: ")
	fmt.Println(PartTwo(db, ingredients))
}
