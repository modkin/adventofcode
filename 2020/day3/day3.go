package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var grid []string
	file, err := os.Open("2020/day3/input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	maxX := 0
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
		if length := len(grid[len(grid)-1]); length > maxX {
			maxX = length
		}
	}

	slope := func(xStep int, yStep int) int {
		x, trees := 0, 0
		for y := 0; y < len(grid); y = y + yStep {
			if string(grid[y][x]) == "#" {
				trees++
			}
			x = (x + xStep) % (maxX)
		}
		return trees
	}

	fmt.Println("Task 3.1:", slope(3, 1))
	total := slope(3, 1) * slope(1, 1) * slope(5, 1) * slope(7, 1) * slope(1, 2)
	fmt.Println("Task 3.2:", total)
}
