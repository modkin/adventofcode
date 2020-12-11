package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"strings"
)

func printFloorplan(floorplan [][]string) {
	for _, line := range floorplan {
		fmt.Println(line)

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

func update2(floorplan [][]string) bool {
	change := false
	floorplanCopy := make([][]string, len(floorplan))
	for i, _ := range floorplanCopy {
		floorplanCopy[i] = utils.CopyStringSlice(floorplan[i])
	}
	for i := 0; i < len(floorplan); i++ {
		for j := 0; j < len(floorplan[0]); j++ {
			occupied := 0

			for dir := i - 1; dir >= 0; dir-- {
				if floorplanCopy[dir][j] == "L" {
					break
				}
				if floorplanCopy[dir][j] == "#" {
					occupied++
					break
				}
			}
			for dir := i + 1; dir < len(floorplan); dir++ {
				if floorplanCopy[dir][j] == "L" {
					break
				}
				if floorplanCopy[dir][j] == "#" {
					occupied++
					break
				}
			}

			for dir := j - 1; dir >= 0; dir-- {
				if floorplanCopy[i][dir] == "L" {
					break
				}
				if floorplanCopy[i][dir] == "#" {
					occupied++
					break
				}
			}
			for dir := j + 1; dir < len(floorplan[0]); dir++ {
				if floorplanCopy[i][dir] == "L" {
					break
				}
				if floorplanCopy[i][dir] == "#" {
					occupied++
					break
				}
			}

			for dirRow, dirCol := i-1, j-1; dirRow >= 0 && dirCol >= 0; dirRow, dirCol = dirRow-1, dirCol-1 {
				if floorplanCopy[dirRow][dirCol] == "L" {
					break
				}
				if floorplanCopy[dirRow][dirCol] == "#" {
					occupied++
					break
				}
			}
			for dirRow, dirCol := i+1, j+1; dirRow < len(floorplan) && dirCol < len(floorplan[0]); dirRow, dirCol = dirRow+1, dirCol+1 {
				if floorplanCopy[dirRow][dirCol] == "L" {
					break
				}
				if floorplanCopy[dirRow][dirCol] == "#" {
					occupied++
					break
				}
			}

			for dirRow, dirCol := i-1, j+1; dirRow >= 0 && dirCol < len(floorplan[0]); dirRow, dirCol = dirRow-1, dirCol+1 {
				if floorplanCopy[dirRow][dirCol] == "L" {
					break
				}
				if floorplanCopy[dirRow][dirCol] == "#" {
					occupied++
					break
				}
			}
			for dirRow, dirCol := i+1, j-1; dirRow < len(floorplan) && dirCol >= 0; dirRow, dirCol = dirRow+1, dirCol-1 {
				if floorplanCopy[dirRow][dirCol] == "L" {
					break
				}
				if floorplanCopy[dirRow][dirCol] == "#" {
					occupied++
					break
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

	//for update2(floorplan) {
	//}
	update2(floorplan)
	update2(floorplan)
	printFloorplan(floorplan)
	fmt.Println()
	for update2(floorplan) {
	}
	printFloorplan(floorplan)
	seats := 0
	for i := range floorplan {
		for j := range floorplan[0] {
			if floorplan[i][j] == "#" {
				seats++
			}
		}
	}
	fmt.Println(seats)
	//main2()

}
