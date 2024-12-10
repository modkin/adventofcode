package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func step(pos, dir [2]int) [2]int {
	return [2]int{pos[0] + dir[0], pos[1] + dir[1]}
}

func findNextStep(hikeMap map[[2]int]int, pos [][2]int) (ret [][2]int) {
	allDir := [][2]int{[2]int{-1, 0}, [2]int{1, 0}, [2]int{0, -1}, [2]int{0, 1}}
	for _, dir := range allDir {
		nextPos := step(dir, pos[len(pos)-1])
		if val, ok := hikeMap[nextPos]; ok {
			if val == hikeMap[pos[len(pos)-1]]+1 {
				ret = append(ret, nextPos)
			}
		}

	}
	return ret

}

func copyPath(in [][2]int) [][2]int {
	ret := [][2]int{}
	for _, ints := range in {
		ret = append(ret, ints)
	}
	return ret
}

func main() {
	lines := utils.ReadFileIntoLines("2024/day10/input")

	//xMax := len(lines[0])
	//yMax := len(lines)

	hikeMap := map[[2]int]int{}

	trailheads := map[[2]int]int{}
	for y, line := range lines {
		for x, pos := range strings.Split(line, "") {
			var stepP int
			if pos == "." {
				stepP = -1
			} else {
				stepP = utils.ToInt(pos)
			}
			hikeMap[[2]int{x, y}] = stepP
			if stepP == 0 {
				trailheads[[2]int{x, y}] = 0
			}
		}
	}
	allPath := [][][2]int{}
	for ints, _ := range trailheads {
		tmp := [][2]int{ints}
		allPath = append(allPath, tmp)

	}

	for i := 0; i < 9; i++ {
		newAllPath := [][][2]int{}
		for _, cur := range allPath {

			for _, newPos := range findNextStep(hikeMap, cur) {
				newPath := append(copyPath(cur), newPos)
				newAllPath = append(newAllPath, newPath)
			}

		}
		allPath = newAllPath
		for _, i3 := range allPath {
			if i3[0] == [2]int{2, 0} {
				fmt.Println("i", i3)
			}
		}
	}

	scores := make(map[[2]int]map[[2]int]int)
	for i, _ := range trailheads {
		scores[i] = make(map[[2]int]int)
	}
	for _, b := range allPath {
		scores[b[0]][b[9]]++
	}

	sum := 0
	for _, i := range scores {
		sum += len(i)
	}
	fmt.Println("Day 10.1:", sum)

	fmt.Println("Day 10.2:", len(allPath))

}
