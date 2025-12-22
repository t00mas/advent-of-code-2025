package main

import (
	"math"
	"sort"
)

type JunctionBox struct {
	PosX int
	PosY int
	PosZ int
}

func ComputeDistance(box1, box2 JunctionBox) float64 {
	dx := float64(box1.PosX - box2.PosX)
	dy := float64(box1.PosY - box2.PosY)
	dz := float64(box1.PosZ - box2.PosZ)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

type Pair struct {
	I, J int
	Dist float64
}

func AllPairsByDistance(boxes []JunctionBox) []Pair {
	var pairs []Pair
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			pairs = append(pairs, Pair{I: i, J: j, Dist: ComputeDistance(boxes[i], boxes[j])})
		}
	}
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].Dist < pairs[b].Dist
	})
	return pairs
}

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := range uf.parent {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	px, py := uf.Find(x), uf.Find(y)
	if px == py {
		return false
	}
	if uf.rank[px] < uf.rank[py] {
		px, py = py, px
	}
	uf.parent[py] = px
	if uf.rank[px] == uf.rank[py] {
		uf.rank[px]++
	}
	return true
}

func (uf *UnionFind) CircuitSizes() []int {
	counts := make(map[int]int)
	for i := range uf.parent {
		root := uf.Find(i)
		counts[root]++
	}
	sizes := make([]int, 0, len(counts))
	for _, c := range counts {
		sizes = append(sizes, c)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return sizes
}
