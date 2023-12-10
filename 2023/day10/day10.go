package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
)

func sum(left [2]int, right [2]int) [2]int {
	return [2]int{left[0] + right[0], left[1] + right[1]}
}

func main() {
	file, err := os.Open("2023/day10/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	pipeMap := make(map[[2]int]string)

	var start [2]int
	nextDirs := make(map[string][][2]int)
	nextDirs["-"] = [][2]int{{1, 0}, {-1, 0}}
	nextDirs["|"] = [][2]int{{0, 1}, {0, -1}}
	nextDirs["L"] = [][2]int{{1, 0}, {0, -1}}
	nextDirs["F"] = [][2]int{{1, 0}, {0, 1}}
	nextDirs["7"] = [][2]int{{-1, 0}, {0, 1}}
	nextDirs["J"] = [][2]int{{-1, 0}, {0, -1}}
	y := 0
	for scanner.Scan() {
		for x, pipe := range scanner.Text() {
			pipeMap[[2]int{x, y}] = string(pipe)
			if pipe == 'S' {
				start = [2]int{x, y}
			}
		}
		y++
	}
	utils.Print2DStringsGrid(pipeMap)
	fmt.Println(start)

	//for ints, s := range pipeMap {
	//	for _, i := range nextDirs[s] {
	//		nbr := pipeMap[sum(ints, i)]
	//		if nbr == "." {
	//			panic(s)
	//		}
	//	}
	//}
	var startType string
	findStartType := func(startPos [2]int, startType string) bool {
		connections := 0
		for _, i := range nextDirs[startType] {
			nbrPos := sum(startPos, i)
			for _, j := range nextDirs[pipeMap[nbrPos]] {
				if startPos == sum(nbrPos, j) {
					connections++
				}
			}
		}
		return connections == 2
	}
	for _, s := range []string{"-", "|", "J", "F", "L", "7"} {
		if findStartType(start, s) {
			startType = s
			fmt.Println("Start", s)
		}
	}
	previousPos := start
	//choose first step random
	currentPos := sum(start, nextDirs[startType][0])
	numberOfSteps := 1
	for currentPos != start {
		fmt.Println(currentPos)
		for _, i := range nextDirs[pipeMap[currentPos]] {
			if nextPos := sum(currentPos, i); nextPos != previousPos {
				previousPos = currentPos
				currentPos = nextPos
				break
			}
		}
		numberOfSteps++

	}
	fmt.Println(numberOfSteps / 2)
}
