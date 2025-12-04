package main

import "strings"

type Grid struct {
	Width  int
	Height int
	Grid   [][]int
}

// String returns a beautiful string representation of the grid for debugging purposes
func (g *Grid) String() string {
	var sb strings.Builder
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if g.HasPaperRoll(x, y) {
				sb.WriteString("@")
			} else {
				sb.WriteString(".")
			}
		}
		if y < g.Height-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

// NewGrid creates a new grid with the given width and height
func NewGrid(width, height int) *Grid {
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}
	return &Grid{Width: width, Height: height, Grid: grid}
}

// SetPaperRoll sets a paper roll at the given position
func (g *Grid) SetPaperRoll(x, y int) {
	g.Grid[y][x] = 1
}

// HasPaperRoll returns true if the given position has a paper roll
func (g *Grid) HasPaperRoll(x, y int) bool {
	return g.Grid[y][x] == 1
}

// CountAdjacentPaperRolls counts the exactnumber of adjacent paper rolls
func (g *Grid) CountAdjacentPaperRolls(x, y int) int {
	// 8 is the maximum number of adjacent paper rolls in a grid anyway
	return g.CountAdjacentPaperRollsWithMax(x, y, 8)
}

// CountAdjacentPaperRollsWithMax counts the number of adjacent paper rolls up to a given maximum count
// Use it to early-stop counting when the count is greater than the maximum given
func (g *Grid) CountAdjacentPaperRollsWithMax(x, y, max int) int {
	count := 0

	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			if dx == 0 && dy == 0 {
				continue
			}

			if x+dx < 0 || x+dx >= g.Width || y+dy < 0 || y+dy >= g.Height {
				continue
			}

			if g.HasPaperRoll(x+dx, y+dy) {
				count++
			}

			if count > max {
				return count
			}
		}
	}

	return count
}

// IsPaperRollAccessible returns true if the paper roll is accessible from the given position
// It's accessible if it is a paper roll and there are fewer than 4 adjacent paper rolls
func (g *Grid) IsPaperRollAccessible(x, y int) bool {
	return g.HasPaperRoll(x, y) && g.CountAdjacentPaperRollsWithMax(x, y, 3) < 4
}

// RemovePaperRoll removes a paper roll from the given position
func (g *Grid) RemovePaperRoll(x, y int) {
	g.Grid[y][x] = 0
}
