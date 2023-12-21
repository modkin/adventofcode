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
	file, err := os.Open("2023/day21/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	garden := make(map[[2]int]string)
	y := 0
	var maxX int
	var maxY int
	var start [2]int
	for scanner.Scan() {
		for x, cost := range scanner.Text() {
			garden[[2]int{x, y}] = string(cost)
			if string(cost) == "S" {
				start = [2]int{x, y}
			}
		}
		y++
		maxX = len(scanner.Text()) - 1

	}
	maxY = y - 1
	fmt.Println(maxX, maxY)
	utils.Print2DStringsGrid(garden)

	allPos := make(map[[2]int]int)
	allPos[start] = 0
	var dirs [][2]int
	dirs = append(dirs, [2]int{1, 0})
	dirs = append(dirs, [2]int{-1, 0})
	dirs = append(dirs, [2]int{0, 1})
	dirs = append(dirs, [2]int{0, -1})
	//reachable := make(map[[2]int]bool)
	for i := 0; i < 64; i++ {

		newAllPos := make(map[[2]int]int)
		for po, oldDist := range allPos {
			for _, dir := range dirs {

				newPos := sum(po, dir)
				//if newPos[0] < minX-1 || newPos[0] > maxX+1 || newPos[1] < minY-1 || newPos[1] > maxY+1 {
				//	continue
				//}
				if garden[newPos] != "#" {
					newAllPos[newPos] = oldDist + 1

				}
			}
		}
		allPos = newAllPos
	}
	maxG := 0
	for _, i := range allPos {
		if i > maxG {
			maxG = i
		}
	}
	fmt.Println(len(allPos))

}
