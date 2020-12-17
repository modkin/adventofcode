package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"strings"
)

func getNeighbors(coord [4]int, fourD bool) [][4]int {
	max := 27
	if fourD {
		max = 81
	}
	ret := make([][4]int, max)
	for i := 0; i < max; i++ {
		dw := 1 - i/27
		dz := 1 - i%27/9
		dy := 1 - i%27%9/3
		dx := 1 - i%27%3
		//fmt.Println(dx, dy, dz)
		if !fourD {
			dw = 0
		}
		nbr := [4]int{coord[0] + dx, coord[1] + dy, coord[2] + dz, coord[3] + dw}
		ret[i] = nbr
	}
	return ret
}

func main() {
	scanner := bufio.NewScanner(utils.OpenFile("2020/day17/input"))

	grid := make(map[[4]int]bool)
	grid2 := make(map[[4]int]bool)
	backupGrid := make(map[[4]int]bool)

	y := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for x, elem := range line {
			if elem == "#" {
				grid[[4]int{x, y, 0, 0}] = true
			} else {
				grid[[4]int{x, y, 0, 0}] = false
			}
		}
		y++
	}
	iterate := func(fourD bool) {
		for coord, val := range grid {
			if val == false {
				delete(grid, coord)
				delete(grid2, coord)
			}
		}
		for coord := range grid2 {
			for _, nbr := range getNeighbors(coord, fourD) {
				if check := grid[nbr]; !check {
					grid[nbr] = false
				}
			}
		}
		for coord, value := range grid {
			countActive := 0
			for _, nbr := range getNeighbors(coord, fourD) {
				if grid[nbr] {
					countActive++
				}
			}
			if value {
				if countActive == 3 || countActive == 4 {
					grid2[coord] = true
				} else {
					grid2[coord] = false
				}
			} else {
				if countActive == 3 {
					grid2[coord] = true
				} else {
					grid2[coord] = false
				}
			}
		}
	}
	countGrid := func() int {
		count := 0
		for _, value := range grid {
			if value {
				count++
			}
		}
		return count
	}

	//print(grid,2)
	for key, value := range grid {
		backupGrid[key] = value
		grid2[key] = value
	}
	for i := 0; i < 6; i++ {
		//fmt.Println(grid)
		iterate(true)
		for key, value := range grid2 {
			grid[key] = value
		}
	}
	fmt.Println("Task 17.1:", countGrid())
	grid = make(map[[4]int]bool)
	grid2 = make(map[[4]int]bool)
	for key, value := range backupGrid {
		grid[key] = value
		grid2[key] = value
	}
	for i := 0; i < 6; i++ {
		//fmt.Println(grid)
		iterate(false)
		for key, value := range grid2 {
			grid[key] = value
		}
	}
	fmt.Println("Task 17.2:", countGrid())
}
