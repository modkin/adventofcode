package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findPath(paths [][]string, connections map[string][]string) [][]string {
	nextPaths := make([][]string, 0)
	foundNext := false
	for _, currentPath := range paths {
		last := currentPath[len(currentPath)-1]
		if last == "end" {
			nextPaths = append(nextPaths, currentPath)
			continue
		}
	outer:
		for _, next := range connections[last] {
			if next == "start" {
				continue
			}
			currentPathCopy := utils.CopyStringSlice(currentPath)
			if strings.ToLower(next) == next {
				//if utils.CountStringinStringSlice(currentPath, next) == 1 {
				if !utils.SliceContains(currentPath, next) {
					currentPathCopy = append(currentPathCopy, next)
					nextPaths = append(nextPaths, utils.CopyStringSlice(currentPathCopy))
					foundNext = true
				} else {
					for key := range connections {
						if utils.CountStringinStringSlice(currentPath, strings.ToLower(key)) == 2 {
							continue outer
						}
					}

					currentPathCopy = append(currentPathCopy, next)
					nextPaths = append(nextPaths, utils.CopyStringSlice(currentPathCopy))
					foundNext = true
				}
				//}
			} else {
				currentPathCopy = append(currentPathCopy, next)
				nextPaths = append(nextPaths, utils.CopyStringSlice(currentPathCopy))
				foundNext = true
			}
		}
	}
	if foundNext {
		nextPaths = findPath(nextPaths, connections)
	}
	return nextPaths
}

func main() {
	file, err := os.Open("2021/day12/input")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}

	connections := make(map[string][]string)
	paths := make([][]string, 0)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "-")
		if _, ok := connections[line[0]]; !ok {
			connections[line[0]] = make([]string, 0)
		}
		if _, ok := connections[line[1]]; !ok {
			connections[line[1]] = make([]string, 0)
		}
		connections[line[0]] = append(connections[line[0]], line[1])
		connections[line[1]] = append(connections[line[1]], line[0])
	}

	paths = append(paths, []string{"start"})
	paths = findPath(paths, connections)
	fmt.Println(len(paths))
}
