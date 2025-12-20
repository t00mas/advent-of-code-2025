package main

import (
	_ "embed"
	"strings"
)

//go:embed example.txt
var example []byte

//go:embed data.txt
var data []byte

func ParseInput() *TachyonManifold {
	var input []byte
	if *Example {
		input = example
	} else {
		input = data
	}

	lines := strings.Split(string(input), "\n")

	manifold := NewTachyonManifold()
	for _, line := range lines {
		manifold.AddLine(line)
	}
	return manifold
}
