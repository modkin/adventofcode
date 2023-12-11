package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func sum(left [2]int, right [2]int) [2]int {
	return [2]int{left[0] + right[0], left[1] + right[1]}
}

func main() {
	file, err := os.Open("2023/day10/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	pipeMap := make(map[[2]int]string)

	var start [2]int
	nextDirs := make(map[string][][2]int)
	nextDirs["-"] = [][2]int{{1, 0}, {-1, 0}}
	nextDirs["|"] = [][2]int{{0, 1}, {0, -1}}
	nextDirs["L"] = [][2]int{{1, 0}, {0, -1}}
	nextDirs["F"] = [][2]int{{1, 0}, {0, 1}}
	nextDirs["7"] = [][2]int{{-1, 0}, {0, 1}}
	nextDirs["J"] = [][2]int{{-1, 0}, {0, -1}}
	y := 0
	var maxX int
	var maxY int
	for scanner.Scan() {
		for x, pipe := range scanner.Text() {
			pipeMap[[2]int{x, y}] = string(pipe)
			if pipe == 'S' {
				start = [2]int{x, y}
			}
		}
		y++
		maxX = len(scanner.Text()) - 1
	}
	maxY = y - 1
	//utils.Print2DStringsGrid(pipeMap)
	fmt.Println("maxX", maxX, "maxY", maxY)

	//for ints, s := range pipeMap {
	//	for _, i := range nextDirs[s] {
	//		nbr := pipeMap[sum(ints, i)]
	//		if nbr == "." {
	//			panic(s)
	//		}
	//	}
	//}
	var startType string
	findStartType := func(startPos [2]int, startType string) bool {
		connections := 0
		for _, i := range nextDirs[startType] {
			nbrPos := sum(startPos, i)
			for _, j := range nextDirs[pipeMap[nbrPos]] {
				if startPos == sum(nbrPos, j) {
					connections++
				}
			}
		}
		return connections == 2
	}
	for _, s := range []string{"-", "|", "J", "F", "L", "7"} {
		if findStartType(start, s) {
			startType = s
		}
	}
	previousPos := start
	//choose first step random
	currentPos := sum(start, nextDirs[startType][0])
	numberOfSteps := 1

	path := make(map[[2]int]bool)
	path[currentPos] = true
	for currentPos != start {
		for _, i := range nextDirs[pipeMap[currentPos]] {
			if nextPos := sum(currentPos, i); nextPos != previousPos {
				previousPos = currentPos
				currentPos = nextPos
				break
			}
		}
		numberOfSteps++
		path[currentPos] = true

	}
	fmt.Println("Day 10.1:", numberOfSteps/2)

	newPipeMap := make(map[[2]int]string)
	newPipeMap[start] = startType
	for ints, pipeType := range pipeMap {
		if _, ok := path[ints]; ok {
			newPipeMap[ints] = pipeType
		} else {
			newPipeMap[ints] = "."
		}
	}
	newPipeMap[start] = startType
	pipeMap = newPipeMap

	checkInOut := func(start [2]int) {
		posList := [][2]int{start}
		currentPatch := make(map[[2]int]bool)
		currentPatch[start] = true
		isOuter := false
		isInnerWithoutWayOut := false
		for len(posList) != 0 {
			var newPosList [][2]int
			for _, pos := range posList {

				for _, nbr := range [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
					nbrPos := sum(pos, nbr)
					if nbrPos[0] < 0 || nbrPos[0] > maxX || nbrPos[1] < 0 || nbrPos[1] > maxY {
						isOuter = true
					} else if pipeMap[nbrPos] == "." {
						if _, ok := currentPatch[nbrPos]; !ok {
							newPosList = append(newPosList, nbrPos)
							currentPatch[nbrPos] = true
						}
					} else if _, ok := path[nbrPos]; ok {
						outWardString := pipeMap[nbrPos]
						tmpPos := nbrPos
						for {
							tmpPos = sum(tmpPos, nbr)
							if _, ok2 := path[tmpPos]; ok2 {
								outWardString += pipeMap[tmpPos]
							}
							if tmpPos[0] < 0 || tmpPos[0] > maxX || tmpPos[1] < 0 || tmpPos[1] > maxY {
								break
							}
						}
						counterOut := 0
						if nbr[0] != 0 {
							horizontal := strings.ReplaceAll(outWardString, "-", "")
							//beforeType := ""
							//for i, pipeType := range outWardString {
							//	if pipeType == '|' {
							//		counterOut++
							//	} else if pipeType == 'F' {
							//		if beforeType == "7" {
							//			counterOut += 2
							//		} else if beforeType == "J" {
							//			counterOut += 1
							//		}
							//		beforeType = "F"
							//	} else if pipeType == ''
							//}
							counterOut += strings.Count(horizontal, "|")
							if nbr[0] == 1 {
								counterOut += strings.Count(horizontal, "FJ")
								counterOut += strings.Count(horizontal, "L7")
							} else {
								counterOut += strings.Count(horizontal, "JF")
								counterOut += strings.Count(horizontal, "7L")
							}
						} else {
							vertical := strings.ReplaceAll(outWardString, "|", "")
							counterOut += strings.Count(vertical, "-")
							if nbr[1] == 1 {
								counterOut += strings.Count(vertical, "FJ")
								counterOut += strings.Count(vertical, "7L")
							} else {
								counterOut += strings.Count(vertical, "JF")
								counterOut += strings.Count(vertical, "L7")
							}

						}

						if counterOut%2 == 1 {
							isInnerWithoutWayOut = true
						}

					}
				}
			}
			posList = newPosList
		}

		if isOuter || !isInnerWithoutWayOut {
			for pos := range currentPatch {
				pipeMap[pos] = "O"
			}
		} else {
			for pos := range currentPatch {
				pipeMap[pos] = "I"
			}
		}
	}

	utils.Print2DStringsGrid(pipeMap)
	for x := 0; x < maxX; x++ {
		for y = 0; y < maxY; y++ {
			cur := pipeMap[[2]int{x, y}]
			if cur == "." {
				checkInOut([2]int{x, y})
			}
		}
	}
	utils.Print2DStringsGrid(pipeMap)
	counterInner := 0
	//for iPos, s := range pipeMap {
	//	if s == "I" {
	//
	//		unevenCounter := 0
	//		for _, nbr := range [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
	//			pos := iPos
	//			loopCounter := 0
	//			for {
	//
	//				pos = sum(pos, nbr)
	//				if pos[0] < 0 || pos[0] > maxX || pos[1] < 0 || pos[1] > maxY {
	//					if loopCounter%2 == 1 {
	//						unevenCounter++
	//					}
	//					break
	//				}
	//				if _, ok := path[pos]; ok {
	//					loopCounter++
	//				}
	//			}
	//		}
	//		if unevenCounter == 4 {
	//			fmt.Println(iPos)
	//		}
	//	}
	//}
	for _, s := range pipeMap {
		if s == "I" {
			counterInner++
		}
	}
	fmt.Println(counterInner)
}
