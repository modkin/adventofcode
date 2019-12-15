package main

import (
	"adventofcode/2019/computer"
	"adventofcode/utils"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func printPaintMap(paintMap map[[2]int]bool) {
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
		for x := minY; x < maxX; x++ {
			tmp := [2]int{x, y}
			if _, ok := paintMap[tmp]; ok {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
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

	intcode[0] = 2

	inputCh := make(chan int64)
	outputCh := make(chan int64)
	quit := make(chan bool)

	shipMap := make(map[[2]int]string)

	go computer.ProcessIntCode(intcode, inputCh, outputCh, quit)

	foundOxy := false
	for foundOxy == false {
		nextInput
	}

}
