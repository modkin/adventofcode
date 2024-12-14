package main

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
)

type robot struct {
	pos [2]int
	dir [2]int
}

func add2Int(one, two [2]int) [2]int {
	return [2]int{one[0] + two[0], one[1] + two[1]}
}

func moveRobot(robot robot, xMax, yMax int) [2]int {
	newPos := add2Int(robot.pos, robot.dir)
	if newPos[0] < 0 {
		newPos[0] = xMax - utils.IntAbs(newPos[0])
	}
	if newPos[0] >= xMax {
		newPos[0] = newPos[0] % xMax
	}
	if newPos[1] < 0 {
		newPos[1] = yMax - utils.IntAbs(newPos[1])
	}
	if newPos[1] >= yMax {
		newPos[1] = newPos[1] % yMax
	}
	return newPos
}

func main() {
	lines := utils.ReadFileIntoLines("2024/day14/input")

	reg := regexp.MustCompile(`p=(-?\d*),(-?\d*) v=(-?\d*),(-?\d*)`)
	robots := []robot{}
	for _, line := range lines {

		tmp := reg.FindStringSubmatch(line)
		newRobot := robot{pos: [2]int{utils.ToInt(tmp[1]), utils.ToInt(tmp[2])}, dir: [2]int{utils.ToInt(tmp[3]), utils.ToInt(tmp[4])}}
		robots = append(robots, newRobot)
	}

	xMax := 101
	yMax := 103

	for step := 0; step < 100; step++ {
		for i, _ := range robots {
			robots[i].pos = moveRobot(robots[i], xMax, yMax)
		}
	}

	allPos := make(map[[2]int]int)
	for _, r := range robots {
		allPos[r.pos]++
	}

	prod := 1

	fmt.Println(robots)

	quad := [4]int{0, 0, 0, 0}

	for _, r := range robots {
		p := r.pos
		if p[0] < xMax/2 && p[1] < yMax/2 {
			quad[0]++
		}
		if p[0] > xMax/2 && p[1] < yMax/2 {
			quad[1]++
		}
		if p[0] < xMax/2 && p[1] > yMax/2 {
			quad[2]++
		}
		if p[0] > xMax/2 && p[1] > yMax/2 {
			quad[3]++
		}
	}

	utils.Print2DIntGrid(allPos)

	fmt.Println(quad)
	fmt.Println(quad[0] * quad[1] * quad[2] * quad[3])

	fmt.Println(prod)
}
