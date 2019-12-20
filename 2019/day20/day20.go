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

func main() {
	file, err := os.Open("./testInput1")
	if err != nil {
		panic(err)
	}

	dungeon := make(map[[2]int]string)
	portalMap := make(map[string][][2]int)
	jumpMap := make(map[[2]int][2]int)

	scanner := bufio.NewScanner(file)
	x := 0
	y := 0
	for scanner.Scan() {
		x = 0
		line := strings.Split(scanner.Text(), "")
		for _, char := range line {
			dungeon[[2]int{x, y}] = char
			x++
		}
		y++
	}

	printPaintMap(dungeon)
	maxX, maxY := math.MinInt32, math.MinInt32
	for pos := range dungeon {
		if pos[0] > maxX {
			maxX = pos[0]
		}
		if pos[1] > maxY {
			maxY = pos[1]
		}
	}

	for pos, char := range dungeon {
		if pos[0] == 0 || pos[0] == maxX || pos[1] == 0 || pos[1] == maxY {
			continue
		}
		//char = X
		if char != "#" && char != "." && char != " " {
			west := dungeon[[2]int{pos[0] - 1, pos[1]}]
			east := dungeon[[2]int{pos[0] + 1, pos[1]}]
			north := dungeon[[2]int{pos[0], pos[1] - 1}]
			south := dungeon[[2]int{pos[0], pos[1] + 1}]

			//// ?X.
			if unicode.IsLetter([]rune(west + "?")[0]) && east == "." {
				newPortal := []string{char, west}
				sort.Strings(newPortal)
				newPos := [2]int{pos[0] + 1, pos[1]}
				portalMap[strings.Join(newPortal, "")] = append(portalMap[strings.Join(newPortal, "")], newPos)
			} else if unicode.IsLetter([]rune(east + "?")[0]) && west == "." {
				newPortal := []string{char, east}
				sort.Strings(newPortal)
				newPos := [2]int{pos[0] - 1, pos[1]}
				portalMap[strings.Join(newPortal, "")] = append(portalMap[strings.Join(newPortal, "")], newPos)
			} else if unicode.IsLetter([]rune(north + "?")[0]) && south == "." {
				newPortal := []string{char, north}
				sort.Strings(newPortal)
				newPos := [2]int{pos[0], pos[1] + 1}
				portalMap[strings.Join(newPortal, "")] = append(portalMap[strings.Join(newPortal, "")], newPos)
			} else if unicode.IsLetter([]rune(south + "?")[0]) && north == "." {
				newPortal := []string{char, south}
				sort.Strings(newPortal)
				newPos := [2]int{pos[0], pos[1] - 1}
				portalMap[strings.Join(newPortal, "")] = append(portalMap[strings.Join(newPortal, "")], newPos)
			}
		}
	}

	for portal, coords := range portalMap {
		if portal != "AA" && portal != "ZZ" {
			jumpMap[coords[0]] = coords[1]
			jumpMap[coords[1]] = coords[0]
		}
	}
	fmt.Println(jumpMap)

	start := portalMap["AA"][0]
	dungeon[start] = "0"
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

				if lookingAt == "#" {
					continue
				}
				if unicode.IsLetter([]rune(lookingAt)[0]) {
					if nP, ok := jumpMap[pos]; ok {
						newPos = nP
					} else {
						//AA or ZZ
						continue
					}
				}

				if lookingAtDistance, err := strconv.Atoi(lookingAt); err == nil {
					if lookingAtDistance <= currentDist+1 {
						continue
					} else {
						running = true
					}
				}
				if lookingAt == "." {
					running = true
				}

				dungeon[newPos] = fmt.Sprint(currentDist + 1)
				newPositions = append(newPositions, newPos)
			}
		}
		positions = newPositions
	}
	fmt.Println(portalMap)
	fmt.Println(dungeon[portalMap["ZZ"][0]])

}
