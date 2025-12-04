package main

import (
	_ "embed"
	"strings"
)

//go:embed example.txt
var example []byte

//go:embed data.txt
var data []byte

func ParseInput() *Grid {
	var input []byte

	if *Example {
		input = example
	} else {
		input = data
	}

	rows := strings.Split(string(input), "\n")
	grid := NewGrid(len(rows[0]), len(rows))
	for y, row := range rows {
		for x, cell := range row {
			if cell == '@' {
				grid.SetPaperRoll(x, y)
			}
		}
	}

	return grid
}
