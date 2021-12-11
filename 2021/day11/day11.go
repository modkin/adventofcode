package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printGrid(grid [][]int) {
	fmt.Println("-------------------------------------")
	for _, g := range grid {
		fmt.Println(g)
	}
}

func increaseOne(grid [][]int) {
	for y := 1; y <= 10; y++ {
		for x := 1; x <= 10; x++ {
			grid[y][x]++
		}
	}
}

func reset(grid [][]int) {
	for y := 1; y <= 10; y++ {
		for x := 1; x <= 10; x++ {
			if grid[y][x] > 9 {
				grid[y][x] = 0
			}
		}
	}
	for i := 1; i < 12; i++ {
		grid[0][i] = 0
		grid[i][0] = 0
		grid[11][i] = 0
		grid[i][11] = 0
	}
}

func flash(grid [][]int) int {
	flashes := 0
	for y := 1; y <= 10; y++ {
		for x := 1; x <= 10; x++ {
			if grid[y][x] >= 10 && grid[y][x] < 100 {
				flashes += 1
				for y1 := y - 1; y1 <= y+1; y1++ {
					for x1 := x - 1; x1 <= x+1; x1++ {
						grid[y1][x1]++
					}
				}
				grid[y][x] = 100
			}
		}
	}
	if flashes > 0 {
		flashes += flash(grid)
	}
	return flashes
}

func main() {
	file, err := os.Open("2021/day11/input")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}

	grid := make([][]int, 0)
	first := true
	topBorder := make([]int, 0)
	for i := 0; i < 12; i++ {
		topBorder = append(topBorder, 0)
	}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		newRow := make([]int, 0)
		newRow = append(newRow, 0)
		if first {
			grid = append(grid, topBorder)
			first = false
		}
		for _, l := range line {
			newRow = append(newRow, utils.ToInt(l))
		}
		newRow = append(newRow, 0)
		grid = append(grid, newRow)
	}
	grid = append(grid, utils.CopyIntSlice(topBorder))
	printGrid(grid)
	flashes := 0
	for timestep := 0; timestep < 100; timestep++ {
		increaseOne(grid)
		flashes += flash(grid)
		reset(grid)
		//printGrid(grid)
	}
	fmt.Println(flashes)
}
