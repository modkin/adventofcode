package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
)

type cur struct {
	pos  [2]int
	cost int
	dir  [2]int
}
type cachePos struct {
	pos       [2]int
	dir       [2]int
	stepsLeft int
}

func turnLeft(vec [2]int) (ret [2]int) {
	ret[0] = vec[1]
	ret[1] = -1 * vec[0]
	return
}

func turnRight(vec [2]int) (ret [2]int) {
	ret[0] = -1 * vec[1]
	ret[1] = vec[0]
	return
}

func main() {
	file, err := os.Open("2023/day17/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	lavaMap := make(map[[2]int]int)
	y := 0
	var maxX int
	var maxY int
	for scanner.Scan() {
		for x, cost := range scanner.Text() {
			lavaMap[[2]int{x, y}] = utils.ToInt(string(cost))
		}
		y++
		maxX = len(scanner.Text()) - 1
	}
	maxY = y - 1
	fmt.Println(maxX, maxY)
	var allPositions []cur
	allPositions = append(allPositions, cur{[2]int{0, 0}, 0, [2]int{1, 0}})
	allPositions = append(allPositions, cur{[2]int{0, 0}, 0, [2]int{0, 1}})
	minCost := make(map[cachePos]int)
	minCost[cachePos{[2]int{0, 0}, [2]int{1, 0}, 3}] = 0
	minCost[cachePos{[2]int{0, 0}, [2]int{0, 1}, 3}] = 0
	for len(allPositions) != 0 {
		var newAllPos []cur
		for _, current := range allPositions {
			var newDirs [][2]int
			newDirs = append(newDirs, turnLeft(current.dir))
			newDirs = append(newDirs, turnRight(current.dir))
			for _, nextDir := range newDirs {

				costTmp := 0

				for steps := 1; steps <= 3; steps++ {
					stepsLeft := 3 - steps
					nextPos := [2]int{current.pos[0] + steps*nextDir[0], current.pos[1] + steps*nextDir[1]}
					if nextPos[0] < 0 || nextPos[0] > maxX || nextPos[1] < 0 || nextPos[1] > maxY {
						continue
					}
					costTmp += lavaMap[nextPos]
					nextCost := current.cost + costTmp
					if value, ok := minCost[cachePos{nextPos, nextDir, stepsLeft}]; ok {
						if nextCost < value {
							newAllPos = append(newAllPos, cur{nextPos, nextCost, nextDir})
							minCost[cachePos{nextPos, nextDir, stepsLeft}] = nextCost
						}
					} else {
						newAllPos = append(newAllPos, cur{nextPos, nextCost, nextDir})
						minCost[cachePos{nextPos, nextDir, stepsLeft}] = nextCost
					}

				}
			}
		}
		allPositions = newAllPos
	}

	utils.Print2DIntGrid(lavaMap)

	for pos, ints := range minCost {
		if pos.pos == [2]int{6, 0} {
			fmt.Println("pos", pos)
			fmt.Println(ints)
			fmt.Println("next")
		}

	}

	minDist := math.MaxInt
	for pos, ints := range minCost {
		if pos.pos == [2]int{maxX, maxY} {
			if ints < minDist {
				minDist = ints
			}
		}
	}
	fmt.Println(minDist)

}
