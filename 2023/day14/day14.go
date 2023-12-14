package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("2023/day14/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	//var lines []string
	//for scanner.Scan() {
	//	lines = append(lines, scanner.Text())
	//
	//}

	table := make(map[[2]int]string)
	y := 0
	var maxX int
	var maxY int
	for scanner.Scan() {
		for x, pipe := range scanner.Text() {
			table[[2]int{x, y}] = string(pipe)
		}
		y++
		maxX = len(scanner.Text()) - 1
	}
	maxY = y - 1
	fmt.Println(maxX, maxY)

	utils.Print2DStringsGrid(table)
	for {
		counter := 0
		for pos, s := range table {
			if s == "O" {
				if table[[2]int{pos[0], pos[1] - 1}] == "." {
					table[[2]int{pos[0], pos[1] - 1}] = "O"
					table[pos] = "."
					counter++
				}
			}
		}
		if counter == 0 {
			break
		}
	}
	utils.Print2DStringsGrid(table)

	load := 0
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if table[[2]int{x, y}] == "O" {
				load += (maxY + 1) - y
			}
		}
	}
	fmt.Println(load)

}
