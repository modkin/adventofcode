package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getDir(start, end [2]int) [2]int {
	return [2]int{end[0] - start[0], end[1] - start[1]}
}

func addVec(start, end [2]int) [2]int {
	return [2]int{start[0] + end[0], start[1] + end[1]}
}

func main() {
	file, err := os.Open("2024/day8/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	antennas := make(map[string]map[[2]int]bool)

	xMax := 0
	yMax := 0
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		xMax = len(split)
		for x, s := range split {
			if s != "." {
				pos := [2]int{x, y}
				if _, ok := antennas[s]; !ok {
					antennas[s] = make(map[[2]int]bool)
				}
				antennas[s][pos] = true
			}
		}
		y++
	}
	yMax = y

	antiNodes := make(map[[2]int]bool)
	for _, allPos := range antennas {
		for one, _ := range allPos {
			for two, _ := range allPos {
				if one != two {
					dir := getDir(one, two)
					antiPos := addVec(addVec(one, dir), dir)
					if antiPos[0] >= 0 && antiPos[0] < xMax && antiPos[1] >= 0 && antiPos[1] < yMax {
						antiNodes[antiPos] = true
					}
				}
			}
		}
	}
	fmt.Println(len(antiNodes))
}
