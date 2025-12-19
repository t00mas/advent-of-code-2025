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

func ParseInput() []Problem {
	var input []byte
	if *Example {
		input = example
	} else {
		input = data
	}

	lines := strings.Split(string(input), "\n")
	split_rows := make([][]string, len(lines))
	for i := range lines {
		split_rows[i] = strings.Fields(lines[i])
	}

	problems := make([]Problem, len(split_rows[0]))
	for i := range split_rows[0] {
		for j := 0; j < len(split_rows)-1; j++ {
			val, err := strconv.Atoi(split_rows[j][i])
			if err != nil {
				panic(err)
			}
			problems[i].Values = append(problems[i].Values, val)
		}
		problems[i].Operator = split_rows[len(split_rows)-1][i]
	}

	return problems
}

func ParseInputCephalopod() []Problem {
	input := data
	if *Example {
		input = example
	}

	L := strings.Split(string(input), "\n")
	w := 0
	for _, l := range L {
		if len(l) > w {
			w = len(l)
		}
	}
	for i := range L {
		L[i] += strings.Repeat(" ", w-len(L[i]))
	}

	var probs []Problem
	var cols []int
	for c := w - 1; c >= -1; c-- {
		space := c < 0
		if !space {
			space = true
			for _, l := range L {
				if l[c] != ' ' {
					space = false
					break
				}
			}
		}
		if space && len(cols) > 0 {
			var vals []int
			var op byte
			for _, col := range cols {
				n := 0
				for r, l := range L {
					if r == len(L)-1 {
						if l[col] != ' ' {
							op = l[col]
						}
					} else if l[col] >= '0' && l[col] <= '9' {
						n = n*10 + int(l[col]-'0')
					}
				}
				vals = append(vals, n)
			}
			probs = append(probs, Problem{Values: vals, Operator: string(op)})
			cols = nil
		} else if !space {
			cols = append(cols, c)
		}
	}
	return probs
}
