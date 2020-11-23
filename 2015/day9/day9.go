package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2015/day9/input.txt")
	if err != nil {
		panic(err)
	}
	distances := make(map[string]map[string]int)

	allCities := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dist := strings.Split(scanner.Text(), " ")
		if _, ok := distances[dist[0]]; !ok {
			distances[dist[0]] = make(map[string]int)
		}
		if _, ok := distances[dist[2]]; !ok {
			distances[dist[2]] = make(map[string]int)
		}
		distances[dist[0]][dist[2]] = utils.ToInt(dist[4])
		distances[dist[2]][dist[0]] = utils.ToInt(dist[4])
		if _, ok := allCities[dist[0]]; !ok {
			allCities[dist[0]] = 0
		}
		if _, ok := allCities[dist[2]]; !ok {
			allCities[dist[2]] = 0
		}
	}
	cityCount := len(allCities)

	allRoutes := make([][]string, 0)
	var nextCity func(current string, visited []string, distance int) int
	nextCity = func(current string, visited []string, distance int) int {
		minLen := math.MaxInt32
		for dst := range allCities {
			if utils.SliceContains(visited, dst) {
				continue
			}
			distance += distances[current][dst]
			newVisited := utils.CopyStringSlice(visited)
			newVisited = append(newVisited, dst)
			if len(newVisited) == cityCount {
				allRoutes = append(allRoutes, newVisited)
				return distance

			} else {
				if newDst := nextCity(dst, newVisited, distance); newDst < minLen {
					minLen = newDst
				}
			}
		}
		//fmt.Println(minLen)
		return minLen
	}
	minimalLength := math.MaxInt32

	for start := range allCities {
		if newDst := nextCity(start, []string{start}, 0); newDst < minimalLength {
			minimalLength = newDst
		}
		fmt.Println(len(allRoutes))
	}
	minimalLength = math.MaxInt32
	for _, route := range allRoutes {
		lenght := 0
		for i := 0; i < len(route)-1; i++ {
			lenght += distances[route[i]][route[i+1]]
		}
		if lenght < minimalLength {
			minimalLength = lenght
		}
	}

	fmt.Println("Task 9.1: ", minimalLength)
}
