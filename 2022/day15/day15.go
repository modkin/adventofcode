package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type sensor struct {
	pos  [2]int
	dist int
}

func checkLine(sensors []sensor, lineNum int, beacons []int) int {
	elements := make(map[int]bool)
	for _, s := range sensors {
		xdist := s.dist - utils.IntAbs(lineNum-s.pos[1])
		if xdist > -1 {
			for x := s.pos[0] - xdist; x <= s.pos[0]+xdist; x++ {
				if !utils.IntSliceContains(beacons, x) {
					elements[x] = true
				}
			}
		}
	}
	return len(elements)
}

func inRange(x, y int, sen sensor) bool {
	xdist := utils.IntAbs(x - sen.pos[0])
	ydist := utils.IntAbs(y - sen.pos[1])
	if xdist+ydist <= sen.dist {
		return true
	} else {
		return false
	}
}

func genBoundary(sen sensor) [][2]int {
	ret := make([][2]int, 0)
	boundaryDist := sen.dist + 1
	for dist := 0; dist <= boundaryDist; dist++ {

		ret = append(ret, [2]int{sen.pos[0] + dist, sen.pos[1] + boundaryDist - dist})
		ret = append(ret, [2]int{sen.pos[0] - dist, sen.pos[1] - boundaryDist + dist})
		ret = append(ret, [2]int{sen.pos[0] + boundaryDist - dist, sen.pos[1] - dist})
		ret = append(ret, [2]int{sen.pos[0] - boundaryDist + dist, sen.pos[1] + dist})
	}
	return ret
}

func checkSensor(allSensors []sensor, sen sensor) (found bool, x, y int) {
	boundary := genBoundary(sen)
outer:
	for _, b := range boundary {
		for _, other := range allSensors {
			if sen.pos == other.pos {
				continue
			}
			if inRange(b[0], b[1], other) {
				continue outer
			}
		}
		return true, b[0], b[1]
	}
	return false, 0, 0
}

func main() {

	file, err := os.Open("2022/day15/input")
	var yCoord int
	if strings.Contains(file.Name(), "test") {
		yCoord = 20
	} else {
		yCoord = 4000000
	}
	if err != nil {
		panic(err)
	}
	sensors := make([]sensor, 0)
	beaconsOnLine := make([]int, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		xSen := utils.ToInt(strings.Split(strings.Trim(line[2], ","), "=")[1])
		ySen := utils.ToInt(strings.Split(strings.Trim(line[3], ":"), "=")[1])
		xBe := utils.ToInt(strings.Split(strings.Trim(line[8], ","), "=")[1])
		yBe := utils.ToInt(strings.Split(line[9], "=")[1])
		xdist := utils.IntAbs(xBe - xSen)
		ydist := utils.IntAbs(yBe - ySen)
		totaldist := xdist + ydist
		newSen := sensor{[2]int{xSen, ySen}, totaldist}
		sensors = append(sensors, newSen)
		if yBe == yCoord/2 {
			beaconsOnLine = append(beaconsOnLine, xBe)
		}
	}

	fmt.Println("Day 15.1:", checkLine(sensors, yCoord/2, beaconsOnLine))

	for _, s := range sensors {
		ret, x, y := checkSensor(sensors, s)
		if ret {
			if 0 <= x && x <= yCoord && 0 <= y && y <= yCoord {
				fmt.Println("Day 15.2:", x*4000000+y)
				break
			}
		}
	}
}
