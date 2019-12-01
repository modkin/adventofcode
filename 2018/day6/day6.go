package day6

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type coordiante struct {
	x int
	y int
}

func createCoordinateList() []coordiante {
	file, err := os.Open("day6/day6-input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var coordList []coordiante

	for scanner.Scan() {
		word := scanner.Text()
		coords := strings.Split(word, ",")
		c := coordiante{utils.ToInt(coords[0]), utils.ToInt(coords[1])}
		coordList = append(coordList, c)
	}
	return coordList
}

func getMax(coords []coordiante) (xmax, ymax int) {
	for _, elem := range coords {
		if xmax < elem.x {
			xmax = elem.x
		}
		if ymax < elem.y {
			ymax = elem.y
		}
	}
	return
}

func shortestDistance(dists []coordiante, x int, y int) (coord int) {
	minDist := 10000000000
	for idx, elem := range dists {
		dist := utils.IntAbs(elem.x-x) + utils.IntAbs(elem.y-y)
		if dist < minDist {
			minDist = dist
			coord = idx + 1
		} else if dist == minDist {
			coord = 0
		}
	}
	return
}

func fillDistances(dists [][]int, coords []coordiante) {
	for x := 0; x < len(dists); x++ {
		for y := 0; y < len(dists[0]); y++ {
			dists[x][y] = shortestDistance(coords, x, y)
		}
	}
}

func getInfList(dists [][]int) map[int]bool {
	infs := make(map[int]bool)
	for x := 0; x < len(dists); x++ {
		infs[dists[x][0]] = true
		infs[dists[x][len(dists[0])-1]] = true
	}
	for y := 0; y < len(dists[0]); y++ {
		infs[dists[0][y]] = true
		infs[dists[len(dists)-1][y]] = true
	}
	return infs
}

func isNotInMap(infinity map[int]bool, val int) bool {
	for key, _ := range infinity {
		if key == val {
			return false
		}
	}
	return true
}

func findMaxSize(dists [][]int, infinity map[int]bool) int {
	sum := make([]int, len(dists)*len(dists[0]))
	for x := 0; x < len(dists); x++ {
		for y := 0; y < len(dists[0]); y++ {
			if isNotInMap(infinity, dists[x][y]) {
				sum[dists[x][y]]++
			}
		}
	}
	max := 0
	for _, elem := range sum {
		if elem > max {
			max = elem
		}
	}
	return max
}

func closeToAll(coords []coordiante, x int, y int, distance int) bool {
	for _, elem := range coords {
		dist := utils.IntAbs(elem.x-x) + utils.IntAbs(elem.y-y)
		if dist > distance {
			return false
		}
	}
	return true
}

func getAreaWithAllPoints(coords []coordiante, xmax int, ymax int, distance int) int {
	area := 0
	for x := 0; x < xmax; x++ {
		for y := 0; y < ymax; y++ {
			counter := 0
			for _, elem := range coords {
				dist := utils.IntAbs(elem.x-x) + utils.IntAbs(elem.y-y)
				counter += dist
			}
			if counter < distance {
				area++
			}
		}
	}
	return area
}

func Task1() {
	coordList := createCoordinateList()
	xmax, ymax := getMax(coordList)
	distances := make([][]int, xmax+1)
	for i := range distances {
		distances[i] = make([]int, ymax+1)
	}
	fillDistances(distances, coordList)

	//for j := 0; j < len(distances[0]); j++ {
	//	for i := 0; i < len(distances); i++ {
	//		fmt.Print(distances[i][j])
	//	}
	//	fmt.Println()
	//}
	infin := getInfList(distances)
	max := findMaxSize(distances, infin)
	fmt.Println(max)
	area := getAreaWithAllPoints(coordList, xmax+1, ymax+1, 10000)
	fmt.Println(area)

}
