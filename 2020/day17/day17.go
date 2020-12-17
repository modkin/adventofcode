package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"strings"
)

func printGrid(grid map[[3]int]bool, max int) {
	for z := -max; z < max; z++ {
		fmt.Println("z=", z)
		for y := -max; y < max+2; y++ {
			for x := -max; x < max+2; x++ {
				if grid[[3]int{x, y, z}] {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}

	}
}

func main() {
	scanner := bufio.NewScanner(utils.OpenFile("2020/day17/input"))

	grid := make(map[[3]int]bool)
	grid2 := make(map[[3]int]bool)

	y := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for x, elem := range line {
			if elem == "#" {
				grid[[3]int{x, y, 0}] = true
			} else {
				grid[[3]int{x, y, 0}] = false
			}
		}
		y++
	}
	fmt.Println(grid)
	iterate := func() {
		for coord := range grid {
			for i := 0; i < 27; i++ {
				dz := 1 - int(i/9)
				dy := 1 - (i%9)/3
				dx := 1 - i%3
				//fmt.Println(dx, dy, dz)
				nbr := [3]int{coord[0] + dx, coord[1] + dy, coord[2] + dz}
				if check := grid[nbr]; !check {
					grid[nbr] = false
				}
			}
		}
		for coord, value := range grid {
			countActive := 0
			for i := 0; i < 27; i++ {
				dz := 1 - int(i/9)
				dy := 1 - (i%9)/3
				dx := 1 - i%3
				//fmt.Println(dx, dy, dz)
				nbr := [3]int{coord[0] + dx, coord[1] + dy, coord[2] + dz}
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

	//print(grid,2)
	for i := 0; i < 6; i++ {
		//fmt.Println(grid)
		iterate()
		for key, value := range grid2 {
			grid[key] = value
		}

	}
	//print(grid,2)
	count := 0
	for _, value := range grid {
		if value {
			count++
		}
	}
	fmt.Println(count)
}
