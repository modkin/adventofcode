package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	grid := make(map[[2]int]string)
	file, err := os.Open("2020/day3/input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	row := 0
	maxX := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for i, elem := range line {
			grid[[2]int{i, row}] = elem
			if i > maxX {
				maxX = i
			}
		}
		row++
	}

	slope := func(xStep int, yStep int) int {
		x, trees := 0, 0
		for y := 0; y < row; y = y + yStep {
			if grid[[2]int{x, y}] == "#" {
				trees++
			}
			x = (x + xStep) % (maxX + 1)
		}
		return trees
	}

	fmt.Println("Task 3.1:", slope(3, 1))
	total := slope(3, 1) * slope(1, 1) * slope(5, 1) * slope(7, 1) * slope(1, 2)
	fmt.Println("Task 3.2:", total)
}
