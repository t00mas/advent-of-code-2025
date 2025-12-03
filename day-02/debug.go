package main

import "fmt"

// P prints a message if Debug is true
func P(a ...any) {
	if *Debug {
		fmt.Println(a...)
	}
}
