package main

import (
	"adventofcode/2019/computer"
	"adventofcode/utils"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func printPaintMap(paintMap map[[2]int]string) {
	minX, minY, maxX, maxY := math.MaxInt32, math.MaxInt32, math.MinInt32, math.MinInt32
	for pos, _ := range paintMap {
		if pos[0] < minX {
			minX = pos[0]
		}
		if pos[0] > maxX {
			maxX = pos[0]
		}
		if pos[1] < minY {
			minY = pos[1]
		}
		if pos[1] > maxY {
			maxY = pos[1]
		}
	}
	fmt.Println("==================================================")
	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			tmp := [2]int{x, y}
			if x == 0 && y == 0 {
				fmt.Print("X")
			} else {
				if val, ok := paintMap[tmp]; ok {
					fmt.Print(val)
				} else {
					fmt.Print(" ")
				}
			}

		}
		fmt.Println()
	}
	fmt.Println("==================================================")
}

func getDirection(dir int) [2]int {
	switch dir {

	case 1: //north
		return [2]int{0, 1}
	case 2: //south
		return [2]int{0, -1}
	case 3: //west
		return [2]int{-1, 0}
	case 4: // east
		return [2]int{1, 0}
	}
	panic("Wrong direction")
}

type Robot struct {
	code     []int64
	distance int
	input    chan int64
	output   chan int64
	squence  []int
}

func main() {
	content, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	contentString := strings.Split(string(content), ",")
	intcode := make([]int64, len(contentString))
	for pos, elem := range contentString {
		intcode[pos] = utils.ToInt64(elem)
	}
	//inputCh := make(chan int64)
	//outputCh := make(chan int64)
	//quit := make(chan bool)

	shipMap := make(map[[2]int]string)

	//go computer.ProcessIntCode(intcode, inputCh, outputCh, quit)

	//foundOxy := false
	//currentPos := [2]int{0, 0}
	//for foundOxy == false {
	//
	//	nextInput := rand.Intn(4) + 1
	//	inputCh <- int64(nextInput)
	//	output := <-outputCh
	//	switch output {
	//	//wall in front
	//	case 0:
	//		newWall := [2]int{currentPos[0] + getDirection(nextInput)[0], currentPos[1] + getDirection(nextInput)[1]}
	//		shipMap[newWall] = "#" //"█"
	//	case 1:
	//		shipMap[currentPos] = "."
	//		currentPos = [2]int{currentPos[0] + getDirection(nextInput)[0], currentPos[1] + getDirection(nextInput)[1]}
	//	case 2:
	//		shipMap[currentPos] = "."
	//		oxyPos = [2]int{currentPos[0] + getDirection(nextInput)[0], currentPos[1] + getDirection(nextInput)[1]}
	//		shipMap[oxyPos] = "D"
	//		foundOxy = true
	//	}
	//}
	//printPaintMap(shipMap)
	var distanceToOxy int
	lastPos := [2]int{0, 0}
	postions := make(map[[2]int]Robot)
	postions[lastPos] = Robot{
		code:     intcode,
		distance: 0,
		input:    make(chan int64),
		output:   make(chan int64),
		squence:  nil,
	}
	go computer.ProcessIntCode(postions[lastPos].code, postions[lastPos].input, postions[lastPos].output, make(chan bool))
	looking := true
	for looking {
		looking = false
		newPostions := make(map[[2]int]Robot)
		for pos, robot := range postions {
			for i := 1; i < 5; i++ {
				dir := getDirection(i)
				lookingAtPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
				if _, ok := shipMap[lookingAtPos]; !ok {
					looking = true
					newRobot := Robot{
						code:     intcode,
						distance: robot.distance + 1,
						input:    make(chan int64),
						output:   make(chan int64),
						squence:  utils.CopyIntSlice(robot.squence),
					}
					go computer.ProcessIntCode(intcode, newRobot.input, newRobot.output, make(chan bool))
					for _, step := range robot.squence {
						newRobot.input <- int64(step)
						<-newRobot.output
					}
					newRobot.input <- int64(i)
					output := <-newRobot.output
					switch output {
					//wall in front
					case 0:
						shipMap[lookingAtPos] = "#"
					case 1:
						shipMap[lookingAtPos] = "."
						newRobot.distance += 1
						newRobot.squence = append(newRobot.squence, i)
						newPostions[lookingAtPos] = newRobot
					case 2:
						shipMap[lookingAtPos] = "D"
						if distanceToOxy == 0 {
							distanceToOxy = robot.distance + 1
						}
						newRobot.distance += 1
						newRobot.squence = append(newRobot.squence, i)
						newPostions[lookingAtPos] = newRobot
					}
				}
			}
		}
		postions = newPostions
	}

	printPaintMap(shipMap)
	//fmt.Println("Task 15.1: ", distanceToOxy)

	startPos := [2]int{0, 0}
	var oxyPos [2]int
	postions2 := make(map[[2]int]int)
	postions2[startPos] = 0
	looking = true
	shortest := 0
	for looking {
		nextPositions := make(map[[2]int]int)
		for pos, distance := range postions2 {
			for i := 1; i < 5; i++ {
				dir := getDirection(i)
				lookingAtPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
				if shipMap[lookingAtPos] == "." {
					nextPositions[lookingAtPos] = distance + 1
				} else if shipMap[lookingAtPos] == "D" {
					looking = false
					oxyPos = lookingAtPos
					shortest = distance + 1
				}
			}
		}
		postions2 = nextPositions
	}
	fmt.Println("Task 15.1: ", shortest)

	postions3 := make(map[[2]int]int)
	postions3[oxyPos] = 0
	looking = true
	longest := 0
	for looking {
		looking = false
		nextPositions := make(map[[2]int]int)
		for pos, distance := range postions3 {
			looking = true
			for i := 1; i < 5; i++ {
				dir := getDirection(i)
				lookingAtPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
				if shipMap[lookingAtPos] == "." || shipMap[lookingAtPos] == "x" {
					nextPositions[lookingAtPos] = distance + 1
					shipMap[pos] = ";"
					if (distance + 1) > longest {
						longest = distance + 1
					}
				}

			}
		}
		postions3 = nextPositions
	}
	printPaintMap(shipMap)
	fmt.Println("Task 15.2: ", longest)

}
