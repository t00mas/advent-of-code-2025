package main

import "slices"

type IngredientDB struct {
	FreshRanges []Range
}

func NewIngredientDB() *IngredientDB {
	return &IngredientDB{
		FreshRanges: make([]Range, 0),
	}
}

func (idb *IngredientDB) AddFreshRange(min, max int) {
	idb.FreshRanges = append(idb.FreshRanges, NewRange(min, max))
}

func (idb *IngredientDB) IsFresh(n int) bool {
	if len(idb.FreshRanges) == 0 {
		return false
	}
	idx, _ := slices.BinarySearchFunc(idb.FreshRanges, n, func(r Range, val int) int {
		return r.Min - val
	})
	if idx > 0 {
		idx--
	}
	if idx < len(idb.FreshRanges) && idb.FreshRanges[idx].Min <= n && n <= idb.FreshRanges[idx].Max {
		return true
	}
	return false
}

func (idb *IngredientDB) MergeFreshRanges() {
	if len(idb.FreshRanges) == 0 {
		return
	}
	slices.SortFunc(idb.FreshRanges, func(a, b Range) int {
		return a.Min - b.Min
	})
	merged := []Range{idb.FreshRanges[0]}
	for i := 1; i < len(idb.FreshRanges); i++ {
		last := &merged[len(merged)-1]
		if last.Max >= idb.FreshRanges[i].Min {
			if idb.FreshRanges[i].Max > last.Max {
				last.Max = idb.FreshRanges[i].Max
			}
		} else {
			merged = append(merged, idb.FreshRanges[i])
		}
	}
	idb.FreshRanges = merged
}

func (idb *IngredientDB) GetTotalFreshRangeCoverage() int {
	coverage := 0
	for _, r := range idb.FreshRanges {
		coverage += r.Max - r.Min + 1
	}
	return coverage
}

type Range struct {
	Min int
	Max int
}

func NewRange(min, max int) Range {
	return Range{Min: min, Max: max}
}

func (r *Range) Contains(n int) bool {
	return n >= r.Min && n <= r.Max
}
