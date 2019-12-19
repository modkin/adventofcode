package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
			if char != "#" && char != "." {
				keyMap[char] = [2]int{x, y}
			}
			x++
		}
		y++
	}
	dungeonCopy := make(map[[2]int]string)
	for key, val := range dungeon {
		dungeonCopy[key] = val
	}
	dungeon[keyMap["@"]] = "0"
	printPaintMap(dungeon)
	keyDistance := calcDistances(dungeon, keyMap["@"])
	printPaintMap(dungeon)
	fmt.Println(keyDistance)

	possiblePath := make(map[string]int)
	for key, dis := range keyDistance {
		possiblePath[key] = dis
	}
	running := true
	for running {
		running = false
		newPossiblePath := make(map[string]int)

		for keys, distance := range possiblePath {
			//reset dungeon
			for key, val := range dungeonCopy {
				dungeon[key] = val
			}
			dungeon[keyMap["@"]] = "."
			if len(keys) < len(keyMap)/2 {
				running = true
			}
			allKeys := strings.Split(keys, "")
			for _, key := range allKeys {
				dungeon[keyMap[key]] = "."
				dungeon[keyMap[strings.ToUpper(key)]] = "."
			}
			startPoint := keyMap[allKeys[len(allKeys)-1]]
			dungeon[startPoint] = "0"
			keyDistance := calcDistances(dungeon, startPoint)
			for newKey, newDist := range keyDistance {
				newPossiblePath[keys+newKey] = distance + newDist
			}
		}
		if len(newPossiblePath) != 0 {
			possiblePath = newPossiblePath
			counter := 0
			duplicates := make(map[string]string)
			for keys, _ := range possiblePath {
				keysSplit := strings.Split(keys, "")
				lastKey := keysSplit[len(keysSplit)-1]
				keysSplit = keysSplit[0 : len(keysSplit)-1]
				sort.Strings(keysSplit)
				sortedKeys := strings.Join(keysSplit, "")
				if dupKey, ok := duplicates[sortedKeys]; ok {
					if dupKey == lastKey {
						delete(possiblePath, keys)
						counter++
					}
				} else {
					duplicates[sortedKeys] = lastKey
				}
			}
			min := math.MaxInt32
			for _, dis := range possiblePath {
				if dis < min {
					min = dis
				}
			}
			for keys, distance := range possiblePath {
				if distance > int(float64(min)*1.7) {
					delete(possiblePath, keys)
					counter++
				}
			}
			fmt.Println("Paths: ", len(possiblePath))
			fmt.Println("Removed ", counter)
		}
	}
	min := math.MaxInt32
	for _, dis := range possiblePath {
		if dis < min {
			min = dis
		}
	}
	fmt.Println(possiblePath)
	fmt.Println(min)
}
