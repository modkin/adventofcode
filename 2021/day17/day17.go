package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type targetArea struct {
	xmin int
	xmax int
	ymin int
	ymax int
}

func checkTarget(x int, y int, target targetArea) bool {
	if x <= target.xmax && x >= target.xmin && y <= target.ymax && y >= target.ymin {
		return true
	}
	return false
}

func fire(xSpeed int, ySpeed int, target targetArea) (bool, int) {
	x, y, ymax := 0, 0, 0
	for y >= target.ymin {
		x += xSpeed
		y += ySpeed
		if y > ymax {
			ymax = y
		}
		if checkTarget(x, y, target) {
			return true, ymax
		}
		ySpeed--
		if xSpeed > 0 {
			xSpeed--
		} else if xSpeed < 0 {
			xSpeed++
		}
	}
	return false, ymax
}

func main() {
	file, err := os.Open("2021/day17/input")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")
	xCoords := strings.Split(strings.Trim(line[2], ","), "..")
	yCoords := strings.Split(line[3], "..")
	xCoords[0] = strings.Trim(xCoords[0], "x=")
	yCoords[0] = strings.Trim(yCoords[0], "y=")
	target := targetArea{utils.ToInt(xCoords[0]), utils.ToInt(xCoords[1]), utils.ToInt(yCoords[0]), utils.ToInt(yCoords[1])}

	ymax := 0
	hits := 0
	for x := -1000; x < 1000; x++ {
		for y := -1000; y < 1000; y++ {
			hit, height := fire(x, y, target)
			if hit {
				if height > ymax {
					ymax = height
				}
				hits += 1
			}
		}
	}
	fmt.Println(ymax)
	fmt.Println(hits)

}
