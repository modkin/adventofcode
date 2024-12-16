package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"strings"
)

type posType struct {
	pos  [2]int
	dir  [2]int
	cost int
}

func findMin(cost map[[2]int]posType, visited map[[2]int]bool) posType {
	minC := math.MaxInt
	var minP posType
	for _, i2 := range cost {
		if _, ok := visited[i2.pos]; !ok {
			if i2.cost < minC {
				minC = i2.cost
				minP = i2
			}
		}
	}
	return minP
}

func AddPoints(one, two [2]int) [2]int {
	return [2]int{one[0] + two[0], one[1] + two[1]}
}

func Neighbours4Pt(p [2]int, g map[[2]int]string) map[[2]int]string {
	ret := make(map[[2]int]string)
	for _, offset := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		tmp := AddPoints(p, offset)
		if g[tmp] == "" {
			ret[tmp] = g[tmp]
		}
	}
	return ret
}

func printCost(grid map[[2]int]posType) {
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
	for y := 0; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
			if _, ok := grid[[2]int{x, y}]; ok {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	lines := utils.ReadFileIntoLines("2024/day16/input")

	maze := make(map[[2]int]string)

	cost := make(map[[2]int]posType)
	visited := make(map[[2]int]bool)
	var target [2]int
	for y, line := range lines {
		split := strings.Split(line, "")
		for x, s := range split {
			if s == "S" {
				maze[[2]int{x, y}] = "."
				p := posType{dir: [2]int{1, 0}, pos: [2]int{x, y}, cost: 0}
				cost[[2]int{x, y}] = p
			} else if s == "E" {
				maze[[2]int{x, y}] = "."
				target = [2]int{x, y}
				cost[[2]int{x, y}] = posType{dir: [2]int{1, 0}, pos: [2]int{x, y}, cost: math.MaxInt}
			} else if s == "." {
				maze[[2]int{x, y}] = s
				cost[[2]int{x, y}] = posType{dir: [2]int{1, 0}, pos: [2]int{x, y}, cost: math.MaxInt}
			} else {
				maze[[2]int{x, y}] = s
			}
		}
	}

	printCost(cost)

	for minP := findMin(cost, visited); minP.pos != target; minP = findMin(cost, visited) {
		visited[minP.pos] = true
		for _, offset := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			nextPos := AddPoints(minP.pos, offset)
			if maze[nextPos] == "#" {
				continue
			}
			nextDir := offset
			turnCost := 0
			if minP.dir[0] != nextDir[0] {
				turnCost = 1000
			}
			nextCost := minP.cost + 1 + turnCost
			if cost[nextPos].cost > nextCost {
				cost[nextPos] = posType{pos: nextPos, dir: nextDir, cost: nextCost}
			}

		}
		//utils.Print2DStringGrid(visited)
	}

	fmt.Println(cost[[2]int{1, 13}])
	fmt.Println(cost[[2]int{1, 12}])
	fmt.Println(cost[[2]int{1, 11}])
	fmt.Println(cost[[2]int{1, 10}])
	fmt.Println(cost[[2]int{1, 9}])
	fmt.Println(cost[[2]int{2, 9}])

	fmt.Println(cost[target].cost)
}
