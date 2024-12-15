package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func add2Int(one, two [2]int) [2]int {
	return [2]int{one[0] + two[0], one[1] + two[1]}
}

func move(wh map[[2]int]string, pos [2]int, dirStr string) [2]int {
	dirMap := map[string][2]int{"<": {-1, 0}, ">": {1, 0}, "v": {0, 1}, "^": {0, -1}}

	dir := dirMap[dirStr]

	var boxes [][2]int
	newPos := add2Int(pos, dir)
	for {
		if wh[newPos] == "O" {
			boxes = append(boxes, newPos)
		} else if wh[newPos] == "." {
			break
		} else if wh[newPos] == "#" {
			boxes = [][2]int{}
			break
		}
		newPos = add2Int(newPos, dir)
	}
	if len(boxes) != 0 {
		for i := len(boxes) - 1; i >= 0; i-- {
			wh[add2Int(boxes[i], dir)] = "O"
		}
		wh[add2Int(pos, dir)] = "."
	}
	newRoboPos := add2Int(pos, dir)
	if wh[newRoboPos] == "." {
		return newRoboPos
	} else {
		return pos
	}
}

func move2(wh map[[2]int]string, pos [2]int, dirStr string) [2]int {
	dirMap := map[string][2]int{"<": {-1, 0}, ">": {1, 0}, "v": {0, 1}, "^": {0, -1}}

	dir := dirMap[dirStr]

	boxes := make(map[[2]int]string)
	if dirStr == "<" || dirStr == ">" {

		newPos := add2Int(pos, dir)
		for {
			if wh[newPos] == "[" || wh[newPos] == "]" {
				boxes[newPos] = wh[newPos]
				newPos = add2Int(newPos, dir)
				boxes[newPos] = wh[newPos]
			} else if wh[newPos] == "." {
				break
			} else if wh[newPos] == "#" {
				boxes = make(map[[2]int]string)
				break
			}
			newPos = add2Int(newPos, dir)
		}
		if len(boxes) != 0 {
			for pb, s := range boxes {
				wh[add2Int(pb, dir)] = s
			}
			wh[add2Int(pos, dir)] = "."
		}

	} else {
		var pushFront [][2]int
		var newBoxes [][][2]int

		pushFront = append(pushFront, pos)
	outer:
		for {
			var newPushFront [][2]int
			for _, ints := range pushFront {
				newPos := add2Int(ints, dir)
				if wh[newPos] == "#" {
					newBoxes = [][][2]int{}
					break outer
				}
				if wh[newPos] == "[" {
					newPushFront = append(newPushFront, newPos)
					nbr := add2Int(newPos, [2]int{1, 0})
					newPushFront = append(newPushFront, nbr)

					if wh[nbr] != "]" {
						fmt.Println("ERROR")
					}
				}
				if wh[newPos] == "]" {
					newPushFront = append(newPushFront, newPos)
					nbr := add2Int(newPos, [2]int{-1, 0})
					newPushFront = append(newPushFront, nbr)

					if wh[nbr] != "[" {
						fmt.Println("ERROR")
					}
				}
			}
			if len(newPushFront) == 0 {
				break outer
			} else {

				pushFront = newPushFront
				newBoxes = append(newBoxes, newPushFront)
			}
		}

		if len(newBoxes) != 0 {
			newWh := make(map[[2]int]string)
			for i := len(newBoxes) - 1; i >= 0; i-- {
				for _, ints := range newBoxes[i] {
					newWh[add2Int(ints, dir)] = wh[ints]
					newWh[ints] = "."
				}
			}
			fmt.Println(newBoxes)
			utils.Print2DStringsGrid(newWh)
			for ints, s := range newWh {
				wh[ints] = s
			}
		}

	}

	newRoboPos := add2Int(pos, dir)
	if wh[newRoboPos] == "." {
		return newRoboPos
	} else {
		return pos
	}
}

func main() {
	lines := utils.ReadFileIntoLines("2024/day15/input")

	wh := make(map[[2]int]string)
	whWorking := true

	//xMax := 0
	yMax := 0

	var moves []string
	var roboPos [2]int
	for _, line := range lines {
		if whWorking {
			if line == "" {
				whWorking = false
			} else {
				//xMax = len(line)
				split := strings.Split(line, "")
				for i2, s := range split {
					if s == "@" {
						roboPos = [2]int{i2, yMax}
						wh[[2]int{i2, yMax}] = "."
					} else {
						wh[[2]int{i2, yMax}] = s
					}
				}
				yMax++
			}
		} else {
			split := strings.Split(line, "")
			for _, i2 := range split {
				moves = append(moves, i2)
			}
		}

	}

	//for _, s := range moves {
	//	roboPos = move(wh, roboPos, s)
	//}

	sum := 0
	for ints, s := range wh {
		if s == "O" {
			sum += 100*ints[1] + ints[0]
		}
	}
	//utils.Print2DStringsGrid(wh)
	fmt.Println(roboPos)
	fmt.Println("Day 15.1:", sum)

	wwh := make(map[[2]int]string)
	wideMap := map[string][2]string{"#": {"#", "#"}, "O": {"[", "]"}, ".": {".", "."}}
	for ints, s := range wh {
		wwh[[2]int{ints[0] * 2, ints[1]}] = wideMap[s][0]
		wwh[[2]int{ints[0]*2 + 1, ints[1]}] = wideMap[s][1]
	}
	roboPos2 := [2]int{roboPos[0] * 2, roboPos[1]}

	utils.Print2DStringsGrid(wwh)
	for _, s := range moves {
		roboPos2 = move2(wwh, roboPos2, s)
		wwh[roboPos2] = "@"
		//utils.Print2DStringsGrid(wwh)
		wwh[roboPos2] = "."
	}

	sum2 := 0
	for ints, s := range wwh {
		if s == "[" {
			sum2 += 100*ints[1] + ints[0]
		}
	}
	fmt.Println(sum2)

}
