package main

import (
	"bufio"
	"fmt"
	"os"
)

type hikeStruct struct {
	pos     [2]int
	lastPos [2]int
	length  int
	id      int
}

var dirs = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func sum(left [2]int, right [2]int) [2]int {
	return [2]int{left[0] + right[0], left[1] + right[1]}
}

func main() {
	file, err := os.Open("2023/day23/testinput")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	forestMap := make(map[[2]int]string)
	y := 0
	var maxX int
	var maxY int
	var start [2]int
	for scanner.Scan() {
		for x, cost := range scanner.Text() {
			forestMap[[2]int{x, y}] = string(cost)
			if y == 0 && cost == '.' {
				start = [2]int{x, y}
			}
		}
		y++
		maxX = len(scanner.Text()) - 1

	}
	maxY = y - 1
	fmt.Println("maxX:", maxX, "maxY:", maxY)
	fmt.Println(start)

	allHikeCounter := 0
	newHike := hikeStruct{
		pos:    start,
		length: 0,
		id:     allHikeCounter,
	}
	var allHikes []hikeStruct
	allHikes = append(allHikes, newHike)

	allHikeCounter++
	hikeLenght := make(map[int]int)
	for len(allHikes) != 0 {
		var newAllHikes []hikeStruct
		for _, hike := range allHikes {
			dirCounter := 0
			for _, dir := range dirs {
				newPos := sum(hike.pos, dir)
				if newPos[1] < 0 {
					continue
				}
				if newPos == hike.lastPos {
					continue
				}
				if forestMap[newPos] == "#" {
					continue
				}
				if forestMap[newPos] == ">" && dir != [2]int{1, 0} {
					continue
				}
				if forestMap[newPos] == "<" && dir != [2]int{-1, 0} {
					continue
				}
				if forestMap[newPos] == "v" && dir != [2]int{0, 1} {
					continue
				}
				if forestMap[newPos] == "^" && dir != [2]int{0, -1} {
					continue
				}
				if newPos[1] == maxY {
					hikeLenght[hike.id] = hike.length + 1
				} else {
					newHikeId := hike.id
					if dirCounter != 0 {
						newHikeId = allHikeCounter
						allHikeCounter++
					}
					newHike = hikeStruct{
						lastPos: hike.pos,
						pos:     newPos,
						length:  hike.length + 1,
						id:      newHikeId,
					}
					newAllHikes = append(newAllHikes, newHike)
					dirCounter++
				}
			}
		}
		allHikes = newAllHikes
	}
	fmt.Println(hikeLenght)
	maxLen := 0
	for _, i2 := range hikeLenght {
		if i2 > maxLen {
			maxLen = i2
		}
	}
	fmt.Println("Day 23.1:", maxLen)
}
