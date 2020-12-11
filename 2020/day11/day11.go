package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"strings"
)

func printFloorplan(floorplan [][]string) {
	for i, line := range floorplan {
		if i < 3 {
			fmt.Println(line)
		}
	}
}

func getNeighbors(floorplan [][]string, i int, j int) [][2]int {
	neighbors := make([][2]int, 0)
	for di := -1; di < 2; di++ {
		for dj := -1; dj < 2; dj++ {
			if (i+di) >= 0 && (j+dj) >= 0 && (i+di) < len(floorplan) && (j+dj) < len(floorplan[0]) {
				neighbors = append(neighbors, [2]int{i + di, j + dj})
			}
		}
	}
	return neighbors
}

func update(floorplan [][]string) bool {
	change := false
	floorplanCopy := make([][]string, len(floorplan))
	for i, _ := range floorplanCopy {
		floorplanCopy[i] = utils.CopyStringSlice(floorplan[i])
	}
	for i := 0; i < len(floorplan); i++ {
		for j := 0; j < len(floorplan[0]); j++ {
			occupied := 0
			for _, elem := range getNeighbors(floorplan, i, j) {
				if floorplanCopy[elem[0]][elem[1]] == "#" {
					occupied++
				}
			}
			if floorplanCopy[i][j] == "L" {
				if occupied == 0 {
					floorplan[i][j] = "#"
					change = true
				}
			}
			if floorplanCopy[i][j] == "#" {
				if occupied >= 5 {
					floorplan[i][j] = "L"
					change = true
				}
			}
		}
	}
	return change
}

func main() {
	floorplanMap := make(map[[2]int]string)

	scanner := bufio.NewScanner(utils.OpenFile("2020/day11/input"))
	row, col := 0, 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		col = len(line)
		for i, elem := range line {
			floorplanMap[[2]int{row, i}] = elem
		}
		row++
	}
	floorplan := make([][]string, row)
	for i := range floorplan {
		floorplan[i] = make([]string, col)
		for j := 0; j < col; j++ {
			floorplan[i][j] = floorplanMap[[2]int{i, j}]
		}
	}
	//fmt.Println(getNeighbors(floorplan,len(floorplan)-1,len(floorplan[0])-1))
	fmt.Println(getNeighbors(floorplan, 10, len(floorplan[0])-1))
	fmt.Println(getNeighbors(floorplan, 0, len(floorplan[0])-1))
	fmt.Println(getNeighbors(floorplan, len(floorplan)-1, 0))
	fmt.Println(getNeighbors(floorplan, len(floorplan)-2, len(floorplan[0])-2))
	for update(floorplan) {
	}

	seats := 0
	for i := range floorplan {
		for j := range floorplan[0] {
			if floorplan[i][j] == "#" {
				seats++
			}
		}
	}
	fmt.Println(seats)
	main2()

}
