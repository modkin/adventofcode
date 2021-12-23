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

type scene struct {
	hallway [11]string
	rooms   [4][2]string
	cost    int
}

type sceneNoCost struct {
	hallway [11]string
	rooms   [4][2]string
}

func unique(intSlice []scene) []scene {
	keys := make(map[scene]bool)
	list := []scene{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
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
	for i := 1; i >= 0; i-- {
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
	fmt.Println(input.cost)
}

func checkScenes(allScenes []scene) (int, int, bool) {
	min := math.MaxInt
	idx := 0
	finished := false
	for i, s := range allScenes {
		if s.cost < min {
			min = s.cost
			idx = i
		}
	}
	for i := 0; i < 4; i++ {
		tmp, _ := getOwnRoom(allScenes[idx].rooms[i][0])
		first := i == tmp
		tmp, _ = getOwnRoom(allScenes[idx].rooms[i][1])
		second := i == tmp
		if first && second {
			finished = true
		}
	}
	return min, idx, finished
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

func allNextPos(init scene) (output []scene) {
	for i, s := range init.hallway {
		if s != "" {
			//for _, targetPos := range hallwayPos {
			//	if targetPos == i {
			//		continue
			//	}
			//	if isPathFree(i, targetPos, init) {
			//		newScene := init
			//		newScene.hallway[targetPos] = s
			//		newScene.hallway[i] = ""
			//		output = append(output, newScene)
			//	}
			//}
			dstRoomIdx, hwIdx := getOwnRoom(s)
			if isPathFree(i, hwIdx, init) {
				dstRoomPos := 0
				dstRoom := init.rooms[dstRoomIdx]
				if dstRoom[1] == "" && (dstRoom[0] == "" || dstRoom[1] == s) {
					if init.rooms[dstRoomIdx][0] == s {
						dstRoomPos = 1
					}
					newScene := init
					newScene.hallway[i] = ""
					newScene.rooms[dstRoomIdx][dstRoomPos] = s
					distance := utils.IntAbs(hwIdx-i) + 2 - dstRoomPos
					newScene.cost += distance * int(math.Pow10(dstRoomIdx))
					output = append(output, newScene)
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
				other := 1 - k
				if room[other] == "" || room[other] == s {
					continue
				}
			}
			hallwayIdx := 2 + i*2
			if s != "" {
				for _, targetPos := range hallwayPos {
					distance := 2 - k
					if isPathFree(hallwayIdx, targetPos, init) {
						newScene := init
						newScene.hallway[targetPos] = s
						newScene.rooms[i][k] = ""
						distance += utils.IntAbs(targetPos - hallwayIdx)
						newScene.cost += distance * int(math.Pow10(dstRoomIdx))
						output = append(output, newScene)
					}
				}
				if dstRoomIdx == i {
					continue
				}
				if isPathFree(hallwayIdx, dsthwIdx, init) {
					distance := 2 - k
					dstRoomPos := 0
					dstRoom := init.rooms[dstRoomIdx]
					if dstRoom[1] == "" && (dstRoom[0] == "" || dstRoom[1] == s) {
						if init.rooms[dstRoomIdx][0] == s {
							dstRoomPos = 1
						}
						newScene := init
						newScene.rooms[i][k] = ""
						newScene.rooms[dstRoomIdx][dstRoomPos] = s
						distance += utils.IntAbs(dsthwIdx-hallwayIdx) + 2 - dstRoomPos
						newScene.cost += distance * int(math.Pow10(dstRoomIdx))
						output = append(output, newScene)
					}
				}
			}
		}
	}
	return
}

func main() {
	file, err := os.Open("2021/day23/testinput")
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
		initScene.rooms[i][1] = line[i]
	}

	scanner.Scan()
	line = strings.Split(scanner.Text(), "#")
	for i := 0; i < 4; i++ {
		initScene.rooms[i][0] = line[i+1]
	}

	//fmt.Println(initScene)

	allScenes := make([]scene, 0)
	allScenes = append(allScenes, initScene)
	//sceneCosts := make(map[sceneNoCost]int)

	min, idx := 0, 0
	finished := false
	for ; finished == false; min, idx, finished = checkScenes(allScenes) {
		newScenes := allNextPos(allScenes[idx])
		allScenes = append(allScenes[:idx], allScenes[idx+1:]...)
		allScenes = append(allScenes, newScenes...)
		for _, ns := range newScenes {
			notVisited := true
			for i, olds := range allScenes {
				if olds.hallway == ns.hallway && olds.rooms == ns.rooms {
					notVisited = false
					if ns.cost < olds.cost {
						allScenes[i].cost = ns.cost
					}
				}
			}
			if notVisited {
				allScenes = append(allScenes, ns)
			}
		}
		fmt.Println(min)
	}
	fmt.Println(min)

	//testScene := scene{}
	//testScene.rooms[2][0] = "A"
	//testScene.hallway[7] = "B"
	//nextPos := allNextPos(testScene)
	//printScene(testScene)
	//fmt.Println()
	//for _, po := range nextPos {
	//	printScene(po)
	//	fmt.Println()
	//}
	//for i := 0; i < 4; i++ {
	//	fmt.Println(math.Pow10(i))
	//}
}
