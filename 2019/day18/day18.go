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
				} else {
					keyDistance[lookingAt] = currentDist + 1
				}
			}
		}
		positions = newPositions
	}
	return keyDistance
}

type Destination struct {
	dependencies []string
	distance     int
}

type Key struct {
	destinations map[string]Destination
	pos          [2]int
}

func main() {
	file, err := os.Open("./testInput2")
	if err != nil {
		panic(err)
	}

	dungeon := make(map[[2]int]string)
	keyMap := make(map[string]Key)

	scanner := bufio.NewScanner(file)
	x := 0
	y := 0
	for scanner.Scan() {
		x = 0
		line := strings.Split(scanner.Text(), "")
		for _, char := range line {
			dungeon[[2]int{x, y}] = char
			if char != "#" && char != "." {
				keyMap[char] = Key{
					destinations: make(map[string]Destination),
					pos:          [2]int{x, y},
				}
			}
			x++
		}
		y++
	}
	dungeonCopy := make(map[[2]int]string)
	for key, val := range dungeon {
		dungeonCopy[key] = val
	}
	for keyName, key := range keyMap {
		dungeon[key.pos] = "0"
		distances := calcDistances(dungeon, key.pos)
		for symbol, dist := range distances {
			newDest := Destination{
				dependencies: nil,
				distance:     dist,
			}
			key.destinations[symbol] = newDest
		}

		keyMap[keyName] = key
		for key, val := range dungeonCopy {
			dungeon[key] = val
		}
	}

	for i := 0; i < 10; i++ {
		newKeyMap := make(map[string]Key)
		for symbol, keyStruct := range keyMap {
			newKeyStruct := Key{
				destinations: make(map[string]Destination),
				pos:          keyStruct.pos,
			}
			for dstName, dst := range keyStruct.destinations {
				newDeps := make([]string, len(dst.dependencies))
				copy(newDeps, dst.dependencies)
				if unicode.IsUpper([]rune(dstName)[0]) {
					newDeps = append(newDeps, dstName)
					//delete(keyStruct.destinations, dstName)
				}
				newKeyStruct.destinations[dstName] = Destination{
					dependencies: newDeps,
					distance:     dst.distance,
				}

				for indirectName, indirectdst := range keyMap[dstName].destinations {
					if _, ok := keyStruct.destinations[indirectName]; !ok && indirectName != symbol {
						newKeyStruct.destinations[indirectName] = Destination{
							dependencies: append(newDeps, indirectdst.dependencies...),
							distance:     indirectdst.distance + dst.distance,
						}
					}
				}
			}
			newKeyMap[symbol] = newKeyStruct
		}
		keyMap = newKeyMap
	}

	var deleteKeys []string
	for key, keyStruct := range keyMap {
		if unicode.IsUpper([]rune(key)[0]) {
			deleteKeys = append(deleteKeys, key)
		} else {
			for symbol, _ := range keyStruct.destinations {
				if unicode.IsUpper([]rune(symbol)[0]) {
					delete(keyStruct.destinations, symbol)
				}
			}
			keyMap[key] = keyStruct
		}
	}

	for _, key := range deleteKeys {
		delete(keyMap, key)
	}

	printPaintMap(dungeon)
	fmt.Println(keyMap)
	fmt.Println(keyMap["@"])

	possiblePath := make(map[string]int)

	possiblePath["@"] = 0

	running := true
	for running {
		running = false
		newPossiblePath := make(map[string]int)

		for keyPath, distance := range possiblePath {
			for symbol, _ := range keyMap {
				if !strings.Contains(keyPath, strings.ToLower(symbol)) {
					running = true
					break
				}
			}
			allKeys := strings.Split(keyPath, "")

			nextPoints := keyMap[allKeys[len(allKeys)-1]].destinations
			for newPos, newDest := range nextPoints {
				if strings.Contains(keyPath, newPos) {
					continue
					/// check if newPos is upper case => door
				}
				allDeps := true
				for _, dep := range newDest.dependencies {
					if !strings.Contains(keyPath, strings.ToLower(dep)) {
						allDeps = false
					}
				}
				if allDeps {
					newPossiblePath[keyPath+newPos] = distance + newDest.distance
				}
			}
		}
		if len(newPossiblePath) != 0 {
			possiblePath = newPossiblePath
		}
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
		fmt.Println("Paths: ", len(possiblePath))
		fmt.Println("Removed ", counter)

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
