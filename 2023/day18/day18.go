package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func print2DStringGridUpward(grid map[[2]int]bool) {
	fmt.Println("----------------------------------------------")
	xMax, yMax := 0, 0
	xMin, yMin := math.MaxInt, math.MaxInt
	for i := range grid {
		if i[0] > xMax {
			xMax = i[0]
		}
		if i[1] > yMax {
			yMax = i[1]
		}
		if i[0] < xMin {
			xMin = i[0]
		}
		if i[1] < yMin {
			yMin = i[1]
		}
	}
	for y := yMax; y >= yMin; y-- {
		for x := xMin; x <= xMax; x++ {
			if _, ok := grid[[2]int{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("----------------------------------------------")
}

func distance(first [2]int, second [2]int) int {
	ret := 0
	ret += utils.IntAbs(second[0] - first[0])
	ret += utils.IntAbs(second[1] - first[1])
	return ret
}

func sum(left [2]int, right [2]int) [2]int {
	return [2]int{left[0] + right[0], left[1] + right[1]}
}

func main() {
	file, err := os.Open("2023/day18/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	var points [][2]int
	pos := [2]int{0, 0}

	maxY2 := 0
	minX2 := 0
	for _, line := range lines {
		points = append(points, pos)
		hex := strings.Trim(strings.Fields(line)[2], "#()")
		hexDir := hex[5]
		var dir [2]int
		if hexDir == '0' {
			dir = [2]int{1, 0}
		} else if hexDir == '2' {
			dir = [2]int{-1, 0}
		} else if hexDir == '1' {
			dir = [2]int{0, 1}
		} else if hexDir == '3' {
			dir = [2]int{0, -1}
		}
		dist, _ := strconv.ParseInt(hex[0:5], 16, 32)
		fmt.Println(hex)
		fmt.Println(dist, dir)
		pos = [2]int{pos[0] + dir[0]*int(dist), pos[1] + dir[1]*int(dist)}
		if pos[1] > maxY2 {
			maxY2 = pos[1]
		}
		if pos[0] < minX2 {
			minX2 = pos[0]
		}
	}
	//fmt.Println(points)

	dig := make(map[[2]int]bool)
	maxY := 0
	maxX := 0
	minX := math.MaxInt
	minY := math.MaxInt
	pos = [2]int{0, 0}
	var points1 [][2]int
	for _, line := range lines {
		points1 = append(points1, pos)
		split := strings.Fields(line)
		var dir [2]int

		if split[0] == "R" {
			dir = [2]int{1, 0}
		} else if split[0] == "L" {
			dir = [2]int{-1, 0}
		} else if split[0] == "D" {
			dir = [2]int{0, 1}
		} else if split[0] == "U" {
			dir = [2]int{0, -1}
		}
		for i := 0; i < utils.ToInt(split[1]); i++ {
			dig[pos] = true
			pos = sum(pos, dir)
			if pos[0] > maxX {
				maxX = pos[0]
			}
			if pos[1] > maxY {
				maxY = pos[1]
			}
			if pos[0] < minX {
				minX = pos[0]
			}
			if pos[1] < minY {
				minY = pos[1]
			}
		}
	}
	//utils.Print2DStringGrid(dig)

	outside := make(map[[2]int]bool)
	allPos := make(map[[2]int]bool)
	allPos[[2]int{-1, 0}] = true
	for len(allPos) > 0 {
		newAllpos := make(map[[2]int]bool)

		for po, _ := range allPos {
			outside[po] = true

			var dirs [][2]int
			dirs = append(dirs, [2]int{1, 0})
			dirs = append(dirs, [2]int{-1, 0})
			dirs = append(dirs, [2]int{0, 1})
			dirs = append(dirs, [2]int{0, -1})

			for _, dir := range dirs {

				newPos := sum(po, dir)
				if newPos[0] < minX-1 || newPos[0] > maxX+1 || newPos[1] < minY-1 || newPos[1] > maxY+1 {
					continue
				}
				if _, ok := dig[newPos]; !ok {
					if _, ok2 := outside[newPos]; !ok2 {
						newAllpos[newPos] = true
					}
				}
			}

		}
		allPos = newAllpos

	}

	fmt.Println((maxX+utils.IntAbs(minX)+3)*(maxY+utils.IntAbs(minY)+3) - len(outside))

	//utils.Print2DStringGrid(dig)
	numPoints := len(points)
	//maxY2 = maxY
	//minX2 = minX

	points1Rev := [][2]int{{0, 0}}
	for i := numPoints - 1; i >= 0; i-- {
		points1Rev = append(points1Rev, points[i])
	}

	//points1Rev = [][2]int{{1, 6}, {3, 1}, {7, 2}, {4, 4}, {8, 5}, {1, 6}}
	var shiftedPoints [][2]int
	for _, ints := range points1Rev {
		newPoint := [2]int{utils.IntAbs(minX2) + ints[0], maxY2 - ints[1]}
		shiftedPoints = append(shiftedPoints, newPoint)
	}
	points1Rev = shiftedPoints
	fmt.Println(points1Rev)

	var outerPoints [][2]int
	prevDir := [2]int{points1Rev[0][0] - points1Rev[len(points1Rev)-2][0], points1Rev[0][1] - points1Rev[len(points1Rev)-2][1]}
	for i, currentPoint := range points1Rev {
		if i == len(points1Rev)-1 {
			break
		}
		newPoint := currentPoint
		dir := [2]int{points1Rev[i+1][0] - currentPoint[0], points1Rev[i+1][1] - currentPoint[1]}
		if dir[1] < 0 {
			if prevDir[0] < 0 {
				newPoint[1] += 1
			}
		}
		if dir[1] > 0 {
			if prevDir[0] > 0 {
				newPoint[0] += 1
			} else {
				newPoint[0] += 1
				newPoint[1] += 1
			}
		}
		if dir[0] > 0 {
			if prevDir[1] > 0 {
				newPoint[0] += 1
			}
		}
		if dir[0] < 0 {
			if prevDir[1] > 0 {
				newPoint[0] += 1
				newPoint[1] += 1
			} else {
				newPoint[1] += 1
			}

		}
		outerPoints = append(outerPoints, newPoint)
		prevDir = dir
	}
	outerPoints = append(outerPoints, outerPoints[0])
	fmt.Println(outerPoints)
	points1Rev = outerPoints
	tmpgrid := make(map[[2]int]bool)
	for _, ints := range points1Rev {
		tmpgrid[ints] = true
	}
	//print2DStringGridUpward(tmpgrid)

	area := 0
	rectArea := 0
	for i := 0; i < len(points1Rev)-1; i++ {
		pi := points1Rev[i]
		pi1 := points1Rev[i+1]
		if pi[0] == pi1[0] {
			//if rectArea < 0 {
			//	area += utils.IntAbs(pi1[1] - pi[1])
			//	fmt.Println("adjust Area", area)
			//}
			continue
		}
		//tmp := float64(pi[0]*pi1[1] - pi1[0]*pi[1])

		rectArea = (distance(pi, pi1)) * (pi[1]) * ((pi[0] - pi1[0]) / utils.IntAbs(pi[0]-pi1[0]))
		//rectArea = pi[0]*pi1[1] - pi1[0]*pi[1]
		//fmt.Println("rect area", rectArea)
		area += rectArea
		//if rectArea < 0 {
		//	area += utils.IntAbs(pi1[0] - pi[0])
		//}

		//fmt.Println("Area", area)
		//fmt.Println()

	}

	fmt.Println(area)

}
