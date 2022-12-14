package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("2022/day14/input")
	if err != nil {
		panic(err)
	}

	grid := make(map[[2]int]string)
	scanner := bufio.NewScanner(file)

	yMax := 0
	grid[[2]int{500, 0}] = "+"
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		stx := utils.ToInt(strings.Split(line[0], ",")[0])
		sty := utils.ToInt(strings.Split(line[0], ",")[1])
		grid[[2]int{stx, sty}] = "#"
		for i := 1; i < len(line); i++ {
			stopX := utils.ToInt(strings.Split(line[i], ",")[0])
			stopY := utils.ToInt(strings.Split(line[i], ",")[1])
			if stopY > yMax {
				yMax = stopY
			}
			incx := utils.Sgn(stopX - stx)
			incy := utils.Sgn(stopY - sty)
			x := stx
			y := sty
			for {
				grid[[2]int{x, y}] = "#"
				if x == stopX && y == stopY {
					break
				}
				x += incx
				y += incy

			}
			stx = stopX
			sty = stopY
		}
	}
	fmt.Println(grid)
	utils.Print2DStringsGrid(grid)

	source := [2]int{500, 0}
	pos := source
	counter := 0
	for {
		if _, down := grid[[2]int{pos[0], pos[1] + 1}]; down {
			if _, downLeft := grid[[2]int{pos[0] - 1, pos[1] + 1}]; downLeft {
				if _, downRight := grid[[2]int{pos[0] + 1, pos[1] + 1}]; downRight {
					grid[pos] = "o"
					pos = source
					//utils.Print2DStringsGrid(grid)
					counter++
				} else {
					pos = [2]int{pos[0] + 1, pos[1] + 1}
				}
			} else {
				pos = [2]int{pos[0] - 1, pos[1] + 1}
			}
		} else {
			pos = [2]int{pos[0], pos[1] + 1}
		}
		if pos[1] > yMax {
			break
		}
	}
	fmt.Println("Day 14.1:", counter)

}
