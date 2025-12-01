package main

import "testing"

func TestRotate(t *testing.T) {
	tests := []struct {
		starts      int
		instruction Instruction
		want        int
	}{
		{0, Instruction{Direction: "L", Steps: 100}, 0},
		{0, Instruction{Direction: "R", Steps: 100}, 0},
		{1, Instruction{Direction: "L", Steps: 1}, 0},
		{1, Instruction{Direction: "R", Steps: 1}, 2},
		{8, Instruction{Direction: "L", Steps: 782}, 26},
		{8, Instruction{Direction: "R", Steps: 782}, 90},
	}

	for _, test := range tests {
		got := Rotate(test.starts, test.instruction)
		if got != test.want {
			t.Errorf("rotate(%d, %+v) = %d, want %d", test.starts, test.instruction, got, test.want)
		}
	}
}

func TestRotateCountingZeroes(t *testing.T) {
	tests := []struct {
		starts      int
		instruction Instruction
		zeroes      int
	}{
		{0, Instruction{Direction: "L", Steps: 100}, 1},
		{0, Instruction{Direction: "R", Steps: 100}, 1},
		{1, Instruction{Direction: "L", Steps: 1}, 1},
		{1, Instruction{Direction: "R", Steps: 1}, 0},
		{8, Instruction{Direction: "L", Steps: 782}, 8},
		{8, Instruction{Direction: "R", Steps: 782}, 7},
	}

	for _, test := range tests {
		_, got := RotateCountingZeroes(test.starts, test.instruction)
		if got != test.zeroes {
			t.Errorf("rotateCountingZeroes(%d, %+v) = %d, want %d", test.starts, test.instruction, got, test.zeroes)
		}
	}
}
