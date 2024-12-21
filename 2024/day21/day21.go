package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func AddPoints(one, two [2]int) [2]int {
	return [2]int{one[0] + two[0], one[1] + two[1]}
}

func SubPoints(one, two [2]int) [2]int {
	return [2]int{one[0] - two[0], one[1] - two[1]}
}

var numCoordsMap = map[string][2]int{"7": [2]int{0, 0}, "8": [2]int{1, 0}, "9": [2]int{2, 0},
	"4": [2]int{0, 1}, "5": [2]int{1, 1}, "6": [2]int{2, 1},
	"1": [2]int{0, 2}, "2": [2]int{1, 2}, "3": [2]int{2, 2},
	"0": [2]int{1, 3}, "A": [2]int{2, 3}}

var dirCoordsMap = map[string][2]int{"^": {1, 0}, "A": {2, 0}, "<": {0, 1}, "v": {1, 1}, ">": {2, 1}}

func movesNumeric(start, end string) string {
	offset := SubPoints(numCoordsMap[end], numCoordsMap[start])
	var moves, xMoves, yMoves string
	for x := 0; x < utils.IntAbs(offset[0]); x++ {
		if offset[0] > 0 {
			xMoves += ">"
		} else {
			xMoves += "<"
		}
	}
	for y := 0; y < utils.IntAbs(offset[1]); y++ {
		if offset[1] > 0 {
			yMoves += "v"
		} else {
			yMoves += "^"
		}
	}
	if offset[0] > 0 {
		moves = xMoves + yMoves
	} else {
		moves = yMoves + xMoves
	}
	return moves
}

func movesDir(start, end string) string {
	offset := SubPoints(dirCoordsMap[end], dirCoordsMap[start])
	var moves, xMoves, yMoves string
	for x := 0; x < utils.IntAbs(offset[0]); x++ {
		if offset[0] > 0 {
			xMoves += ">"
		} else {
			xMoves += "<"
		}
	}
	for y := 0; y < utils.IntAbs(offset[1]); y++ {
		if offset[1] > 0 {
			yMoves += "v"
		} else {
			yMoves += "^"
		}
	}
	if offset[0] > 0 {
		moves = xMoves + yMoves
	} else {
		moves = yMoves + xMoves
	}
	return moves
}

func GenerateVariations(elements []string, start string, numeric bool) []string {
	var result [][]string
	var generate func([]string, int)
	generate = func(current []string, index int) {
		if index == len(current)-1 {
			temp := make([]string, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}
		for i := index; i < len(current); i++ {
			current[index], current[i] = current[i], current[index]
			generate(current, index+1)
			current[index], current[i] = current[i], current[index]
		}
	}
	generate(elements, 0)
	cleanedElement := make(map[string]bool)
	cleanedList := []string{}
	for _, elem := range result {
		safe := true
		var pos [2]int
		if numeric {
			pos = numCoordsMap[start]
		} else {
			pos = dirCoordsMap[start]
		}
		for _, e := range elem {
			if e == ">" {
				pos[0] += 1
			} else if e == "<" {
				pos[0] -= 1
			} else if e == "v" {
				pos[1] += 1
			} else if e == "^" {
				pos[1] -= 1
			}
			if numeric && pos[0] == 0 && pos[1] == 3 {
				safe = false
				break
			}
			if !numeric && pos[0] == 0 && pos[1] == 0 {
				safe = false
				break
			}
		}
		//safe = true
		if safe {
			tmp := append(elem, "A")
			cleanedElement[strings.Join(tmp, "")] = true
			cleanedList = append(cleanedList, strings.Join(tmp, ""))
		}
	}
	result = [][]string{}
	for k := range cleanedElement {
		result = append(result, strings.Split(k, ""))
	}
	return cleanedList
}

func getInputfromString(input string) string {
	pos := [2]int{2, 0}
	var ret string
	for _, s := range strings.Split(input, "") {
		if s == ">" {
			pos[0] += 1
		} else if s == "<" {
			pos[0] -= 1
		} else if s == "v" {
			pos[1] += 1
		} else if s == "^" {
			pos[1] -= 1
		} else if s == "A" {
			for i, i2 := range dirCoordsMap {
				if i2 == pos {
					ret += i
				}
			}
		}
	}
	return ret
}

func main() {

	//positions := []string{"A", "A", "A", "A"}
	codes := []string{}

	lines := utils.ReadFileIntoLines("2024/day21/input")
	for _, line := range lines {

		codes = append(codes, line)
	}

	sum := 0
	for i, code := range codes {
		//if i != 1 {
		//	continue
		//}
		fmt.Println(i)

		doorPos := "A"
		pos1 := "A"
		pos2 := "A"
		allDoorCombs := []string{}
		for _, s := range strings.Split(code, "") {
			basicDoorMoves := movesNumeric(doorPos, s)
			permDoor := GenerateVariations(strings.Split(basicDoorMoves, ""), doorPos, true)
			doorPos = s

			if len(allDoorCombs) == 0 {
				allDoorCombs = permDoor
			} else {
				newAllDoorCombs := []string{}
				for _, m := range permDoor {
					for _, comb := range allDoorCombs {
						newAllDoorCombs = append(newAllDoorCombs, comb+m)
					}
				}
				allDoorCombs = newAllDoorCombs
			}
		}

		allFirstCombs2 := []string{}
		for _, doorMoves := range allDoorCombs {
			allFirstCombs := []string{}
			for _, move1 := range strings.Split(doorMoves, "") {
				basicFirstMoves := movesDir(pos1, move1)
				//basicFirstMoves = append(basicFirstMoves, "A")
				var permFirst []string
				if len(basicFirstMoves) == 0 {
					permFirst = []string{"A"}
				} else {
					permFirst = GenerateVariations(strings.Split(basicFirstMoves, ""), pos1, false)
				}
				pos1 = move1

				if len(allFirstCombs) == 0 {
					allFirstCombs = permFirst
				} else {
					newallFirstCombs := []string{}
					for _, m := range permFirst {
						for _, comb := range allFirstCombs {
							newallFirstCombs = append(newallFirstCombs, comb+m)
						}
					}
					allFirstCombs = newallFirstCombs
				}
			}
			//fmt.Println(allFirstCombs)
			allFirstCombs2 = append(allFirstCombs2, allFirstCombs...)
		}

		minLen := math.MaxInt32
		for _, firstMoves := range allFirstCombs2 {
			l := 0
			var endMoves string
			for _, move2 := range strings.Split(firstMoves, "") {
				secondMoves := movesDir(pos2, move2)
				secondMoves += "A"
				pos2 = move2

				l += len(secondMoves)
				endMoves += secondMoves

			}

			//fmt.Println(firstMoves, l)
			if l < minLen {
				fmt.Println(firstMoves, l)
				fmt.Println(endMoves)
				minLen = l
			}
		}

		//moveStr := strings.Join(allMoves, "")

		numPart, _ := strconv.Atoi(code[:3])
		fmt.Println(numPart, minLen)
		sum += numPart * minLen

	}

	fmt.Println("Day 21.1: ", sum)

	fmt.Println(getInputfromString("<v<A>>^AAAvA^A<vA<AA>>^AvAA<^A>A<v<A>A>^AAAvA<^A>A<vA>^A<A>A"))
}
