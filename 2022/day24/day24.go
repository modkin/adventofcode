package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input
var input string

func add(first [2]int, second [2]int) [2]int {
	x := wrapPos(first[0]+second[0], maxX)
	y := wrapPos(first[1]+second[1], maxY)
	return [2]int{x, y}
}

func plainAdd(first [2]int, second [2]int) [2]int {
	return [2]int{first[0] + second[0], first[1] + second[1]}
}

func wrapPos(pos, max int) int {
	pos -= 1
	pos += max
	pos = pos % max
	pos += 1
	return pos
}

var dirMap = map[string][2]int{"^": {0, -1}, "v": {0, 1}, ">": {1, 0}, "<": {-1, 0}}
var allDir = [][2]int{{0, 0}, {1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var maxX, maxY int

func moveWinds(valley map[[2]int]string) {
	newValley := make(map[[2]int]string)
	for pos, s := range valley {
		if s != "#" {
			for _, wind := range s {
				dir := dirMap[string(wind)]
				newPos := add(pos, dir)
				newValley[newPos] += string(wind)
			}
			delete(valley, pos)
		}

	}
	for pos, s := range newValley {
		valley[pos] = s
	}
}

func printValley(grid map[[2]int]string) {
	fmt.Println("----------------------------------------------")
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
	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			if val, ok := grid[[2]int{x, y}]; ok {
				if len(val) > 1 {
					fmt.Print(len(val))
				} else {
					fmt.Print(val)
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	valley := make(map[[2]int]string)

	y := 0
	for _, line := range strings.Split(input, "\n") {
		for x, v := range line {
			if v != '.' {
				valley[[2]int{x, y}] = string(v)
			}
		}
		y++
		maxX = len(line) - 2
	}
	maxY = y - 2

	start := [2]int{1, 0}
	exit := [2]int{maxX, maxY + 1}

	getMinSteps := func(start, target [2]int) int {
		minute := 0
		allPos := make(map[[2]int]bool)
		allPos[start] = true
		for {
			newPositions := make(map[[2]int]bool)
			for pos := range allPos {
				for _, dir := range allDir {
					newPos := plainAdd(pos, dir)
					if newPos == target {
						//fmt.Println("Day 24.1:", minute)
						return minute
					}
					if newPos[1] >= 0 {
						if _, ok := valley[newPos]; !ok {
							newPositions[newPos] = true
						}
					}
				}
			}
			allPos = newPositions
			minute++
			moveWinds(valley)
		}
	}
	first := getMinSteps(start, exit)
	fmt.Println("Day 24.1:", first)
	second := getMinSteps(exit, start)
	third := getMinSteps(start, exit)
	fmt.Println("Day 24.2:", first+second+third)

	//printValley(valley)
}
