package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
)

func rotate90left(vec [2]int) (ret [2]int) {
	ret[0] = -1 * vec[1]
	ret[1] = vec[0]
	return
}

func rotate90right(vec [2]int) (ret [2]int) {
	ret[0] = vec[1]
	ret[1] = -1 * vec[0]
	return
}

func main() {

	pos := [2]int{0, 0}
	pos2 := [2]int{0, 0}
	waypoint := [2]int{10, 1}

	scanner := bufio.NewScanner(utils.OpenFile("2020/day12/input"))
	currentDirection := [2]int{1, 0}
	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0]
		distance := utils.ToInt(line[1:])
		if dir == 'F' {
			pos[0], pos[1] = pos[0]+distance*currentDirection[0], pos[1]+distance*currentDirection[1]
			pos2[0], pos2[1] = pos2[0]+distance*waypoint[0], pos2[1]+distance*waypoint[1]
		} else if dir == 'N' {
			pos[1] += distance
			waypoint[1] += distance
		} else if dir == 'S' {
			pos[1] -= distance
			waypoint[1] -= distance
		} else if dir == 'E' {
			pos[0] += distance
			waypoint[0] += distance
		} else if dir == 'W' {
			pos[0] -= distance
			waypoint[0] -= distance
		} else if dir == 'L' {
			for i := 0; i < distance; i += 90 {
				currentDirection = rotate90left(currentDirection)
				waypoint = rotate90left(waypoint)
			}
		} else if dir == 'R' {
			for i := 0; i < distance; i += 90 {
				currentDirection = rotate90right(currentDirection)
				waypoint = rotate90right(waypoint)
			}
		}
	}
	fmt.Println("Task 12.1:", utils.IntAbs(pos[0])+utils.IntAbs(pos[1]))
	fmt.Println("Task 12.2", utils.IntAbs(pos2[0])+utils.IntAbs(pos2[1]))
}
