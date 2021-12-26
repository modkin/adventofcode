package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var hallwayPos = [7]int{0, 1, 3, 5, 7, 9, 10}

const LINES int = 4

type scene struct {
	hallway [11]string
	rooms   [4][LINES]string
	//cost    int
}

func printScene(input scene) {
	fmt.Println("#############")
	fmt.Print("#")
	for _, h := range input.hallway {
		if h == "" {
			fmt.Print(".")
		} else {
			fmt.Print(h)
		}
	}
	fmt.Println("#")
	for i := LINES - 1; i >= 0; i-- {
		fmt.Print("###")
		for r := 0; r < 4; r++ {
			if input.rooms[r][i] == "" {
				fmt.Print(".")
			} else {
				fmt.Print(input.rooms[r][i])
			}
			fmt.Print("#")
		}
		fmt.Println("##")
	}
	fmt.Println("#############")
}

func checkScenes(allScenes map[scene]int) (int, scene, bool) {
	min := math.MaxInt
	var minScene scene
	finished := [4]bool{}
	for i, s := range allScenes {
		if s < min {
			min = s
			minScene = i
		}
	}
	for i := 0; i < 4; i++ {
		tmp, _ := getOwnRoom(minScene.rooms[i][0])
		first := i == tmp
		tmp, _ = getOwnRoom(minScene.rooms[i][1])
		second := i == tmp
		if first && second {
			finished[i] = true
		}
	}
	finishedRet := finished[0] && finished[1] && finished[2] && finished[3]
	return min, minScene, finishedRet
}

func getOwnRoom(amp string) (roomIdx int, hwIdx int) {
	switch amp {
	case "A":
		return 0, 2
	case "B":
		return 1, 4
	case "C":
		return 2, 6
	case "D":
		return 3, 8
	default:
		return 100, 100
	}
}

func isPathFree(start, stop int, current scene) bool {
	step := (stop - start) / utils.IntAbs(stop-start)
	for pos := start + step; pos != stop+step; pos += step {
		if current.hallway[pos] != "" {
			//something in the way
			return false
		}
	}
	return true
}

func allNextPos(init scene, currentCost int) (output map[scene]int) {
	output = make(map[scene]int)
	for i, s := range init.hallway {
		if s != "" {
			dstRoomIdx, hwIdx := getOwnRoom(s)
			if isPathFree(i, hwIdx, init) {
				//dstRoomPos := 0
				dstRoom := init.rooms[dstRoomIdx]
				for roomIdx, s2 := range dstRoom {
					if s2 != "" && s2 != s {
						break
					}
					if s2 == "" {
						newScene := init
						newScene.hallway[i] = ""
						newScene.rooms[dstRoomIdx][roomIdx] = s
						distance := utils.IntAbs(hwIdx-i) + len(dstRoom) - roomIdx
						output[newScene] = currentCost + distance*int(math.Pow10(dstRoomIdx))
						break
					}
				}
			}
		}
	}
	for i, room := range init.rooms {
		for k, s := range room {

			if k == 0 && room[1] != "" {
				continue
			}
			dstRoomIdx, dsthwIdx := getOwnRoom(s)
			if dstRoomIdx == i {
				other := len(room) - 1 - k
				if room[other] == "" || room[other] == s {
					continue
				}
			}
			hallwayIdx := 2 + i*2
			if s != "" {
				for _, targetPos := range hallwayPos {
					distance := len(room) - k
					if isPathFree(hallwayIdx, targetPos, init) {
						newScene := init
						newScene.hallway[targetPos] = s
						newScene.rooms[i][k] = ""
						distance += utils.IntAbs(targetPos - hallwayIdx)
						output[newScene] = currentCost + distance*int(math.Pow10(dstRoomIdx))
					}
				}
				if dstRoomIdx == i {
					continue
				}
				if isPathFree(hallwayIdx, dsthwIdx, init) {
					distance := len(room) - k
					//dstRoomPos := 0
					dstRoom := init.rooms[dstRoomIdx]
					for dstRoomPos, s2 := range dstRoom {
						if s2 != "" && s2 != s {
							break
						}
						if s2 == "" {
							newScene := init
							newScene.rooms[i][k] = ""
							newScene.rooms[dstRoomIdx][dstRoomPos] = s
							distance += utils.IntAbs(dsthwIdx-hallwayIdx) + len(dstRoom) - dstRoomPos
							output[newScene] = currentCost + distance*int(math.Pow10(dstRoomIdx))
						}
					}
				}
			}
		}
	}
	return
}

func main() {
	file, err := os.Open("2021/day23/testinput2")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}
	//memory := make(map[scene]int)
	initScene := scene{}

	scanner.Scan()
	scanner.Scan()
	scanner.Scan()
	line := strings.Split(strings.Trim(scanner.Text(), "###"), "#")

	for i := 0; i < 4; i++ {
		initScene.rooms[i][LINES-1] = line[i]
	}

	for k := LINES - 1 - 1; k >= 0; k-- {
		scanner.Scan()
		line = strings.Split(scanner.Text(), "#")
		for i := 0; i < 4; i++ {
			initScene.rooms[i][k] = line[i+1]
		}
	}

	allScenes := make(map[scene]int)
	visitedScenes := make(map[scene]int)
	allScenes[initScene] = 0

	min := 0
	var minScene scene
	finished := false
	for ; finished == false; min, minScene, finished = checkScenes(allScenes) {
		newScenes := allNextPos(minScene, allScenes[minScene])
		delete(allScenes, minScene)
		visitedScenes[minScene] = 1
		for key, value := range newScenes {
			if _, ok := visitedScenes[key]; ok {
				continue
			}
			if ret, ok := allScenes[key]; ok {
				if value < ret {
					allScenes[key] = value
				}
			} else {
				allScenes[key] = value
			}

		}
		fmt.Println(min, len(allScenes))
		//break
	}
	fmt.Println(min)

	printScene(minScene)

	//testScene := scene{}
	//testScene.hallway[0] = "C"
	//testScene.hallway[3] = "A"
	//testScene.hallway[9] = "B"
	//testScene.hallway[10] = "D"

	//nextPos := allNextPos(testScene, 0)
	//
	//fmt.Println("----------")
	//for po, _ := range nextPos {
	//	printScene(po)
	//	fmt.Println()
	//}
}
