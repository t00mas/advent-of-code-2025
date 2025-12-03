package main

import (
	"fmt"
	"strconv"
)

// BatteryBank represents a bank of batteries
type BatteryBank struct {
	Size          int
	Batteries     []int
	batteries_str string
}

// NewBatteryBank creates a new battery bank
func NewBatteryBank(batteries string) *BatteryBank {
	size := len(batteries)
	batteries_ := make([]int, size)
	for i, b := range batteries {
		batteries_[i], _ = strconv.Atoi(string(b))
	}
	return &BatteryBank{Size: size, Batteries: batteries_, batteries_str: batteries}
}

// MaxJoltage returns the maximum joltage of the battery bank
// It's the biggest number that can be formed with any 2 digits of the bank
func (b *BatteryBank) MaxJoltage() int {
	max := 0
	for i := 0; i < b.Size; i++ {
		for j := i + 1; j < b.Size; j++ {
			joltage_str := fmt.Sprintf("%d%d", b.Batteries[i], b.Batteries[j])
			joltage, _ := strconv.Atoi(joltage_str)
			if joltage > max {
				max = joltage
			}
		}
	}
	return max
}

// SuperJoltage returns the super joltage of the battery bank
// It's the biggest number that can be formed with any 12 digits of the bank
func (b *BatteryBank) SuperJoltage() int {
	toRemove := b.Size - 12
	stack := make([]int, 0, 12)

	for i := 0; i < b.Size; i++ {
		digit := b.Batteries[i]
		for len(stack) > 0 && toRemove > 0 && digit > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			toRemove--
		}
		stack = append(stack, digit)
	}

	stack = stack[:len(stack)-toRemove] // remove from the end
	result := 0
	for _, d := range stack {
		result = result*10 + d // int math no allocs
	}

	return result
}
