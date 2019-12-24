package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printMap(bugMap [5][5]bool) {
	fmt.Println("=============")
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if bugMap[x][y] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("=============")
}

func calcBioDiff(bugMap [5][5]bool) (bioDiff int64) {
	bioPoints := int64(1)
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if bugMap[x][y] {
				bioDiff += bioPoints
			}
			bioPoints *= 2
		}
	}
	return
}

type recBugMap struct {
	bugMap [5][5]bool
	inner  *recBugMap
	outer  *recBugMap
}

func updateBugMap(currentMap *recBugMap) {
	bugMap := currentMap.bugMap
	var newBugMap [5][5]bool
	hasOuter := true
	if currentMap.outer == nil {
		hasOuter = false
	}
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			bugCount := 0
			if x == 0 {
				if hasOuter {
					if currentMap.outer.bugMap[2][3] {
						bugCount++
					}
				}
			} else if bugMap[x-1][y] {
				bugCount++
			}

			if x == 4 {
				if hasOuter {
					if currentMap.outer.bugMap[4][3] {
						bugCount++
					}
				}
			} else if bugMap[x+1][y] {
				bugCount++
			}

			if y == 0 {
				if hasOuter {
					if currentMap.outer.bugMap[3][2] {
						bugCount++
					}
				}
			} else if bugMap[x][y-1] {
				bugCount++
			}

			if y == 4 {
				if hasOuter {
					if currentMap.outer.bugMap[3][4] {
						bugCount++
					}
				}
			} else if bugMap[x][y+1] {
				bugCount++
			}

			if bugMap[x][y] {
				// bug
				if bugCount != 1 {
					// dies if there is one around
					newBugMap[x][y] = false
				} else {
					newBugMap[x][y] = true
				}
			} else {
				// empty
				if bugCount == 1 || bugCount == 2 {
					newBugMap[x][y] = true
				} else {
					newBugMap[x][y] = false
				}
			}
		}
	}
	currentMap.bugMap = newBugMap
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	var bugMap [5][5]bool

	bioDiffs := make(map[int64]bool)
	scanner := bufio.NewScanner(file)

	y := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for x, elem := range line {
			if elem == "#" {
				bugMap[x][y] = true
			} else if elem == "." {
				bugMap[x][y] = false
			} else {
				fmt.Println("ERROR")
			}
		}
		y++
	}

	initMap := recBugMap{
		bugMap: bugMap,
		inner:  nil,
		outer:  nil,
	}
	printMap(bugMap)
	for true {

		newBioDiff := calcBioDiff(initMap.bugMap)
		updateBugMap(&initMap)
		printMap(initMap.bugMap)

		if _, ok := bioDiffs[newBioDiff]; ok {
			fmt.Println("Task 24.1: ", newBioDiff)
			break
		}
		bioDiffs[newBioDiff] = true
	}

}
