package main

import (
	"fmt"
	"time"
)

// clearScreen clears the terminal using ANSI escape codes
func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

// SuperJoltageExplained demonstrates the greedy stack algorithm step by step
// This is for educational purposes only - it prints each step of the algorithm
func SuperJoltageExplained(batteries []int, k int) int {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		fmt.Printf("Elapsed time: %s\n", elapsed)
	}()

	clearScreen()
	fmt.Printf("\n=== Super Joltage Algorithm Explanation ===\n")
	fmt.Printf("Input: %v\n", batteries)
	fmt.Printf("Goal: Select %d digits to form the largest number\n", k)
	fmt.Printf("Strategy: Use a greedy stack - remove smaller digits when we find larger ones\n\n")

	n := len(batteries)
	toRemove := n - k
	stack := make([]int, 0, k)

	fmt.Printf("We need to remove %d digits (keep %d out of %d)\n\n", toRemove, k, n)
	time.Sleep(500 * time.Millisecond)

	for i := 0; i < n; i++ {
		digit := batteries[i]
		clearScreen()
		fmt.Printf("=== Super Joltage Algorithm Explanation ===\n")
		fmt.Printf("Input: %v\n", batteries)
		fmt.Printf("Goal: Select %d digits to form the largest number\n\n", k)
		fmt.Printf("Step %d: Processing digit %d at position %d\n", i+1, digit, i)
		fmt.Printf("  Current stack: %v\n", stack)
		fmt.Printf("  Digits remaining to remove: %d\n", toRemove)

		// Try to remove smaller digits from the stack
		removed := false
		for len(stack) > 0 && toRemove > 0 && digit > stack[len(stack)-1] {
			removedDigit := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			toRemove--
			removed = true
			fmt.Printf("  -> Removed %d from stack (found larger digit %d)\n", removedDigit, digit)
			fmt.Printf("  -> Stack after removal: %v, remaining to remove: %d\n", stack, toRemove)
			time.Sleep(500 * time.Millisecond)
		}

		if !removed {
			fmt.Printf("  -> Keeping current stack (digit %d is not larger than top)\n", digit)
			time.Sleep(500 * time.Millisecond)
		}

		stack = append(stack, digit)
		fmt.Printf("  -> Added %d to stack\n", digit)
		fmt.Printf("  -> Final stack: %v\n\n", stack)
		time.Sleep(500 * time.Millisecond)
	}

	clearScreen()
	fmt.Printf("=== Super Joltage Algorithm Explanation ===\n")
	fmt.Printf("Input: %v\n", batteries)
	fmt.Printf("Goal: Select %d digits to form the largest number\n\n", k)
	fmt.Printf("After processing all digits:\n")
	fmt.Printf("  Stack: %v (length: %d)\n", stack, len(stack))
	fmt.Printf("  Still need to remove: %d digits\n", toRemove)
	time.Sleep(500 * time.Millisecond)

	if toRemove > 0 {
		fmt.Printf("  -> Removing %d smallest digits from the end\n", toRemove)
		fmt.Printf("  -> Removing: %v\n", stack[len(stack)-toRemove:])
		stack = stack[:len(stack)-toRemove]
		fmt.Printf("  -> Final stack: %v\n", stack)
		time.Sleep(500 * time.Millisecond)
	}

	clearScreen()
	fmt.Printf("=== Super Joltage Algorithm Explanation ===\n")
	fmt.Printf("Input: %v\n", batteries)
	fmt.Printf("Goal: Select %d digits to form the largest number\n\n", k)
	fmt.Printf("Final stack: %v\n\n", stack)
	fmt.Printf("Building the number from stack:\n")
	result := 0
	for i, d := range stack {
		oldResult := result
		result = result*10 + d
		fmt.Printf("  Position %d: digit %d -> result = %d * 10 + %d = %d\n", i, d, oldResult, d, result)
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Printf("\n=== Final Result: %d ===\n\n", result)
	time.Sleep(1 * time.Second)
	return result
}
