package main

import (
	_ "embed"
	"flag"
	"fmt"
)

var Debug = flag.Bool("debug", false, "Debug")
var Example = flag.Bool("example", false, "Use example input")

func PartOne(points []Point) int {
	maxArea := 0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			dx := abs(points[j][0] - points[i][0])
			dy := abs(points[j][1] - points[i][1])
			area := (dx + 1) * (dy + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func PartTwo(points []Point) int {
	poly := NewPolygon(points)

	maxArea := 0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]
			if p1[0] == p2[0] || p1[1] == p2[1] {
				continue
			}

			x1, x2 := min(p1[0], p2[0]), max(p1[0], p2[0])
			y1, y2 := min(p1[1], p2[1]), max(p1[1], p2[1])

			area := (x2 - x1 + 1) * (y2 - y1 + 1)
			if area <= maxArea {
				continue
			}

			if poly.ContainsRect(x1, y1, x2, y2) {
				maxArea = area
			}
		}
	}

	return maxArea
}

func main() {
	flag.Parse()

	parsedInput := ParseInput()

	fmt.Println("Part 1: ")
	fmt.Println(PartOne(parsedInput))

	fmt.Println("Part 2: ")
	fmt.Println(PartTwo(parsedInput))
}
