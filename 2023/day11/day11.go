package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func dist(first [2]int, second [2]int) int {
	ret := 0
	ret += utils.IntAbs(second[0] - first[0])
	ret += utils.IntAbs(second[1] - first[1])
	return ret
}

func main() {
	file, err := os.Open("2023/day11/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}

	countY := func(col int) int {
		counter := 0
		for _, line := range lines {
			if line[col] == '.' {
				counter++
			}
		}
		return counter
	}
	galaxies := make(map[[2]int]string)
	y := 0
	x := 0
	startXLength := len(lines[0])
	startYLength := len(lines[1])

	//var maxY int
	for _, line := range lines {
		if strings.Count(line, ".") == startXLength {
			y++
		}
		x = 0
		for i, start := range line {
			if countY(i) == startYLength {
				x++
			}
			if start == '#' {
				galaxies[[2]int{x, y}] = string(start)
			}
			x++
		}
		y++
		//maxX = len(scanner.Text())
	}
	//maxY = y - 1

	utils.Print2DStringsGrid(galaxies)
	//fmt.Println(maxX, maxY)

	totalDist := 0
	for left, _ := range galaxies {
		counter := 0
		for right, _ := range galaxies {
			if left != right {
				totalDist += dist(left, right)
			}
			counter++
		}
	}
	fmt.Println(totalDist / 2)
}
