package main

import "testing"

func TestCountAdjacentPaperRollsWithMax(t *testing.T) {
	tests := []struct {
		rawGrid [][]string
		want    int
	}{
		{[][]string{
			{".", ".", "."},
			{".", "@", "."},
			{".", ".", "."},
		}, 0},
		{[][]string{
			{".", "@", "@"},
			{".", "@", "@"},
			{".", "@", "@"},
		}, 5},
		{[][]string{
			{".", "@", "@"},
			{".", "@", "@"},
			{"@", "@", "@"},
		}, 5}, // max is at 4, stops counting at 5 and returns 5
	}

	for _, test := range tests {
		grid := NewGrid(len(test.rawGrid[0]), len(test.rawGrid))
		for y, row := range test.rawGrid {
			for x, cell := range row {
				if cell == "@" {
					grid.SetPaperRoll(x, y)
				}
			}
		}

		got := grid.CountAdjacentPaperRollsWithMax(1, 1, 4)
		if got != test.want {
			t.Errorf("CountAdjacentPaperRollsWithMax(%v, 1, 1, 4) = %d, want %d", test.rawGrid, got, test.want)
		}
	}
}
