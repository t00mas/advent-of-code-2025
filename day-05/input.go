package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed example.txt
var example []byte

//go:embed data.txt
var data []byte

func ParseInput() (*IngredientDB, []int) {
	var input []byte
	if *Example {
		input = example
	} else {
		input = data
	}
	lines := strings.Split(string(input), "\n")
	db := NewIngredientDB()
	ingredients := make([]int, 0)
	for _, line := range lines {
		if strings.Contains(line, "-") {
			min, max, _ := strings.Cut(line, "-")
			minInt, _ := strconv.Atoi(min)
			maxInt, _ := strconv.Atoi(max)
			db.AddFreshRange(minInt, maxInt)
		} else {
			ingredient, _ := strconv.Atoi(line)
			ingredients = append(ingredients, ingredient)
		}
	}
	db.MergeFreshRanges()
	return db, ingredients
}
