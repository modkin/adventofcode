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

func move(guard guardPos, lab map[[2]int]string, block [2]int) guardPos {
	if guard.dir == "^" {
		newPos := [2]int{guard.pos[0], guard.pos[1] - 1}
		if _, ok := lab[newPos]; ok || newPos == block {
			guard.dir = ">"
		} else {
			guard.pos = newPos
		}
	}
	if guard.dir == ">" {
		newPos := [2]int{guard.pos[0] + 1, guard.pos[1]}
		if _, ok := lab[newPos]; ok || newPos == block {
			guard.dir = "v"
		} else {
			guard.pos = newPos
		}
	}
	if guard.dir == "v" {
		newPos := [2]int{guard.pos[0], guard.pos[1] + 1}
		if _, ok := lab[newPos]; ok || newPos == block {
			guard.dir = "<"
		} else {
			guard.pos = newPos
		}
	}
	if guard.dir == "<" {
		newPos := [2]int{guard.pos[0] - 1, guard.pos[1]}
		if _, ok := lab[newPos]; ok || newPos == block {
			guard.dir = "^"
		} else {
			guard.pos = newPos
		}
	}
	return guard
}

func getDefaultRoute(guard guardPos, labor map[[2]int]string, xMax, yMax int) (int, map[[2]int]bool) {
	visited := make(map[[2]int]bool)
	for guard.pos[0] >= 0 && guard.pos[0] < xMax && guard.pos[1] >= 0 && guard.pos[1] < yMax {
		visited[guard.pos] = true
		nonBlockBlock := [2]int{-10, -10}
		guard = move(guard, labor, nonBlockBlock)
	}
	return len(visited), visited
}

func checkLoop(guard guardPos, labor map[[2]int]string, block [2]int, xMax, yMax int, c chan bool) bool {
	loop := make(map[[2]int]string)
	for guard.pos[0] >= 0 && guard.pos[0] < xMax && guard.pos[1] >= 0 && guard.pos[1] < yMax {
		loop[guard.pos] += guard.dir
		guard = move(guard, labor, block)
		if strings.Contains(loop[guard.pos], guard.dir) {
			c <- true
			return true
		}
	}
	c <- false
	return false
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

	part1, guardPath := getDefaultRoute(guard, labor, xMax, yMax)
	fmt.Println("Day 6.1:", part1)

	counter := 0
	c := make(chan bool, len(guardPath))

	delete(guardPath, guard.pos)
	for block := range guardPath {
		go func(block [2]int) {
			checkLoop(guard, labor, block, xMax, yMax, c)
		}(block)
	}

	for range guardPath {
		if <-c {
			counter++
		}
	}

	fmt.Println("Day 6.2:", counter)
}
