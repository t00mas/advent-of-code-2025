package main

import (
	"slices"
	"strconv"
)

// Range represents a range of numbers
type Range struct {
	Min   string
	Max   string
	min_i int
	max_i int
}

// NewRange creates a new range
func NewRange(min, max string) *Range {
	r := Range{Min: min, Max: max}
	r.min_i, _ = strconv.Atoi(r.Min)
	r.max_i, _ = strconv.Atoi(r.Max)
	return &r
}

// Valid returns true if the range is valid
func (r *Range) Valid() bool {
	return r.min_i <= r.max_i
}

// Length returns the length of the range
func (r *Range) Length() int {
	return r.max_i - r.min_i + 1
}

// InRange returns true if the number is in the range
func (r *Range) InRange(num int) bool {
	return num >= r.min_i && num <= r.max_i
}

// _2Sequences returns all digit sequences that are 2-sequenced in the given range
func (r *Range) _2Sequences() []string {
	sequences := make([]string, 0)

	for num := r.min_i; num <= r.max_i; num++ {
		numStr := strconv.Itoa(num)
		if isNSequenced(numStr, 2) {
			sequences = append(sequences, numStr)
		}
	}
	return sequences
}

// _NSequences returns all digit sequences that are n-sequenced in the given range
func (r *Range) _NSequences() []string {
	seen := make(map[string]bool) // keep track of seen sequences to avoid duplicates
	sequences := make([]string, 0)

	for num := r.min_i; num <= r.max_i; num++ {
		numStr := strconv.Itoa(num)
		for n := 2; n <= len(numStr); n++ {
			if isNSequenced(numStr, n) && !seen[numStr] {
				seen[numStr] = true
				sequences = append(sequences, numStr)
				break
			}
		}
	}

	slices.SortFunc(sequences, func(a, b string) int {
		na, _ := strconv.Atoi(a)
		nb, _ := strconv.Atoi(b)
		return na - nb
	})

	return sequences
}

// isNSequenced returns true if the string is a sequence of digits repeated n times
func isNSequenced(s string, n int) bool {
	if len(s)%n != 0 { // s must be divisible in n equal chunks
		return false
	}

	seq_len := len(s) / n
	for i := 0; i < n-1; i++ {
		prev_chunk := s[i*seq_len : (i+1)*seq_len]
		next_chunk := s[(i+1)*seq_len : (i+2)*seq_len]
		if prev_chunk != next_chunk {
			return false
		}
	}

	return true
}
