package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"slices"
	"sort"
)

type hikeStruct struct {
	pos     [2]int
	visited [][2]int
	length  int
	id      int
}

var dirs = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func sum(left [2]int, right [2]int) [2]int {
	return [2]int{left[0] + right[0], left[1] + right[1]}
}

type byYX [][2]int

func (c byYX) Len() int           { return len(c) }
func (c byYX) Less(i, j int) bool { return c[i][1] < c[j][1] && c[i][0] < c[j][0] }
func (c byYX) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func filterSame(hikes []hikeStruct) {
	ids := make(map[int][]int)
	for _, hike := range hikes {
		sort.Sort(byYX(hike.visited))
	}
	for _, hike := range hikes {
		for _, hike2 := range hikes {
			if hike.id == hike2.id {
				continue
			}
			if reflect.DeepEqual(hike.visited, hike2.visited) {
				ids[hike.id] = append(ids[hike.id], hike2.id)
			}
		}
	}
	fmt.Println(ids)
}

func part1(forestMap map[[2]int]string, start [2]int, maxY int) {
	allHikeCounter := 0

	newHike := hikeStruct{
		pos:     start,
		length:  0,
		id:      allHikeCounter,
		visited: [][2]int{start},
	}
	var allHikes []hikeStruct
	allHikes = append(allHikes, newHike)

	allHikeCounter++
	hikeLenght := make(map[int]int)
	for len(allHikes) != 0 {
		var newAllHikes []hikeStruct
		for _, hike := range allHikes {
			dirCounter := 0
			for _, dir := range dirs {
				newPos := sum(hike.pos, dir)
				if newPos[1] < 0 {
					continue
				}
				if slices.Contains(hike.visited, newPos) {
					continue
				}
				if forestMap[newPos] == "#" {
					continue
				}
				if forestMap[newPos] == ">" && dir != [2]int{1, 0} {
					continue
				}
				if forestMap[newPos] == "<" && dir != [2]int{-1, 0} {
					continue
				}
				if forestMap[newPos] == "v" && dir != [2]int{0, 1} {
					continue
				}
				if forestMap[newPos] == "^" && dir != [2]int{0, -1} {
					continue
				}
				if newPos[1] == maxY {
					hikeLenght[hike.id] = hike.length + 1
				} else {
					newHikeId := hike.id
					if dirCounter != 0 {
						newHikeId = allHikeCounter
						allHikeCounter++
					}
					newHike = hikeStruct{
						visited: append(slices.Clone(hike.visited), newPos),
						pos:     newPos,
						length:  hike.length + 1,
						id:      newHikeId,
					}
					newAllHikes = append(newAllHikes, newHike)
					dirCounter++
				}
			}
		}
		allHikes = newAllHikes
	}
	//fmt.Println(hikeLenght)
	maxLen := 0
	for _, i2 := range hikeLenght {
		if i2 > maxLen {
			maxLen = i2
		}
	}
	fmt.Println("Day 23.1:", maxLen)
}

func main() {
	file, err := os.Open("2023/day23/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	forestMap := make(map[[2]int]string)
	y := 0
	var maxX int
	var maxY int
	var start [2]int
	var end [2]int
	for scanner.Scan() {
		for x, cost := range scanner.Text() {
			forestMap[[2]int{x, y}] = string(cost)
			if y == 0 && cost == '.' {
				start = [2]int{x, y}
			}
			if cost == '.' {
				end = [2]int{x, y}
			}
		}
		y++
		maxX = len(scanner.Text()) - 1

	}
	maxY = y - 1
	fmt.Println("maxX:", maxX, "maxY:", maxY)
	fmt.Println("start", start)
	fmt.Println("end:", end)

	part1(forestMap, start, maxY)

	var allCrossRoads [][2]int
	allCrossRoads = append(allCrossRoads, start)
	allCrossRoads = append(allCrossRoads, end)
	for ints, s := range forestMap {
		if !(s == "#" || s == ".") {
			forestMap[ints] = "."
		}
	}
	//utils.Print2DStringsGrid(forestMap)

	for pos, s := range forestMap {
		if s != "." {
			continue
		}
		dotCount := 0
		for _, dir := range dirs {
			tmpPos := sum(pos, dir)
			if forestMap[tmpPos] == "." {
				dotCount++
			}
		}
		if dotCount >= 3 {
			allCrossRoads = append(allCrossRoads, pos)
		}
	}
	simpleDistance := func(start [2]int, prev [2]int) (end [2]int, dist int) {
		dist = 1
		for {
			for _, dir := range dirs {
				nextPos := sum(start, dir)
				if nextPos == prev {
					continue
				}
				if slices.Contains(allCrossRoads, nextPos) {
					dist++
					end = nextPos
					return
				}
				if forestMap[nextPos] == "." {
					dist++
					prev = start
					start = nextPos
				}
			}
		}
	}

	//fmt.Println(allCrossRoads)

	crossRoadDistances := make(map[[2]int]map[[2]int]int)
	for _, cr := range allCrossRoads {
		crossRoadDistances[cr] = make(map[[2]int]int)
		for _, dir := range dirs {
			nextPos := sum(cr, dir)
			if forestMap[nextPos] == "." {
				endPos, dist := simpleDistance(nextPos, cr)
				crossRoadDistances[cr][endPos] = dist
			}
		}
	}
	//fmt.Println(crossRoadDistances)

	getPathLengh := func(path [][2]int) int {
		dist := 0
		for i, p := range path {
			if i == len(path)-1 {
				continue
			}
			dist += crossRoadDistances[p][path[i+1]]
		}
		return dist
	}

	fmt.Println(len(allCrossRoads))
	for i, road := range crossRoadDistances {
		fmt.Println(i)
		fmt.Println(road)
		fmt.Println()
	}
	var allDists []int
	var allPaths [][][2]int
	allPaths = append(allPaths, [][2]int{start})

	for len(allPaths) > 0 {
		var newAllPaths [][][2]int
		for _, path := range allPaths {
			for pathEnd, _ := range crossRoadDistances[path[len(path)-1]] {
				if slices.Contains(path, pathEnd) {
					continue
				}
				newPath := append(slices.Clone(path), pathEnd)
				if pathEnd == end {
					length := getPathLengh(newPath)
					allDists = append(allDists, length)
				} else {
					newAllPaths = append(newAllPaths, newPath)
				}
			}
		}
		allPaths = newAllPaths

	}

	fmt.Println("Day 23.2:", slices.Max(allDists))
}
