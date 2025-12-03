package main

import (
	_ "embed"
	"strings"
)

//go:embed example.txt
var example []byte

//go:embed data.txt
var data []byte

// ParseInput parses the input and returns a list of BatteryBanks
func ParseInput() []*BatteryBank {
	var input []byte
	if *Example {
		input = example
	} else {
		input = data
	}
	lines := strings.Split(string(input), "\n")
	banks := make([]*BatteryBank, len(lines))
	for n, line := range lines {
		banks[n] = NewBatteryBank(line)
	}
	return banks
}
