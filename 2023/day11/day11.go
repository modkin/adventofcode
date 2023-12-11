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

	//galaxies2 := make(map[[2]int]string)

	expandGalaxies := func(expansionFaktor int) (expGalaxy map[[2]int]string) {
		expGalaxy = make(map[[2]int]string)
		y := 0
		startXLength := len(lines[0])
		startYLength := len(lines[1])
		for _, line := range lines {
			if strings.Count(line, ".") == startXLength {
				y += expansionFaktor - 1
			}
			x := 0
			for i, start := range line {
				if countY(i) == startYLength {
					x += expansionFaktor - 1
				}
				if start == '#' {
					expGalaxy[[2]int{x, y}] = string(start)
				}
				x += 1
			}
			y++
		}
		return
	}
	galaxies1 := expandGalaxies(2)
	galaxies2 := expandGalaxies(1000000)

	calculateDistances := func(galaxy map[[2]int]string) int {
		totalDist := 0
		for left, _ := range galaxy {
			counter := 0
			for right, _ := range galaxy {
				if left != right {
					totalDist += dist(left, right)
				}
				counter++
			}
		}
		return totalDist / 2
	}
	fmt.Println("Day 11.1:", calculateDistances(galaxies1))
	fmt.Println("Day 11.2:", calculateDistances(galaxies2))
}
