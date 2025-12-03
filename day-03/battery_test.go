package main

import "testing"

func TestMaxJoltage(t *testing.T) {
	tests := []struct {
		batteries string
		want      int
	}{
		{"987654321111111", 98},
		{"811111111111119", 89},
		{"234234234234278", 78},
		{"818181911112111", 92},
	}

	for _, test := range tests {
		battery := NewBatteryBank(test.batteries)
		got := battery.MaxJoltage()
		if got != test.want {
			t.Errorf("MaxJoltage(%s) = %d, want %d", test.batteries, got, test.want)
		}
	}
}

func TestSuperJoltage(t *testing.T) {
	tests := []struct {
		batteries string
		want      int
	}{
		{"987654321111111", 987654321111},
		{"811111111111119", 811111111119},
		{"234234234234278", 434234234278},
		{"818181911112111", 888911112111},
	}

	for _, test := range tests {
		battery := NewBatteryBank(test.batteries)
		got := battery.SuperJoltage()
		if got != test.want {
			t.Errorf("SuperJoltage(%s) = %d, want %d", test.batteries, got, test.want)
		}
	}
}
