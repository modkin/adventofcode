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
	file, err := os.Open("./testInput2")
	if err != nil {
		panic(err)
	}

	dungeon := make(map[[2]int]string)
	portalMap := make(map[string][][2]int)
	jumpLevelOut := make(map[[2]int][2]int)
	jumpLevelIn := make(map[[2]int][2]int)

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

	createDungeonCopy := func() map[[2]int]string {
		dungeonCopy := make(map[[2]int]string)
		for key, val := range dungeon {
			dungeonCopy[key] = val
		}
		return dungeonCopy
	}

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

			// ?X.
			if unicode.IsLetter([]rune(west + "?")[0]) && east == "." {
				newPortal := []string{char, west}
				sort.Strings(newPortal)
				newPos := [2]int{pos[0] + 1, pos[1]}
				portalMap[strings.Join(newPortal, "")] = append(portalMap[strings.Join(newPortal, "")], newPos)
				// .X?
			} else if unicode.IsLetter([]rune(east + "?")[0]) && west == "." {
				newPortal := []string{char, east}
				sort.Strings(newPortal)
				newPos := [2]int{pos[0] - 1, pos[1]}
				portalMap[strings.Join(newPortal, "")] = append(portalMap[strings.Join(newPortal, "")], newPos)
				// ?
				// X
				// .
			} else if unicode.IsLetter([]rune(north + "?")[0]) && south == "." {
				newPortal := []string{char, north}
				sort.Strings(newPortal)
				newPos := [2]int{pos[0], pos[1] + 1}
				portalMap[strings.Join(newPortal, "")] = append(portalMap[strings.Join(newPortal, "")], newPos)
				// .
				// X
				// ?
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
			if coords[0][0] == 2 || coords[0][0] == maxX-2 || coords[0][1] == 2 || coords[0][1] == maxY-1 {
				jumpLevelOut[coords[0]] = coords[1]
			} else {
				jumpLevelIn[coords[0]] = coords[1]
			}
			if coords[1][0] == 2 || coords[1][0] == maxX-2 || coords[1][1] == 2 || coords[1][1] == maxY-1 {
				jumpLevelOut[coords[1]] = coords[0]
			} else {
				jumpLevelIn[coords[1]] = coords[0]
			}
		}
		if len(coords) != 2 {
			fmt.Println("ERROR ", portal)
		}
	}
	fmt.Println("Out: ", jumpLevelOut)
	fmt.Println("In: ", jumpLevelIn)

	dungeonsMap := make(map[int]map[[2]int]string)
	dungeonsMap[0] = createDungeonCopy()

	start := portalMap["AA"][0]
	dungeonsMap[0][start] = "0"
	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	positions := [][3]int{{start[0], start[1], 0}}
	running := true
	for running {
		running = false
		var newPositions [][3]int
		for _, pos := range positions {
			if pos[2] > 8 {
				continue
			}
			twoDpos := [2]int{pos[0], pos[1]}
			for _, dir := range directions {
				currentDist, _ := strconv.Atoi(dungeonsMap[pos[2]][twoDpos])
				twoDnewPos := utils.Sum(twoDpos, dir)
				newPos := [3]int{twoDnewPos[0], twoDnewPos[1], pos[2]}
				lookingAt := dungeonsMap[pos[2]][twoDnewPos]

				if lookingAt == "#" {
					continue
				}
				if unicode.IsLetter([]rune(lookingAt)[0]) {
					if nP, ok := jumpLevelOut[twoDpos]; ok {
						newPos[2]--
						newPos[0], newPos[1] = nP[0], nP[1]
					} else if nP, ok := jumpLevelIn[twoDpos]; ok {
						newPos[2]++
						running = true
						newPos[0], newPos[1] = nP[0], nP[1]
					} else {
						//AA or ZZ
						continue
					}
					if newPos[2] == -1 {
						continue
					}
					if _, ok := dungeonsMap[newPos[2]]; !ok {
						dungeonsMap[newPos[2]] = createDungeonCopy()
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

				dungeonsMap[newPos[2]][twoDnewPos] = fmt.Sprint(currentDist + 1)
				newPositions = append(newPositions, newPos)
			}
		}
		positions = newPositions
	}
	//fmt.Println(dungeonsMap)
	//fmt.Println(portalMap)
	printPaintMap(dungeonsMap[1])
	fmt.Println(dungeonsMap[0][portalMap["ZZ"][0]])

}
