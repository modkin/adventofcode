package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type guardPos struct {
	pos [2]int
	dir string
}

func move(guard guardPos, lab map[[2]int]string) guardPos {
	if guard.dir == "^" {
		newPos := [2]int{guard.pos[0], guard.pos[1] - 1}
		if _, ok := lab[newPos]; ok {
			guard.dir = ">"
		} else {
			guard.pos = newPos
		}
	}
	if guard.dir == ">" {
		newPos := [2]int{guard.pos[0] + 1, guard.pos[1]}
		if _, ok := lab[newPos]; ok {
			guard.dir = "v"
		} else {
			guard.pos = newPos
		}
	}
	if guard.dir == "v" {
		newPos := [2]int{guard.pos[0], guard.pos[1] + 1}
		if _, ok := lab[newPos]; ok {
			guard.dir = "<"
		} else {
			guard.pos = newPos
		}
	}
	if guard.dir == "<" {
		newPos := [2]int{guard.pos[0] - 1, guard.pos[1]}
		if _, ok := lab[newPos]; ok {
			guard.dir = "^"
		} else {
			guard.pos = newPos
		}
	}
	return guard
}

func getLeavingSteps(guard guardPos, labor map[[2]int]string, xMax, yMax int) int {
	visited := make(map[[2]int]bool)
	loop := make(map[[2]int]string)
	for guard.pos[0] >= 0 && guard.pos[0] < xMax && guard.pos[1] >= 0 && guard.pos[1] < yMax {
		loop[guard.pos] += guard.dir
		visited[guard.pos] = true
		guard = move(guard, labor)
		if strings.Contains(loop[guard.pos], guard.dir) {
			return 0
		}
	}
	return len(visited)
}

func main() {
	file, err := os.Open("2024/day6/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	labor := make(map[[2]int]string)
	var guard guardPos

	y := 0
	xMax := 0
	yMax := 0
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, "")
		xMax = len(split)
		for x, i2 := range split {
			if i2 == "#" {
				labor[[2]int{x, y}] = "#"
			} else if i2 == "^" {
				guard = guardPos{[2]int{x, y}, "^"}
			}
		}
		y++
	}
	yMax = y

	fmt.Println("Day 6.1:", getLeavingSteps(guard, labor, xMax, yMax))

	counter := 0
	for x := 0; x < xMax; x++ {
		for y = 0; y < yMax; y++ {
			if !(x == guard.pos[0] && y == guard.pos[1]) {
				if _, ok := labor[[2]int{x, y}]; !ok {
					labor[[2]int{x, y}] = "O"
					//utils.Print2DStringsGrid(labor)
					steps := getLeavingSteps(guard, labor, xMax, yMax)
					if steps == 0 {
						counter++

					}
					delete(labor, [2]int{x, y})
				}
			}
		}
	}
	fmt.Println("Day 6.2:", counter)
}
