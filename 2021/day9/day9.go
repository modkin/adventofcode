package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type vList struct {
	p [][2]int
}

func notVisited(visited [][2]int, pos [2]int) bool {
	for _, v := range visited {
		if pos[0] == v[0] && pos[1] == v[1] {
			return false
		}
	}
	return true
}

func getSize(grid [][]int, positions [][2]int, visited *vList) int {
	found := false
	newPositions := make([][2]int, 0)
	size := 0
	for _, p := range positions {
		if grid[p[0]-1][p[1]] != 9 {
			newPos := [2]int{p[0] - 1, p[1]}
			if notVisited(visited.p, newPos) {
				newPositions = append(newPositions, newPos)
				visited.p = append(visited.p, newPos)
				size++
				found = true
			}
		}
		if grid[p[0]+1][p[1]] != 9 {
			newPos := [2]int{p[0] + 1, p[1]}
			if notVisited(visited.p, newPos) {
				newPositions = append(newPositions, newPos)
				visited.p = append(visited.p, newPos)
				size++
				found = true
			}
		}
		if grid[p[0]][p[1]-1] != 9 {
			newPos := [2]int{p[0], p[1] - 1}
			if notVisited(visited.p, newPos) {
				newPositions = append(newPositions, newPos)
				visited.p = append(visited.p, newPos)
				size++
				found = true
			}
		}
		if grid[p[0]][p[1]+1] != 9 {
			newPos := [2]int{p[0], p[1] + 1}
			if notVisited(visited.p, newPos) {
				newPositions = append(newPositions, newPos)
				visited.p = append(visited.p, newPos)
				size++
				found = true
			}
		}
		if found {
			size += getSize(grid, newPositions, visited)
		}
	}
	return size
}

func main() {
	file, err := os.Open("2021/day9/input")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}

	grid := make([][]int, 0)
	topBorder := make([]int, 0)
	first := true
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		newRow := make([]int, 0)
		newRow = append(newRow, 9)
		if first {
			for i := 0; i < len(line)+2; i++ {
				topBorder = append(topBorder, 9)
			}
			grid = append(grid, topBorder)
			first = false
		}
		for _, l := range line {
			newRow = append(newRow, utils.ToInt(l))
		}
		newRow = append(newRow, 9)
		grid = append(grid, newRow)
	}
	grid = append(grid, topBorder)

	lowPoints := make([]int, 0)
	lowPointsCoords := make([][2]int, 0)
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[0])-1; x++ {
			if grid[y][x-1] <= grid[y][x] {
				continue
			}

			if grid[y][x+1] <= grid[y][x] {
				continue
			}

			if grid[y-1][x] <= grid[y][x] {
				continue
			}

			if grid[y+1][x] <= grid[y][x] {
				continue
			}
			lowPoints = append(lowPoints, grid[y][x])
			lowPointsCoords = append(lowPointsCoords, [2]int{y, x})
		}
	}
	sum := 0
	for _, point := range lowPoints {
		sum += point + 1
	}
	fmt.Println("Day 9.1:", sum)

	//start := [2]int{lowPointsCoords[0][0], lowPointsCoords[0][1]}
	sizes := make([]int, 0)
	for _, coord := range lowPointsCoords {
		positions := make([][2]int, 0)
		positions = append(positions, coord)
		visited := make([][2]int, 0)
		visited = append(visited, coord)
		v := vList{visited}
		size := getSize(grid, positions, &v)
		sizes = append(sizes, size+1)
	}

	sort.Ints(sizes)
	length := len(sizes)
	fmt.Println("Day 9.2:", sizes[length-1]*sizes[length-2]*sizes[length-3])
}
