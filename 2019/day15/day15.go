package main

import (
	"adventofcode/2019/computer"
	"adventofcode/utils"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
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
	inputCh := make(chan int64)
	outputCh := make(chan int64)
	quit := make(chan bool)

	shipMap := make(map[[2]int]string)

	go computer.ProcessIntCode(intcode, inputCh, outputCh, quit)

	foundOxy := false
	currentPos := [2]int{0, 0}
	var oxyPos [2]int
	for foundOxy == false {

		nextInput := rand.Intn(4) + 1
		inputCh <- int64(nextInput)
		output := <-outputCh
		switch output {
		//wall in front
		case 0:
			newWall := [2]int{currentPos[0] + getDirection(nextInput)[0], currentPos[1] + getDirection(nextInput)[1]}
			shipMap[newWall] = "#" //"█"
		case 1:
			shipMap[currentPos] = "."
			currentPos = [2]int{currentPos[0] + getDirection(nextInput)[0], currentPos[1] + getDirection(nextInput)[1]}
		case 2:
			shipMap[currentPos] = "."
			oxyPos = [2]int{currentPos[0] + getDirection(nextInput)[0], currentPos[1] + getDirection(nextInput)[1]}
			shipMap[oxyPos] = "D"
			foundOxy = true
		}
	}
	printPaintMap(shipMap)
	lastPos := [2]int{0, 0}
	postions := make(map[[2]int]int)
	postions[lastPos] = 0
	looking := true
	shortest := 0
	for looking {
		newPostions := make(map[[2]int]int)
		for pos, distance := range postions {
			for i := 1; i < 5; i++ {
				dir := getDirection(i)
				lookingAtPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
				if shipMap[lookingAtPos] == "." {
					newPostions[lookingAtPos] = distance + 1
				} else if shipMap[lookingAtPos] == "D" {
					looking = false
					shortest = distance + 1
				}
			}
		}
		postions = newPostions
	}
	fmt.Println("Task 15.1: ", shortest)

}
