package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"math/bits"
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
func find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
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

func keyToUint32(key string) (ret uint32) {
	pos := int([]rune(key)[0]) - 97
	ret |= 1 << pos
	return
}

type Destination struct {
	keysOnTheWay []string
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
	counterLowerCaseLetter := 0
	var keyList []string
	for scanner.Scan() {
		x = 0
		line := strings.Split(scanner.Text(), "")
		for _, char := range line {
			if unicode.IsLower([]rune(char + ".")[0]) {
				counterLowerCaseLetter++
				keyList = append(keyList, char)
			}
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
	fmt.Println("Number of keys: ", counterLowerCaseLetter)
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

	var total uint32
	for _, key := range keyList {
		fmt.Printf("%s == %032b\n", key, keyToUint32(key))
		total = total | keyToUint32(key)
	}
	fmt.Printf("%d== %032b\n", bits.OnesCount32(total), total)

	running := true
	for running {
		running = false
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
					keysOnTheWay: utils.CopyStringSlice(dst.keysOnTheWay),
					dependencies: newDeps,
					distance:     dst.distance,
				}

				for indirectName, indirectdst := range keyMap[dstName].destinations {
					if _, ok := keyStruct.destinations[indirectName]; !ok && indirectName != symbol {
						newKeysOnTheWay := append(utils.CopyStringSlice(dst.keysOnTheWay), indirectdst.keysOnTheWay...)
						if unicode.IsLower([]rune(dstName)[0]) {
							newKeysOnTheWay = append(newKeysOnTheWay, dstName)
						}
						newKeyStruct.destinations[indirectName] = Destination{
							keysOnTheWay: newKeysOnTheWay,
							dependencies: append(newDeps, indirectdst.dependencies...),
							distance:     indirectdst.distance + dst.distance,
						}
					}
				}
			}
			newKeyMap[symbol] = newKeyStruct
		}
		keyMap = newKeyMap
		for key, destinations := range keyMap {
			if unicode.IsLower([]rune(key)[0]) {
				for _, allKey := range keyList {
					if allKey != key {
						if _, ok := destinations.destinations[allKey]; !ok {
							running = true
							break
						}
					}
				}
			}
		}
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

	//printPaintMap(dungeon)
	//fmt.Println(keyMap)
	//fmt.Println(keyMap["@"])

	type position struct {
		distance int
		keys     uint32
	}

	//possiblePath := make(map[string]position)
	//
	//possiblePath["@"] = position{
	//	distance: 0,
	//	keys:     0,
	//}
	possiblePath := make(map[string]int)
	possiblePath["@"] = 0

	var min int
	var minPath string
	running = true
	var cantReach []string
	for running {
		running = false
		/// find shortest path
		min = math.MaxInt32
		for path, dis := range possiblePath {
			if !find(cantReach, path) {
				if dis < min {
					min = dis
					minPath = path
				}
			}
		}

		fmt.Println(min)
		fmt.Println(minPath)

		for _, keyToCheck := range keyList {
			if !strings.Contains(minPath, strings.ToLower(keyToCheck)) {
				running = true
				break
			}
		}

		allKeys := strings.Split(minPath, "")

		nextPoints := keyMap[allKeys[len(allKeys)-1]].destinations
		cantProgress := true
		for newPos, newDest := range nextPoints {
			/// skip self
			if strings.Contains(minPath, newPos) {
				continue
				/// check if newPos is upper case => door
			}
			allDeps := true
			for _, dep := range newDest.dependencies {
				if !strings.Contains(minPath, strings.ToLower(dep)) {
					allDeps = false
					break
				}
			}
			if allDeps {
				possiblePath[minPath+strings.Join(newDest.keysOnTheWay, "")+newPos] = min + newDest.distance
				delete(possiblePath, minPath)
				cantProgress = false
				cantReach = nil
			}
		}
		if cantProgress {
			cantReach = append(cantReach, minPath)
		}
		//counter := 0
		type keyDist struct {
			key      string
			distance int
			orig     string
		}
		duplicates := make(map[string]keyDist)
		for keys, dis := range possiblePath {
			keysSplit := strings.Split(keys, "")
			lastKey := keysSplit[len(keysSplit)-1]
			keysSplit = keysSplit[0 : len(keysSplit)-1]
			sort.Strings(keysSplit)
			sortedKeys := strings.Join(keysSplit, "")
			if dupKey, ok := duplicates[sortedKeys]; ok {
				if dupKey.key == lastKey {
					delete(possiblePath, keys)
					if dis < dupKey.distance {
						possiblePath[dupKey.orig] = dis
					}
				}
			} else {
				duplicates[sortedKeys] = keyDist{
					key:      lastKey,
					distance: dis,
					orig:     keys,
				}
			}
		}
	}
	//fmt.Println("Paths: ", len(possiblePath))
	//fmt.Println("Removed ", counter)

	fmt.Println(minPath)
	fmt.Println(min)
}
