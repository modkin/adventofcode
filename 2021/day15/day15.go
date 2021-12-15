package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type node struct {
	dist    int
	visited bool
}

func printGrid(grid [][]int) {
	fmt.Println("-------------------------------------")
	for _, g := range grid {
		fmt.Println(g)
	}
}

func findSmallestUnvisited(dijstra map[[2]int]*node) ([2]int, int) {
	var smallestCoord [2]int
	smallestDist := math.MaxInt
	for key, coords := range dijstra {
		if coords.visited == false {
			if coords.dist < smallestDist {
				smallestDist = coords.dist
				smallestCoord = key
			}
		}
	}
	return smallestCoord, smallestDist
}

func nextStep(dijstra map[[2]int]*node, distances [][]int) {
	start, currentDist := findSmallestUnvisited(dijstra)
	dijstra[start].visited = true
	directions := [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for _, d := range directions {
		nbr := [2]int{start[0] + d[0], start[1] + d[1]}
		if val, ok := dijstra[nbr]; ok {
			if val.visited == false {
				if newDist := currentDist + distances[nbr[0]][nbr[1]]; newDist < dijstra[nbr].dist {
					dijstra[nbr].dist = newDist
				}
			}
		} else {
			newDist := currentDist + distances[nbr[0]][nbr[1]]
			newNode := node{newDist, false}
			dijstra[nbr] = &newNode
		}
	}
}

func main() {
	file, err := os.Open("2021/day15/input")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}

	grid := make([][]int, 0)
	dijstra := make(map[[2]int]*node)

	var border []int
	first := true
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		if first {
			border = make([]int, len(line)+2)
			for i := range border {
				border[i] = math.MaxInt / 2
			}
			grid = append(grid, border)
			first = false
		}
		newline := make([]int, 0)
		newline = append(newline, math.MaxInt/2)
		for _, l := range line {
			newline = append(newline, utils.ToInt(l))
		}
		newline = append(newline, math.MaxInt/2)
		grid = append(grid, newline)
	}
	grid = append(grid, border)

	target := [2]int{len(grid[0]) - 2, len(grid) - 2}
	start := [2]int{1, 1}
	dijstra[start] = &node{0, false}
	for {
		if val, ok := dijstra[target]; ok {
			if val.visited {
				fmt.Println(val.dist)
				break
			}
		}
		nextStep(dijstra, grid)
	}
	grid = grid[1 : len(grid)-1]
	for i, ints := range grid {
		grid[i] = ints[1 : len(ints)-1]
	}
	xlen := len(grid[0])
	ylen := len(grid)
	for block := 1; block < 9; block++ {
		for y := 0; y < ylen; y++ {
			for x := 0; x < xlen; x++ {
				nextNumber := grid[y][x] + block
				if nextNumber > 9 {
					nextNumber = nextNumber % 9
				}
				grid[y] = append(grid[y], nextNumber)
			}
		}
	}
	for block := 1; block < 5; block++ {
		for y := 0; y < ylen; y++ {
			grid = append(grid, utils.CopyIntSlice(grid[y][xlen*block:xlen*block+xlen*5]))
		}
	}
	for y := 0; y < ylen; y++ {
		grid[y] = grid[y][:ylen*5]
	}
	border = make([]int, len(grid[0])+2)
	for i := 0; i < len(border); i++ {
		border[i] = math.MaxInt / 2
	}
	for i := range grid {
		grid[i] = append(grid[i], math.MaxInt/2)
		grid[i] = append([]int{math.MaxInt / 2}, grid[i]...)
	}
	grid = append([][]int{border}, grid...)
	grid = append(grid, border)
	//printGrid(grid)
	target = [2]int{len(grid[0]) - 2, len(grid) - 2}
	start = [2]int{1, 1}
	dijstra = make(map[[2]int]*node)
	dijstra[start] = &node{0, false}
	for {
		if val, ok := dijstra[target]; ok {
			if val.visited {
				fmt.Println(val.dist)
				break
			}
		}
		nextStep(dijstra, grid)
	}
}
