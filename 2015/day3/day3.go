package main

import (
	"fmt"
	"io/ioutil"
)

func move(step byte) (int, int) {
	x := 0
	y := 0
	if step == '^' {
		y += 1
	} else if step == 'v' {
		y -= 1
	} else if step == '<' {
		x -= 1
	} else if step == '>' {
		x += 1
	}
	return x, y
}

func main() {
	content, err := ioutil.ReadFile("2015/day3/input")
	if err != nil {
		panic(err)
	}

	houses := make(map[[2]int]int)
	housesSanta := make(map[[2]int]int)
	housesRobot := make(map[[2]int]int)
	x := 0
	y := 0
	xSanta := 0
	ySanta := 0
	xRobot := 0
	yRobot := 0
	santa := true
	houses[[2]int{x, y}] = 1
	housesSanta[[2]int{x, y}] = 1
	housesRobot[[2]int{x, y}] = 1
	for _, elem := range content {
		xstep, ystep := move(elem)
		x += xstep
		y += ystep
		if santa {
			xSanta += xstep
			ySanta += ystep
		} else {
			xRobot += xstep
			yRobot += ystep
		}
		if _, ok := houses[[2]int{x, y}]; ok {
			houses[[2]int{x, y}] += 1
		} else {
			houses[[2]int{x, y}] = 1
		}
		if santa {
			if _, ok := housesSanta[[2]int{xSanta, ySanta}]; ok {
				housesSanta[[2]int{xSanta, ySanta}] += 1
			} else {
				housesSanta[[2]int{xSanta, ySanta}] = 1
			}
		} else {
			if _, ok := housesRobot[[2]int{xRobot, yRobot}]; ok {
				housesRobot[[2]int{xRobot, yRobot}] += 1
			} else {
				housesRobot[[2]int{xRobot, yRobot}] = 1
			}
		}
		santa = !santa

	}
	onlyRobot := 0
	for key := range housesRobot {
		if _, ok := housesSanta[key]; !ok {
			onlyRobot += 1
		}
	}
	fmt.Println("Task 3.1:", len(houses))
	fmt.Println("Task 3.2:", len(housesSanta)+onlyRobot)
}
