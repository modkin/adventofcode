package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("2023/day16/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	//cache := make(map[string]string)
	ma := make(map[[2]int]string)
	y := 0
	var maxX int
	var maxY int
	for scanner.Scan() {
		for x, pipe := range scanner.Text() {
			ma[[2]int{x, y}] = string(pipe)
		}
		y++
		maxX = len(scanner.Text()) - 1
	}
	maxY = y - 1
	fmt.Println(maxX, maxY)

	utils.Print2DStringsGrid(ma)
	energy := make(map[[2]int]bool)
	beams := make(map[[2][2]int]bool)
	beams[[2][2]int{[2]int{0, 0}, [2]int{1, 0}}] = true
	move := func(in map[[2][2]int]bool) map[[2][2]int]bool {
		//var out [][2][2]int
		out := make(map[[2][2]int]bool)
		for ints, bo := range in {
			if bo == false {
				continue
			}
			var nextPos [2][2]int
			//var nextPosList [][2][2]int
			oldPos := ints[0]
			oldDir := ints[1]
			posType := ma[ints[0]]
			if posType == "/" {
				if oldDir[0] == 1 {
					nextPos = [2][2]int{{oldPos[0], oldPos[1] - 1}, {0, -1}}
				} else if oldDir[0] == -1 {
					nextPos = [2][2]int{{oldPos[0], oldPos[1] + 1}, {0, 1}}
				} else if oldDir[1] == 1 {
					nextPos = [2][2]int{{oldPos[0] - 1, oldPos[1]}, {-1, 0}}
				} else if oldDir[1] == -1 {
					nextPos = [2][2]int{{oldPos[0] + 1, oldPos[1]}, {1, 0}}
				}
				out[nextPos] = true
			} else if posType == "\\" {
				if oldDir[0] == 1 {
					nextPos = [2][2]int{{oldPos[0], oldPos[1] + 1}, {0, 1}}
				} else if oldDir[0] == -1 {
					nextPos = [2][2]int{{oldPos[0], oldPos[1] - 1}, {0, -1}}
				} else if oldDir[1] == 1 {
					nextPos = [2][2]int{{oldPos[0] + 1, oldPos[1]}, {1, 0}}
				} else if oldDir[1] == -1 {
					nextPos = [2][2]int{{oldPos[0] - 1, oldPos[1]}, {-1, 0}}
				}
				out[nextPos] = true
			} else if posType == "-" {
				out[[2][2]int{{oldPos[0] + 1, oldPos[1]}, {1, 0}}] = true
				out[[2][2]int{{oldPos[0] - 1, oldPos[1]}, {-1, 0}}] = true
			} else if posType == "|" {
				out[[2][2]int{{oldPos[0], oldPos[1] + 1}, {0, 1}}] = true
				out[[2][2]int{{oldPos[0], oldPos[1] - 1}, {0, -1}}] = true
			} else {
				//if oldPos[0]+oldDir[0] <= maxX && oldPos[1]+oldDir[1] <= maxY && oldPos[0]+oldDir[0] >= 0 && oldPos[1]+oldDir[1] >= 0 {
				out[[2][2]int{{oldPos[0] + oldDir[0], oldPos[1] + oldDir[1]}, oldDir}] = true
				//}
			}

		}
		cutOf := make(map[[2][2]int]bool)
		for pos, _ := range out {
			oldPos := pos[0]
			if oldPos[0] > maxX || oldPos[1] > maxY || oldPos[0] < 0 || oldPos[1] < 0 {
				out[pos] = false
			} else {
				cutOf[pos] = true
			}
		}
		out = cutOf
		return out
	}

	for len(beams) != 0 {
		utils.Print2DStringGrid(energy)
		changed := false
		for beam := range beams {
			energy[beam[0]] = true
			changed = true
		}
		counter := 0
		for _, _ = range energy {
			counter++
		}
		fmt.Println("count:", counter)
		if changed {
			beams = move(beams)

		} else {
			break
		}
	}
	counter := 0
	for _, _ = range energy {
		counter++
	}
	fmt.Println(counter)
}
