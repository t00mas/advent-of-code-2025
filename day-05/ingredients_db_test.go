package main

import (
	"reflect"
	"testing"
)

func TestMergeFreshRanges(t *testing.T) {
	tests := []struct {
		testRanges []Range
		wantRanges []Range
	}{
		{
			[]Range{
				NewRange(1, 5),
				NewRange(10, 14),
				NewRange(16, 20),
				NewRange(12, 18),
			},
			[]Range{
				NewRange(1, 5),
				NewRange(10, 20),
			},
		},
	}

	for _, test := range tests {
		db := IngredientDB{}
		for _, r := range test.testRanges {
			db.AddFreshRange(r.Min, r.Max)
		}
		db.MergeFreshRanges()
		if !reflect.DeepEqual(db.FreshRanges, test.wantRanges) {
			t.Errorf("MergeFreshRanges(%v) = %v, want %v", test.testRanges, db.FreshRanges, test.wantRanges)
		}
	}
}
