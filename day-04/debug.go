package main

import "fmt"

func P(a ...any) {
	if *Debug {
		fmt.Println(a...)
	}
}
