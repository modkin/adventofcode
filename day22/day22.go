package day22

import (
	"fmt"
	"math"
)

type field struct {
	distance int
	//torch 1; climbing gear 2
	tool    int
	visited bool
}

func typeMap(depth int, x int, y int) [][]int {
	typeMap := make([][]int, x+1)
	for i := range typeMap {
		typeMap[i] = make([]int, y+1)
	}
	var geoIndex int
	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			if i == 0 && j == 0 {
				geoIndex = 0
			} else if i == x && j == y {
				geoIndex = 0
			} else if j == 0 {
				geoIndex = i * 16807
			} else if i == 0 {
				geoIndex = j * 48271
			} else {
				geoIndex = typeMap[i-1][j] * typeMap[i][j-1]
			}
			erosionLevel := (geoIndex + depth) % 20183
			typeMap[i][j] = erosionLevel
		}
	}
	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			typeMap[i][j] = typeMap[i][j] % 3
		}
	}
	return typeMap
}

func printTypeMap(typmap [][]int) {
	for y := 0; y < len(typmap[0]); y++ {
		for x := 0; x < len(typmap); x++ {
			switch typmap[x][y] {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("=")
			case 2:
				fmt.Print("|")
			}
		}
		fmt.Println()
	}
}

func sumTypeMap(typemap [][]int) (ret int) {
	for y := 0; y < len(typemap[0]); y++ {
		for x := 0; x < len(typemap); x++ {
			ret += typemap[x][y]
		}
	}
	return
}

func accessible(tool int, ftype int) bool {
	switch tool {
	case 0:
		if ftype == 0 {
			return false
		}
	case 1:
		if ftype == 1 {
			return false
		}
	case 2:
		if ftype == 2 {
			return false
		}
	}
	return true
}

func findTool(sourceRegion int, targetRegion int) int {
	switch sourceRegion {
	case 0:
		if targetRegion == 1 {
			return 2
		} else {
			return 1
		}
	case 1:
		if targetRegion == 0 {
			return 2
		} else {
			return 0
		}
	case 2:
		if targetRegion == 0 {
			return 1
		} else {
			return 0
		}
	}
	return math.MaxInt64
}

func findMin(distance [][]field) []int {
	min := []int{0, 0}
	minDist := math.MaxInt64
	for idx, x := range distance {
		for idy, elem := range x {
			if elem.distance < minDist && !elem.visited {
				minDist = elem.distance
				min = []int{idx, idy}
			}
		}
	}
	return min
}

func findShortestPath(depth int, x int, y int) int {
	typemap := typeMap(4080, x*2, y*2)

	distance := make([][]field, x*2)
	for i := range distance {
		distance[i] = make([]field, y*2)
		for idx, _ := range distance[i] {
			distance[i][idx].distance = math.MaxInt64
			distance[i][idx].tool = 13
		}
	}
	distance[0][0].tool = 1
	distance[0][0].distance = 0

	finalTarget := []int{x, y}
	min := []int{0, 0}

	offsets := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -0}}

	totalDistance := 0
	for min[0] != finalTarget[0] && min[1] != finalTarget[1] {
		min = findMin(distance)
		distance[min[0]][min[1]].visited = true
		stepcost := 1
		source := distance[min[0]][min[1]]
		for _, offset := range offsets {
			if min[0]+offset[0] < 0 || min[1]+offset[1] < 0 {
				continue
			}
			tx := min[0] + offset[0]
			ty := min[1] + offset[1]
			tool := source.tool
			sourceRegion := typemap[min[0]][min[1]]
			targetRegion := typemap[min[0]+offset[0]][min[1]+offset[1]]
			if !accessible(tool, targetRegion) {
				stepcost += 7
			}
			if source.distance+stepcost < distance[tx][ty].distance {
				distance[tx][ty].distance = source.distance + stepcost
				totalDistance = distance[tx][ty].distance
				if sourceRegion == targetRegion {
					distance[tx][ty].tool = source.tool
				} else {
					distance[tx][ty].tool = findTool(source.tool, distance[tx][ty].tool)
				}
			}
		}
	}
	return totalDistance
}

func Task1() {
	typemap := typeMap(4080, 14, 785)
	//printTypeMap(typemap)
	fmt.Println(sumTypeMap(typemap))
}

func Task2() {
	fmt.Println(findShortestPath(510, 10, 10))
}
