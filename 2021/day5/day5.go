package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ventline struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func main() {
	file, err := os.Open("2021/day5/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	lines := make([]ventline, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		start := strings.Split(line[0], ",")
		end := strings.Split(line[1], ",")
		x1, y1 := utils.ToInt(start[0]), utils.ToInt(start[1])
		x2, y2 := utils.ToInt(end[0]), utils.ToInt(end[1])
		lines = append(lines, ventline{x1, y1, x2, y2})
	}
	rectangularLines := make([]ventline, 0)
	diagLines := make([]ventline, 0)
	for _, line := range lines {
		if line.x1 == line.x2 || line.y1 == line.y2 {
			rectangularLines = append(rectangularLines, line)
		} else {
			diagLines = append(diagLines, line)
		}
	}
	if len(diagLines)+len(rectangularLines) != len(lines) {
		fmt.Println("error")
	}
	ocean := make(map[[2]int]int)
	for _, line := range rectangularLines {
		xStart, xEnd, yStart, yEnd := line.x1, line.x2, line.y1, line.y2
		if line.x1 > line.x2 {
			xStart = line.x2
			xEnd = line.x1
		}
		if line.y1 > line.y2 {
			yStart = line.y2
			yEnd = line.y1
		}
		for x := xStart; x <= xEnd; x++ {
			for y := yStart; y <= yEnd; y++ {
				if _, ok := ocean[[2]int{x, y}]; ok {
					ocean[[2]int{x, y}] += 1
				} else {
					ocean[[2]int{x, y}] = 1
				}
			}
		}
	}
	count := 0
	for _, i := range ocean {
		if i >= 2 {
			count++
		}
	}
	fmt.Println("Day 5.1:", count)
	for _, line := range diagLines {
		xStart, xEnd, yStart, yEnd := line.x1, line.x2, line.y1, line.y2
		xstep, ystep := 1, 1
		if line.x1 > line.x2 {
			xstep = -1
		}
		if line.y1 > line.y2 {
			ystep = -1
		}
		y := yStart
		for x := xStart; x != (xEnd+xstep) || y != (yEnd+ystep); x += xstep {
			if _, ok := ocean[[2]int{x, y}]; ok {
				ocean[[2]int{x, y}] += 1
			} else {
				ocean[[2]int{x, y}] = 1
			}
			y += ystep
		}
	}
	count = 0
	for _, i := range ocean {
		if i >= 2 {
			count++
		}
	}

	fmt.Println("Day 5.2:", count)
}
