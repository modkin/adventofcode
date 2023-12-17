package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
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

type byCost []cur

func (c byCost) Len() int           { return len(c) }
func (c byCost) Less(i, j int) bool { return c[i].cost < c[j].cost }
func (c byCost) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

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

	findShortes := func(minSteps int, maxSteps int) int {
		var allPositions []cur
		allPositions = append(allPositions, cur{[2]int{0, 0}, 0, [2]int{1, 0}})
		allPositions = append(allPositions, cur{[2]int{0, 0}, 0, [2]int{0, 1}})
		minCost := make(map[cachePos]int)
		minCost[cachePos{[2]int{0, 0}, [2]int{1, 0}, 10}] = 0
		minCost[cachePos{[2]int{0, 0}, [2]int{0, 1}, 10}] = 0
		for {
			var newAllPos []cur
			current := allPositions[0]

			var newDirs [][2]int
			newDirs = append(newDirs, turnLeft(current.dir))
			newDirs = append(newDirs, turnRight(current.dir))
			for _, nextDir := range newDirs {

				costTmp := 0
				for steps := 1; steps < minSteps; steps++ {
					nextPos := [2]int{current.pos[0] + steps*nextDir[0], current.pos[1] + steps*nextDir[1]}
					if nextPos[0] < 0 || nextPos[0] > maxX || nextPos[1] < 0 || nextPos[1] > maxY {
						continue
					}
					costTmp += lavaMap[nextPos]
				}

				for steps := minSteps; steps <= maxSteps; steps++ {
					stepsLeft := 10 - steps
					nextPos := [2]int{current.pos[0] + steps*nextDir[0], current.pos[1] + steps*nextDir[1]}
					if nextPos[0] < 0 || nextPos[0] > maxX || nextPos[1] < 0 || nextPos[1] > maxY {
						continue
					}
					costTmp += lavaMap[nextPos]
					nextCost := current.cost + costTmp
					value, exists := minCost[cachePos{nextPos, nextDir, stepsLeft}]
					if !exists || nextCost < value {
						newAllPos = append(newAllPos, cur{nextPos, nextCost, nextDir})
						minCost[cachePos{nextPos, nextDir, stepsLeft}] = nextCost
					}
				}
			}
			allPositions = append(allPositions[1:], newAllPos...)
			sort.Sort(byCost(allPositions))
			if allPositions[0].pos[0] == maxX && allPositions[0].pos[1] == maxY {
				return allPositions[0].cost
			}

		}
	}

	fmt.Println("Day 17.1:", findShortes(1, 3))
	fmt.Println("Day 17.2:", findShortes(4, 10))

}
