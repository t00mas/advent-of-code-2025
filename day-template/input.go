package main

import (
	_ "embed"
	"strings"
)

//go:embed example.txt
var example []byte

//go:embed data.txt
var data []byte

func ParseInput() []string {
	var input []byte
	if *Example {
		input = example
	} else {
		input = data
	}
	return strings.Split(string(input), "\n")
}
