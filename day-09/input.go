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

func ParseInput() [][2]int {
	var input []byte
	if *Example {
		input = example
	} else {
		input = data
	}
	lines := strings.Split(string(input), "\n")
	points := make([][2]int, len(lines))
	for i, line := range lines {
		pos_values := strings.Split(line, ",")
		points[i][0], _ = strconv.Atoi(pos_values[0])
		points[i][1], _ = strconv.Atoi(pos_values[1])
	}
	return points
}
