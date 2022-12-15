package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("2022/day15/input")
	if err != nil {
		panic(err)
	}

	grid := make(map[[2]int]string)
	scanner := bufio.NewScanner(file)

	grid[[2]int{500, 0}] = "+"
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		xSen := utils.ToInt(strings.Split(strings.Trim(line[2], ","), "=")[1])
		ySen := utils.ToInt(strings.Split(strings.Trim(line[3], ":"), "=")[1])
		xBe := utils.ToInt(strings.Split(strings.Trim(line[8], ","), "=")[1])
		yBe := utils.ToInt(strings.Split(line[9], "=")[1])
		xdist := utils.IntAbs(xBe - xSen)
		ydist := utils.IntAbs(yBe - ySen)
		grid[[2]int{xSen, ySen}] = "S"
		walk := make(map[[2]int]bool)
		walk[[2]int{xSen, ySen}] = true

		for i := 0; i <= xdist+ydist; i++ {
			newWalk := make(map[[2]int]bool)
			for pos := range walk {
				grid[pos] = "#"
				newWalk[[2]int{pos[0] + 1, pos[1]}] = true
				newWalk[[2]int{pos[0], pos[1] + 1}] = true
				newWalk[[2]int{pos[0] - 1, pos[1]}] = true
				newWalk[[2]int{pos[0], pos[1] - 1}] = true
			}
			walk = newWalk

		}
		grid[[2]int{xSen, ySen}] = "S"
		grid[[2]int{xBe, yBe}] = "B"
		//utils.Print2DStringsGrid(grid)
		fmt.Println("new Line")
	}
	xMin, yMin, xMax, yMax := math.MaxInt, math.MaxInt, 0, 0
	for i := range grid {
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

	counter := 0
	yCoord := 2000000
	for i := xMin; i < xMax; i++ {
		//fmt.Print(grid[[2]int{i, yCoord}])
		if grid[[2]int{i, yCoord}] == "#" {
			counter++
		}
	}
	fmt.Println()
	fmt.Println(counter)

}
