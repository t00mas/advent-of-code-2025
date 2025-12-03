package main

import (
	"reflect"
	"testing"
)

func TestSequencesRepeatedOnce(t *testing.T) {
	tests := []struct {
		r    Range
		want []string
	}{
		{*NewRange("11", "22"), []string{"11", "22"}},
		{*NewRange("95", "115"), []string{"99"}},
		{*NewRange("998", "1012"), []string{"1010"}},
		{*NewRange("1188511880", "1188511890"), []string{"1188511885"}},
		{*NewRange("222220", "222224"), []string{"222222"}},
		{*NewRange("1698522", "1698528"), []string{}},
		{*NewRange("446443", "446449"), []string{"446446"}},
		{*NewRange("38593856", "38593862"), []string{"38593859"}},
	}

	for _, test := range tests {
		got := test.r._2Sequences()
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("SequencesRepeatedTwice(%v) = %v, want %v", test.r, got, test.want)
		}
	}
}

func TestSequencesRepeatedNTimes(t *testing.T) {
	tests := []struct {
		r    Range
		want []string
	}{
		{*NewRange("11", "22"), []string{"11", "22"}},
		{*NewRange("95", "115"), []string{"99", "111"}},
		{*NewRange("998", "1012"), []string{"999", "1010"}},
		{*NewRange("1188511880", "1188511890"), []string{"1188511885"}},
		{*NewRange("222220", "222224"), []string{"222222"}},
		{*NewRange("1698522", "1698528"), []string{}},
		{*NewRange("446443", "446449"), []string{"446446"}},
		{*NewRange("38593856", "38593862"), []string{"38593859"}},
		{*NewRange("565653", "565659"), []string{"565656"}},
		{*NewRange("824824821", "824824827"), []string{"824824824"}},
		{*NewRange("2121212118", "2121212124"), []string{"2121212121"}},
	}

	for _, test := range tests {
		got := test.r._NSequences()
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("SequencesRepeatedNTimes(%v) = %v, want %v", test.r, got, test.want)
		}
	}
}
