package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	lines := utils.ReadFileIntoLines("2025/day4/input")

	diagram := make(map[[2]int]string)

	for y, line := range lines {
		for x, s := range line {
			diagram[[2]int{x, y}] = string(s)
		}
	}

	utils.Print2DStringsGrid(diagram)

	task1 := 0
	totalAccessable := 0
	for {
		accessable := 0
		newDiagram := make(map[[2]int]string)
		for ints, s := range diagram {
			counter := 0
			if s == "@" {
				for y := -1; y <= 1; y++ {
					for x := -1; x <= 1; x++ {
						if diagram[[2]int{ints[0] + x, ints[1] + y}] == "@" {
							counter++
						}
					}
				}
				if counter <= 4 {
					accessable++
				} else {
					newDiagram[ints] = "@"
				}
			}
		}
		diagram = newDiagram
		if task1 == 0 {
			task1 = accessable
		}
		totalAccessable += accessable
		if accessable == 0 {
			break
		}
	}
	fmt.Println("Day 4.1:", task1)
	fmt.Println("Day 4.2:", totalAccessable)

}
