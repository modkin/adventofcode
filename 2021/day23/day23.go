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

func SceneFinished(input scene) bool {
	finished := true
	for i, room := range input.rooms {
		for _, s := range room {
			roomIdx, _ := getOwnRoom(s)
			finished = finished && (roomIdx == i)
		}
	}
	return finished
}

func checkScenes(allScenes map[scene]int) (int, scene, bool) {
	min := math.MaxInt
	var minScene scene

	for i, s := range allScenes {
		if s < min {
			min = s
			minScene = i
		}
	}
	return min, minScene, SceneFinished(minScene)
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

func roomDone(room [LINES]string, idx int) bool {
	for _, i := range room {
		if i != "" {
			roomIdx, _ := getOwnRoom(i)
			if roomIdx != idx {
				return false
			}
		}
	}
	return true
}

func moveIn(init scene) (scene, int) {
	newScene := init
	newCost := 0
	moved := true
	for moved {
		moved = false
		for i, s := range newScene.hallway {
			if s != "" {
				dstRoomIdx, hwIdx := getOwnRoom(s)
				if isPathFree(i, hwIdx, newScene) {
					//dstRoomPos := 0
					dstRoom := newScene.rooms[dstRoomIdx]
					for roomIdx, s2 := range dstRoom {
						if s2 != "" && s2 != s {
							break
						}
						if s2 == "" {
							newScene.hallway[i] = ""
							newScene.rooms[dstRoomIdx][roomIdx] = s
							distance := utils.IntAbs(hwIdx-i) + len(dstRoom) - roomIdx
							newCost += distance * int(math.Pow10(dstRoomIdx))
							moved = true
							break
						}
					}
				}
			}
		}

		for i, room := range newScene.rooms {
			for k := len(room) - 1; k >= 0; k-- {
				s := room[k]
				if s != "" {
					dstRoomIdx, dsthwIdx := getOwnRoom(s)
					hallwayIdx := 2 + i*2
					if hallwayIdx == dsthwIdx {
						break
					}
					if isPathFree(hallwayIdx, dsthwIdx, newScene) {
						distance := len(room) - k
						//dstRoomPos := 0
						dstRoom := newScene.rooms[dstRoomIdx]
						for dstRoomPos, s2 := range dstRoom {
							if s2 != "" && s2 != s {
								break
							}
							if s2 == "" {
								newScene.rooms[i][k] = ""
								newScene.rooms[dstRoomIdx][dstRoomPos] = s
								distance += utils.IntAbs(dsthwIdx-hallwayIdx) + (len(dstRoom) - dstRoomPos)
								newCost += distance * int(math.Pow10(dstRoomIdx))
								moved = true
								break
							}
						}
					}
					break
				}
			}
		}
	}
	return newScene, newCost
}

func allNextPos(init scene, currentCost int) (output map[scene]int) {
	output = make(map[scene]int)
	for i, room := range init.rooms {
		if roomDone(room, i) {
			continue
		}
		hallwayIdx := 2 + i*2
		for k := len(room) - 1; k >= 0; k-- {
			s := room[k]
			if s != "" {
				dstRoomIdx, _ := getOwnRoom(s)
				for _, targetPos := range hallwayPos {
					distance := len(room) - k
					if isPathFree(hallwayIdx, targetPos, init) {
						newScene := init
						newCost := 0
						newScene.hallway[targetPos] = s
						newScene.rooms[i][k] = ""
						distance += utils.IntAbs(targetPos - hallwayIdx)
						newCost += distance * int(math.Pow10(dstRoomIdx))
						moveInScene, moveInCost := moveIn(newScene)
						totalCost := currentCost + newCost + moveInCost
						if oldCost, ok := output[moveInScene]; ok {
							if oldCost > totalCost {
								output[moveInScene] = totalCost
							}
						} else {
							output[moveInScene] = totalCost
						}
					}
				}
				break
			}
		}
	}
	return
}

func main() {
	file, err := os.Open("2021/day23/input")
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

	initScene.hallway[0] = "A"

	initScene.hallway[7] = "B"
	initScene.hallway[9] = "B"
	initScene.hallway[10] = "D"

	initScene.rooms[3][3] = ""
	initScene.rooms[3][2] = ""
	initScene.rooms[2][3] = ""
	initScene.rooms[2][2] = ""

	//initScene.hallway[1] = "A"
	//initScene.rooms[1][3] = "."
	//initScene.rooms[1][2] = "."
	//initScene.rooms[2][1] = "C"
	//initScene.rooms[2][2] = "C"

	//initScene.rooms[1][1] = "B"
	//initScene.rooms[2][0] = "C"
	//initScene.rooms[2][1] = "C"
	//initScene.rooms[3][0] = "A"
	//initScene.rooms[3][1] = "D"

	//min := 0
	//var minScene scene
	//finished := false
	min, minScene, finished := checkScenes(allScenes)
	for ; finished == false; min, minScene, finished = checkScenes(allScenes) {
		if minScene == initScene {
			fmt.Println("SLKFJLK:DFS:", min)
			printScene(minScene)

		}
		visitedScenes[minScene] = 1
		newScenes := allNextPos(minScene, allScenes[minScene])
		delete(allScenes, minScene)
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
	fmt.Println("Day 23:", min)

	//printScene(minScene)

	initScene.hallway[3] = "B"
	initScene.rooms[3][1] = ""
	//initScene.rooms[0][1] = ""
	//initScene.rooms[1][0] = "B"
	//initScene.rooms[1][1] = "B"
	//initScene.rooms[2][0] = "C"
	//initScene.rooms[2][1] = "C"
	//initScene.rooms[3][0] = "A"
	//initScene.rooms[3][1] = "D"

	//testScene := scene{}
	//testScene.rooms[3][1] = "A"
	//testScene.rooms[3][0] = "A"
	//testScene.hallway[7] = "D"
	//testScene.hallway[5] = "D"
	//
	//fmt.Println("START")
	//printScene(testScene)
	//nextPos := allNextPos(testScene, 0)
	//
	//fmt.Println("----------")
	//for po, cost := range nextPos {
	//	printScene(po)
	//	fmt.Println(cost)
	//	fmt.Println()
	//}
}
