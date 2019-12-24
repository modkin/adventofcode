package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printMap(bugMap [7][7]bool) {
	fmt.Println("=============")
	for y := 0; y < 7; y++ {
		for x := 0; x < 7; x++ {
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

func calcBioDiff(bugMap [7][7]bool) (bioDiff int64) {
	bioPoints := int64(1)
	for y := 1; y < 6; y++ {
		for x := 1; x < 6; x++ {
			if bugMap[x][y] {
				bioDiff += bioPoints
			}
			bioPoints *= 2
		}
	}
	return
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	var bugMap [7][7]bool
	var newBugMap [7][7]bool

	bioDiffs := make(map[int64]bool)
	scanner := bufio.NewScanner(file)

	y := 1
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for x, elem := range line {
			if elem == "#" {
				bugMap[x+1][y] = true
			} else if elem == "." {
				bugMap[x+1][y] = false
			} else {
				fmt.Println("ERROR")
			}
		}
		y++
	}

	printMap(bugMap)
	for true {
		for y := 1; y < 6; y++ {
			for x := 1; x < 6; x++ {
				bugCount := 0
				if bugMap[x+1][y] {
					bugCount++
				}
				if bugMap[x-1][y] {
					bugCount++
				}
				if bugMap[x][y+1] {
					bugCount++
				}
				if bugMap[x][y-1] {
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
		bugMap = newBugMap
		newBioDiff := calcBioDiff(bugMap)
		if _, ok := bioDiffs[newBioDiff]; ok {
			fmt.Println("Task 24.1: ", newBioDiff)
			break
		}
		bioDiffs[newBioDiff] = true
	}

}
