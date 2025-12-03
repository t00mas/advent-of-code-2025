package main

import (
	_ "embed"
	"strings"
)

//go:embed example.txt
var example []byte

//go:embed data.txt
var data []byte

// ParseInput parses the input and returns a list of Ranges
func ParseInput() []Range {
	var input []byte
	if *Example {
		input = example
	} else {
		input = data
	}

	ranges := make([]Range, 0)
	for r := range strings.SplitSeq(string(input), ",") {
		min, max, _ := strings.Cut(r, "-")
		ranges = append(ranges, *NewRange(min, max))
	}
	return ranges
}
