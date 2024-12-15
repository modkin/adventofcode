package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func add2Int(one, two [2]int) [2]int {
	return [2]int{one[0] + two[0], one[1] + two[1]}
}

func move(wh map[[2]int]string, pos [2]int, dirStr string) [2]int {
	dirMap := map[string][2]int{"<": {-1, 0}, ">": {1, 0}, "v": {0, 1}, "^": {0, -1}}

	dir := dirMap[dirStr]
	var boxes [][2]int
	newPos := add2Int(pos, dir)
	for {
		if wh[newPos] == "O" {
			boxes = append(boxes, newPos)
		} else if wh[newPos] == "." {
			break
		} else if wh[newPos] == "#" {
			boxes = [][2]int{}
			break
		}
		newPos = add2Int(newPos, dir)
	}
	if len(boxes) != 0 {
		for i := len(boxes) - 1; i >= 0; i-- {
			wh[add2Int(boxes[i], dir)] = "O"
		}
		wh[add2Int(pos, dir)] = "."
	}
	newRoboPos := add2Int(pos, dir)
	if wh[newRoboPos] == "." {
		return newRoboPos
	} else {
		return pos
	}
}

func main() {
	lines := utils.ReadFileIntoLines("2024/day15/input")

	wh := make(map[[2]int]string)
	whWorking := true

	xMax := 0
	yMax := 0

	var moves []string
	var roboPos [2]int
	for _, line := range lines {
		if whWorking {
			if line == "" {
				whWorking = false
			} else {
				xMax = len(line)
				split := strings.Split(line, "")
				for i2, s := range split {
					if s == "@" {
						roboPos = [2]int{i2, yMax}
						wh[[2]int{i2, yMax}] = "."
					} else {
						wh[[2]int{i2, yMax}] = s
					}
				}
				yMax++
			}
		} else {
			split := strings.Split(line, "")
			for _, i2 := range split {
				moves = append(moves, i2)
			}
		}

	}
	fmt.Println(xMax, yMax)

	fmt.Println(roboPos)

	for i, s := range moves {
		if wh[roboPos] != "." {
			fmt.Println(i)
		}
		//wh[roboPos] = s
		//utils.Print2DStringsGrid(wh)
		//wh[roboPos] = "."
		roboPos = move(wh, roboPos, s)
		wh[roboPos] = "@"
		//utils.Print2DStringsGrid(wh)
		wh[roboPos] = "."
	}

	sum := 0
	for ints, s := range wh {

		if s == "O" {
			sum += 100*ints[1] + ints[0]
		}
	}
	utils.Print2DStringsGrid(wh)
	fmt.Println("Day 15.1:", sum)
}
