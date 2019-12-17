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
	maxX, maxY := math.MinInt32, math.MinInt32
	for pos, _ := range paintMap {
		if pos[0] > maxX {
			maxX = pos[0]
		}

		if pos[1] > maxY {
			maxY = pos[1]
		}
	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x < maxX; x++ {
			fmt.Print(paintMap[[2]int{x, y}])
		}
		fmt.Println()
	}
}

func findItersections(paintMap map[[2]int]string) (result int) {
	maxX, maxY := math.MinInt32, math.MinInt32
	for pos, _ := range paintMap {
		if pos[0] > maxX {
			maxX = pos[0]
		}

		if pos[1] > maxY {
			maxY = pos[1]
		}
	}
	for y := 1; y <= maxY-1; y++ {
		for x := 1; x < maxX-1; x++ {
			if paintMap[[2]int{x, y}] == "#" {
				if paintMap[[2]int{x - 1, y}] == "#" && paintMap[[2]int{x + 1, y}] == "#" && paintMap[[2]int{x, y - 1}] == "#" && paintMap[[2]int{x, y + 1}] == "#" {
					paintMap[[2]int{x, y}] = "O"
					result += y * x
				}
			}
		}
	}
	return
}

func runCamera(shipMap map[[2]int]string, outputCh <-chan int64, quit <-chan bool) {
	running := true
	x := 0
	y := 0
	for running {
		select {
		case input := <-outputCh:
			if input == 10 {
				y++
				x = 0
			} else {
				shipMap[[2]int{x, y}] = string(rune(input))
				x++
			}
		case <-quit:
			running = false
		}
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
	inputCh := make(chan int64)
	outputCh := make(chan int64)
	quit := make(chan bool)

	shipMap := make(map[[2]int]string)

	go computer.ProcessIntCode(intcode, inputCh, outputCh, quit)

	runCamera(shipMap, outputCh, quit)
	task1 := findItersections(shipMap)
	//printPaintMap(shipMap)
	fmt.Println("Task 17.1: ", task1)

	maxX, maxY := math.MinInt32, math.MinInt32
	for pos, _ := range shipMap {
		if pos[0] > maxX {
			maxX = pos[0]
		}
		if pos[1] > maxY {
			maxY = pos[1]
		}
	}

	intcode[0] = 2
	go computer.ProcessIntCode(intcode, inputCh, outputCh, quit)
	//for i := 0; i < maxX*maxY; i++ {
	//	<-outputCh
	//}
	//runCamera(shipMap,outputCh,quit)

	mainProgramm := []rune("A\n")
	aProgramm := []rune("R,8\n8\n8\n")
	videoFeed := []rune("n\n")
	var total []rune
	total = append(append(mainProgramm, aProgramm...), videoFeed...)
	counter := 0
	running := true
	for running {
		select {
		case <-outputCh:
		case <-quit:
			running = false
		case inputCh <- int64(total[counter]):
			fmt.Println(string(total[counter]), ": ", total[counter])
			counter++
			if counter == len(total) {
				running = false
			}
		}
	}

	shipMap = make(map[[2]int]string)

	runCamera(shipMap, outputCh, quit)
	findItersections(shipMap)
	printPaintMap(shipMap)
}
