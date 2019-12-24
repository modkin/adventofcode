package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printMap(bugMap [5][5]bool) {
	fmt.Println("=============")
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if y == 2 && x == 2 {
				fmt.Print("?")

			} else if bugMap[x][y] {
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
	level  int
}

func updateBugMap(currentMap *recBugMap, currentMaxLvl int) {
	bugMap := currentMap.bugMap
	var newBugMap [5][5]bool
	hasOuter := true
	hasInner := true
	if currentMap.outer == nil {
		hasOuter = false
		newOuter := recBugMap{
			bugMap: [5][5]bool{},
			inner:  currentMap,
			outer:  nil,
			level:  currentMap.level + 1,
		}
		currentMap.outer = &newOuter
	}
	if currentMap.inner == nil {
		hasInner = false
		newInner := recBugMap{
			bugMap: [5][5]bool{},
			inner:  nil,
			outer:  currentMap,
			level:  currentMap.level - 1,
		}
		currentMap.inner = &newInner
	}
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if x == 2 && y == 2 {
				continue
			}
			bugCount := 0
			///WEST
			if x == 0 {
				if hasOuter {
					if currentMap.outer.bugMap[1][2] {
						bugCount++
					}
				}
			} else if x == 3 && y == 2 {
				if hasInner {
					for innerY := 0; innerY < 5; innerY++ {
						if currentMap.inner.bugMap[4][innerY] {
							bugCount++
						}
					}
				}
			} else if bugMap[x-1][y] {
				bugCount++
			}

			///EAST
			if x == 4 {
				if hasOuter {
					if currentMap.outer.bugMap[3][2] {
						bugCount++
					}
				}
			} else if x == 1 && y == 2 {
				if hasInner {
					for innerY := 0; innerY < 5; innerY++ {
						if currentMap.inner.bugMap[0][innerY] {
							bugCount++
						}
					}
				}
			} else if bugMap[x+1][y] {
				bugCount++
			}

			///NORTH
			if y == 0 {
				if hasOuter {
					if currentMap.outer.bugMap[2][1] {
						bugCount++
					}
				}
			} else if y == 3 && x == 2 {
				if hasInner {
					for innerX := 0; innerX < 5; innerX++ {
						if currentMap.inner.bugMap[innerX][4] {
							bugCount++
						}
					}
				}
			} else if bugMap[x][y-1] {
				bugCount++
			}

			///SOUTH
			if y == 4 {
				if hasOuter {
					if currentMap.outer.bugMap[2][3] {
						bugCount++
					}
				}
			} else if y == 1 && x == 2 {
				if hasInner {
					for innerX := 0; innerX < 5; innerX++ {
						if currentMap.inner.bugMap[innerX][0] {
							bugCount++
						}
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
	if currentMap.level <= currentMaxLvl && currentMap.level > 0 {
		updateBugMap(currentMap.outer, currentMaxLvl)
	} else if utils.IntAbs(currentMap.level) <= currentMaxLvl && currentMap.level < 0 {
		updateBugMap(currentMap.inner, currentMaxLvl)
	} else if currentMap.level == 0 {
		///Level 0
		updateBugMap(currentMap.outer, currentMaxLvl)
		updateBugMap(currentMap.inner, currentMaxLvl)
	}
	currentMap.bugMap = newBugMap
}

func countBugs(bugMap [5][5]bool) (number int) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if y == 2 && x == 2 {
				continue
			}
			if bugMap[x][y] {
				number++
			}
		}
	}
	return
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	var bugMap [5][5]bool

	//bioDiffs := make(map[int64]bool)
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

	printMap(initMap.bugMap)
	for i := 0; i < 200; i++ {
		fmt.Println(i)
		updateBugMap(&initMap, i)
	}

	totalNumber := countBugs(initMap.bugMap)
	innerMap := initMap.inner
	outerMap := initMap.outer

	printMap(initMap.bugMap)
	printMap(innerMap.bugMap)
	printMap(outerMap.bugMap)
	for true {
		if innerMap == nil && outerMap == nil {
			break
		}
		totalNumber += countBugs(innerMap.bugMap)
		totalNumber += countBugs(outerMap.bugMap)
		innerMap = innerMap.inner
		outerMap = outerMap.outer
	}

	fmt.Println("Task 24.2: ", totalNumber)
	//fmt.Println("Old")
	//for true {
	//
	//	newBioDiff := calcBioDiff(initMap.bugMap)
	//	updateBugMap(&initMap, 0)
	//	printMap(initMap.bugMap)
	//
	//	if _, ok := bioDiffs[newBioDiff]; ok {
	//		fmt.Println("Task 24.1: ", newBioDiff)
	//		break
	//	}
	//	bioDiffs[newBioDiff] = true
	//}

}
