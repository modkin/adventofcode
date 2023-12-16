package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
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

	getHash := func(in map[[2][2]int]bool) string {
		var keys []string
		for i, _ := range in {
			str := string(rune(i[0][0])) + string(rune(i[0][1])) + string(rune(i[1][0])) + string(rune(i[1][1]))
			keys = append(keys, str)
		}
		sort.Strings(keys)
		return strings.Join(keys, "-")
	}
	getEnergy := func(start [2][2]int) int {
		cache := make(map[string]bool)
		energy := make(map[[2]int]bool)
		beams := make(map[[2][2]int]bool)
		beams[start] = true
		for i := 0; i < 1000; i++ {
			for beam := range beams {
				energy[beam[0]] = true
			}
			beams = move(beams)
			hash := getHash(beams)
			if _, ok := cache[hash]; ok {
				break
			}
			cache[hash] = true
		}
		counter := 0
		for _, _ = range energy {
			counter++
		}
		return counter
	}
	fmt.Println(getEnergy([2][2]int{{0, 0}, {1, 0}}))

	maxEnergy := 0
	for i := 0; i <= maxX; i++ {
		fmt.Println(i)
		energy := getEnergy([2][2]int{{0, i}, {1, 0}})
		if energy > maxEnergy {
			maxEnergy = energy
		}
		energy = getEnergy([2][2]int{{maxX, i}, {-1, 0}})
		if energy > maxEnergy {
			maxEnergy = energy
		}
		energy = getEnergy([2][2]int{{i, 0}, {0, 1}})
		if energy > maxEnergy {
			maxEnergy = energy
		}
		energy = getEnergy([2][2]int{{i, maxY}, {0, -1}})
		if energy > maxEnergy {
			maxEnergy = energy
		}
	}
	fmt.Println(maxEnergy)
}
