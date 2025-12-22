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

func ParseInput() []JunctionBox {
	var input []byte
	if *Example {
		input = example
	} else {
		input = data
	}
	lines := strings.Split(string(input), "\n")
	boxes := make([]JunctionBox, len(lines))
	for i, line := range lines {
		pos_values := strings.Split(line, ",")
		boxes[i].PosX, _ = strconv.Atoi(pos_values[0])
		boxes[i].PosY, _ = strconv.Atoi(pos_values[1])
		boxes[i].PosZ, _ = strconv.Atoi(pos_values[2])
	}
	return boxes
}
