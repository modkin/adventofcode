package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func sum(left [2]int, right [2]int) [2]int {
	return [2]int{left[0] + right[0], left[1] + right[1]}
}

func main() {
	file, err := os.Open("2023/day18/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	dig := make(map[[2]int]bool)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	maxX := 0
	maxY := 0
	minX := math.MaxInt
	minY := math.MaxInt
	pos := [2]int{0, 0}
	for _, line := range lines {
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
	utils.Print2DStringGrid(dig)

	//lagoon := make(map[[2]int]bool)
	//counter := 0
	//for y := 0; y <= maxY; y++ {
	//	inside := false
	//	//wasSpace := false
	//	for x := 0; x <= maxX; x++ {
	//
	//		pos = [2]int{x, y}
	//		if _, ok := dig[pos]; ok {
	//			counter++
	//			lagoon[pos] = true
	//
	//			if _, ok2 := dig[sum(pos, [2]int{1, 0})]; !ok2 {
	//				if _, ok3 := dig[sum(pos, [2]int{-1, 0})]; !ok3 {
	//					if inside {
	//						inside = false
	//					} else {
	//						inside = true
	//					}
	//				}
	//			}
	//		} else {
	//			if inside {
	//				counter++
	//				lagoon[pos] = true
	//			}
	//		}
	//	}
	//}
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

	counter := 0
	for _, _ = range dig {
		counter++
	}

	fmt.Println(counter)
	fmt.Println(len(dig))
	utils.Print2DStringGrid(outside)
	utils.Print2DStringGrid(dig)

	fmt.Println((maxX+utils.IntAbs(minX)+3)*(maxY+utils.IntAbs(minY)+3) - len(outside))

}
