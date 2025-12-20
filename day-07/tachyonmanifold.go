package main

import "bytes"

type TachyonManifold struct {
	Lines [][]byte
}

func NewTachyonManifold() *TachyonManifold {
	return &TachyonManifold{Lines: make([][]byte, 0)}
}

func (t *TachyonManifold) AddLine(line string) {
	t.Lines = append(t.Lines, []byte(line))
}

type BeamSimulation struct {
	CurrentStep   int
	TachyonSplits int
}

func NewBeamSimulation() *BeamSimulation {
	return &BeamSimulation{CurrentStep: 0, TachyonSplits: 0}
}

func (b *BeamSimulation) SimulateBeamStep(t *TachyonManifold) {
	// step is the line index
	current_line := t.Lines[b.CurrentStep]
	for i := range current_line {
		if current_line[i] == 'S' {
			// beam entry point, cell down from it is a beam now
			t.Lines[b.CurrentStep+1][i] = '|'
		}
		if current_line[i] == '|' {
			// if next cell down is empty space, continue beam
			if t.Lines[b.CurrentStep+1][i] == '.' {
				t.Lines[b.CurrentStep+1][i] = '|'
			}
			// if next cell down is a splitter, left/right of cell down is a beam now
			if t.Lines[b.CurrentStep+1][i] == '^' {
				t.Lines[b.CurrentStep+1][i-1] = '|'
				t.Lines[b.CurrentStep+1][i+1] = '|'
				b.TachyonSplits++
			}
		}
	}
}

type QuantumBeamSimulation struct {
	Timelines int
}

func NewQuantumBeamSimulation() *QuantumBeamSimulation {
	return &QuantumBeamSimulation{Timelines: 0}
}

func (q *QuantumBeamSimulation) CalculateTimelines(t *TachyonManifold) int {
	timelines := map[int]int{bytes.IndexByte(t.Lines[0], 'S'): 1}

	for row := 1; row < len(t.Lines); row++ {
		next := make(map[int]int)
		for col, count := range timelines {
			if t.Lines[row][col] == '^' {
				next[col-1] += count
				next[col+1] += count
			} else {
				next[col] += count
			}
		}
		timelines = next
	}

	total := 0
	for _, c := range timelines {
		total += c
	}
	return total
}
