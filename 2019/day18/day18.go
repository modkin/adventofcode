package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func printPaintMap(paintMap map[[2]int]string) {
	maxX, maxY := math.MinInt32, math.MinInt32
	for pos := range paintMap {
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

func calcDistances(dungeon map[[2]int]string, start [2]int) map[string]int {
	keyDistance := make(map[string]int)
	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	positions := [][2]int{start}
	running := true
	for running {
		running = false
		var newPositions [][2]int
		for _, pos := range positions {
			for _, dir := range directions {
				currentDist, _ := strconv.Atoi(dungeon[pos])
				newPos := utils.Sum(pos, dir)
				lookingAt := dungeon[newPos]
				isNumber := true
				if _, err := strconv.Atoi(lookingAt); err != nil {
					isNumber = false
				}
				if lookingAt == "#" || isNumber {
					continue
				} else if lookingAt == "." {
					running = true
					dungeon[newPos] = fmt.Sprint(currentDist + 1)
					newPositions = append(newPositions, newPos)
				} else if unicode.IsLower([]rune(lookingAt)[0]) {
					keyDistance[lookingAt] = currentDist + 1
				}
			}
		}
		positions = newPositions
	}
	return keyDistance
}

func main() {
	file, err := os.Open("./testInput1")
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
			if char != "#" && char != "." {
				keyMap[char] = [2]int{x, y}
			}
			x++
		}
		y++
	}
	dungenCopy := make(map[[2]int]string)
	for key, val := range dungeon {
		dungenCopy[key] = val
	}
	dungeon[keyMap["@"]] = "0"
	printPaintMap(dungeon)
	keyDistance := calcDistances(dungeon, keyMap["@"])
	printPaintMap(dungeon)
	fmt.Println(keyDistance)
}
