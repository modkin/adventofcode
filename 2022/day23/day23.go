package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func elveAround(grid map[[2]int]string, pos [2]int) bool {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if val, ok := grid[[2]int{pos[0] + x, pos[1] + y}]; ok {
				if val == "#" {
					return true
				}
			} else {
				fmt.Println("ERROR - to small")
			}
		}
	}
	return false
}

func add(first [2]int, second [2]int) [2]int {
	return [2]int{first[0] + second[0], first[1] + second[1]}
}

func increaseGrid(grid map[[2]int]string) int {
	xMin, yMin, xMax, yMax := math.MaxInt, math.MaxInt, 0, 0
	for i, val := range grid {
		if val == "#" {
			if i[0] < xMin {
				xMin = i[0]
			}
			if i[0] > xMax {
				xMax = i[0]
			}
			if i[1] < yMin {
				yMin = i[1]
			}
			if i[1] > yMax {
				yMax = i[1]
			}
		}
	}
	for x := xMin - 1; x <= xMax+1; x++ {
		grid[[2]int{x, yMin - 1}] = "."
		grid[[2]int{x, yMax + 1}] = "."
	}
	for y := yMin - 1; y <= yMax+1; y++ {
		grid[[2]int{xMin - 1, y}] = "."
		grid[[2]int{xMax + 1, y}] = "."
	}
	return (yMax + 1 - yMin) * (xMax + 1 - xMin)
}
func checkfree(grid map[[2]int]string, pos [2]int, dir [2]int) bool {
	target := add(pos, dir)
	if dir[0] == 0 {
		for i := -1; i <= 1; i++ {
			tmp := add([2]int{i, 0}, target)
			if grid[tmp] == "#" {
				return false
			}
		}
	} else {
		for i := -1; i <= 1; i++ {
			if grid[add([2]int{0, i}, target)] == "#" {
				return false
			}
		}
	}
	return true
}

func main() {
	file, err := os.Open("2022/day23/input")

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	grid := make(map[[2]int]string)
	possDirs := [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	numElves := 0

	yPos := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for xPos, s := range line {
			if s != " " {
				grid[[2]int{xPos, yPos}] = s
				if s == "#" {
					numElves++
				}
			}
		}
		yPos++
	}
	grid[[2]int{0, 0}] = "*"

	utils.Print2DStringsGrid(grid)
	increaseGrid(grid)
	utils.Print2DStringsGrid(grid)
	for rounds := 0; rounds < 10; rounds++ {
		fmt.Println("Round", rounds)
		increaseGrid(grid)
		proposed := make(map[[2]int]bool)
		moves := make(map[[2]int][2]int)
		collison := make(map[[2]int]bool)

		for elvePos, value := range grid {
			if value == "#" {
				if elveAround(grid, elvePos) {
					for _, dir := range possDirs {
						if checkfree(grid, elvePos, dir) {
							target := add(elvePos, dir)
							moves[elvePos] = target
							if _, already := proposed[target]; already {
								//an elve already wants to go there -> no one moves
								collison[target] = true
							} else {
								proposed[target] = true
							}
							break
						}
					}
				}
			}
		}

		for elvePos := range grid {
			if move, ok := moves[elvePos]; ok {
				if _, col := collison[move]; !col {
					grid[move] = "#"
					grid[elvePos] = "."
				}
			}
		}
		possDirs = append(possDirs[1:4], possDirs[0])
		utils.Print2DStringsGrid(grid)
	}
	area := increaseGrid(grid)
	fmt.Println(area - numElves)

}
