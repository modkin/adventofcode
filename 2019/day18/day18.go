package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func printPaintMap(paintMap map[[2]int]string) {
	maxX, maxY := math.MinInt32, math.MinInt32
	for pos, _ := range paintMap {
		if pos[0] > maxX {
			maxX = pos[0]
		}

		if pos[1] > maxY {
			maxY = pos[1]
		}
	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			fmt.Print(paintMap[[2]int{x, y}])
		}
		fmt.Println()
	}
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	dungeon := make(map[[2]int]string)
	keyMap := make(map[string][2]int)
	scanner := bufio.NewScanner(file)
	x := 0
	y := 0
	for scanner.Scan() {
		x = 0
		line := strings.Split(scanner.Text(), "")
		for _, char := range line {
			dungeon[[2]int{x, y}] = char
			x++
			if char != "#" && char != "." {
				keyMap[char] = [2]int{x, y}
			}
		}
		y++
	}

	fmt.Println(keyMap)
	//printPaintMap(dungeon)
}
