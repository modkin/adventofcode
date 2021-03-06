package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2015/day6/input")
	if err != nil {
		panic(err)
	}
	grid := [1000][1000]bool{}
	brightnessGrid := [1000][1000]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " ")
		coordStart := 2
		if instruction[0] == "toggle" {
			coordStart = 1
		}
		start := strings.Split(instruction[coordStart], ",")
		distance := strings.Split(instruction[coordStart+2], ",")
		xstart, _ := strconv.Atoi(start[0])
		ystart, _ := strconv.Atoi(start[1])
		xdistance, _ := strconv.Atoi(distance[0])
		ydistance, _ := strconv.Atoi(distance[1])
		for x := xstart; x <= xdistance; x++ {
			for y := ystart; y <= ydistance; y++ {
				if instruction[0] == "turn" {
					if instruction[1] == "on" {
						grid[x][y] = true
						brightnessGrid[x][y] += 1
					} else {
						grid[x][y] = false
						if brightnessGrid[x][y] > 0 {
							brightnessGrid[x][y] -= 1
						}
					}
				} else {
					grid[x][y] = !grid[x][y]
					brightnessGrid[x][y] += 2
				}
			}
		}
	}
	count := 0
	for _, x := range grid {
		for _, y := range x {
			if y {
				count += 1
			}
		}
	}
	brightnessCount := 0
	for _, x := range brightnessGrid {
		for _, y := range x {
			brightnessCount += y
		}
	}
	fmt.Println("Task 6.1:", count)
	fmt.Println("Task 6.2:", brightnessCount)
}
