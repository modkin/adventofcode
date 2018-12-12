package day7

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type node struct {
	name rune
	deps []rune
}

func processMap(curr rune, runeMap map[rune][]rune, printed map[rune]bool) {
	for _, elem := range runeMap[curr] {
		runeMap[curr] = runeMap[curr][1:]
		if printed[curr] == false {
			processMap(elem, runeMap, printed)
		}
	}
	if printed[curr] == false {
		fmt.Print(string(curr))
		printed[curr] = true
	}
}

func createNodeSlice() {
	file, err := os.Open("day7/day7-input.txt")
	if err != nil {
		panic(err)
	}

	nodeMap := make(map[rune][]rune)
	var depSlice []rune
	depMap := make(map[rune]bool)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		coords := strings.Split(word, " ")
		dep := rune([]rune(coords[1])[0])
		node := rune([]rune(coords[7])[0])
		nodeMap[node] = append(nodeMap[node], dep)
		//nodeMap[node] = append(nodeMap[node], dep)
		depMap[dep] = true
		//depSlice = append(depSlice, dep)
	}
	for _, elem := range nodeMap {
		sort.Slice(elem, func(i, j int) bool {
			return i < j
		})
	}

	for key, _ := range depMap {
		depSlice = append(depSlice, key)
	}
	fmt.Println(string(depSlice))
	sort.Slice(depSlice, func(i, j int) bool {
		return i < j
	})

	abs := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	depSlice = []rune(abs)

	fmt.Println(len(nodeMap))
	fmt.Println(string(depSlice))
	printMap := make(map[rune]bool)
	for {
		for key, elem := range nodeMap {
			for _, val := range depSlice {
				if val != key || elem == nil {
					if printMap[val] == false {
						fmt.Print(string(val))
						printMap[val] = true
					}
					for idx, entry := range depSlice {
						if val == entry {
							depSlice = append(depSlice[0:idx], depSlice[idx+1:]...)
							break
						}
					}
					delete(nodeMap, val)

					for _, elem2 := range nodeMap {
						for idx, entry := range elem2 {
							if val == entry {
								elem2 = append(elem2[0:idx], elem2[idx+1:]...)
								break
							}
						}
					}
					//processMap(key, nodeMap, make(map[rune]bool))
				}
				//fmt.Println("bla", string(val))
			}
		}
	}
}

func Task1() {
	createNodeSlice()
}
