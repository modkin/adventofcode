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

func Solve() {
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
		depMap[dep] = true
		depMap[node] = true
	}
	for _, elem := range nodeMap {
		sort.Slice(elem, func(i, j int) bool {
			return i < j
		})
	}

	for key, _ := range depMap {
		depSlice = append(depSlice, key)
	}
	sort.Slice(depSlice, func(i, j int) bool {
		return depSlice[i] < depSlice[j]
	})

	for len(depSlice) != 0 {
		for depIdx, letter := range depSlice {
			if nodeMap[letter] == nil || len(nodeMap[letter]) == 0 {
				fmt.Print(string(letter))
				depSlice = append(depSlice[0:depIdx], depSlice[depIdx+1:]...)
				for keys, _ := range nodeMap {
					for idx, elem := range nodeMap[keys] {
						if elem == letter {
							nodeMap[keys] = append(nodeMap[keys][0:idx], nodeMap[keys][idx+1:]...)
							break
						}
					}
				}
				break
			}

		}
	}
}

func Task1() {
	Solve()
}
