package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func hmapprint(grid map[[2]int]Point) {
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
				fmt.Print(val.dist % 9)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func findMin(shortestPath map[[2]int]Point) [2]int {
	var minPos [2]int
	currentMin := math.MaxInt
	for i, i2 := range shortestPath {
		if i2.dist < currentMin {
			currentMin = i2.dist
			minPos = i
		}
	}
	return minPos
}

type Point struct {
	h    rune
	dist int
}

func main() {

	file, err := os.Open("2022/day12/input")
	if err != nil {
		panic(err)
	}

	//grid := make([][]int, 0)
	scanner := bufio.NewScanner(file)

	hmap := make(map[[2]int]Point, 0)

	curPoses := make(map[[2]int]Point)
	//curPoses[[2]int{0, 0}] = Point{'a', 0}
	var endPos [2]int
	var startPos [2]int

	y := 0
	for scanner.Scan() {
		tmp := []rune(scanner.Text())
		for x := 0; x < len(tmp); x++ {
			if tmp[x] == 'E' {
				endPos = [2]int{x, y}
				hmap[[2]int{x, y}] = Point{'z', math.MaxInt}
			} else if tmp[x] == 'S' {
				startPos = [2]int{x, y}
				hmap[[2]int{x, y}] = Point{'a', math.MaxInt}
			} else {
				hmap[[2]int{x, y}] = Point{tmp[x], math.MaxInt}
			}
		}
		y++
	}

	//utils.Print2DStringsGrid(hmap)

	//currentHeight := "A"
	//pos := [2]int{0, 0}
	hmap[startPos] = Point{hmap[startPos].h, 0}
	curPoses[startPos] = Point{'a', 0}

	for {
		//hmapprint(hmap)
		minPos := findMin(curPoses)
		for x := -1; x <= 1; x += 2 {
			newPos := [2]int{minPos[0] + x, minPos[1]}
			if _, ok := hmap[newPos]; ok {
				targetHeight := hmap[newPos].h
				if curPoses[minPos].dist+1 < hmap[newPos].dist && targetHeight-curPoses[minPos].h < 2 {
					curPoses[newPos] = Point{targetHeight, curPoses[minPos].dist + 1}
					hmap[newPos] = Point{hmap[newPos].h, curPoses[minPos].dist + 1}
				}
			}
		}
		for y2 := -1; y2 <= 1; y2 += 2 {
			newPos := [2]int{minPos[0], minPos[1] + y2}
			if _, ok := hmap[newPos]; ok {
				targetHeight := hmap[newPos].h

				if curPoses[minPos].dist+1 < hmap[newPos].dist && targetHeight-curPoses[minPos].h < 2 {
					curPoses[newPos] = Point{targetHeight, curPoses[minPos].dist + 1}
					hmap[newPos] = Point{hmap[newPos].h, curPoses[minPos].dist + 1}
				}

			}
		}
		delete(curPoses, minPos)
		if hmap[endPos].dist != math.MaxInt {
			break
		}
	}

	fmt.Println("Day 12.1: ", hmap[endPos])
}
