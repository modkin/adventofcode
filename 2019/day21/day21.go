package main

import (
	"adventofcode/2019/computer"
	"adventofcode/utils"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func runCamera2(shipMap map[[2]int]string, outputCh <-chan int64, quit <-chan bool) {
	running := true
	x := 0
	y := 0
	for running {
		select {
		case input := <-outputCh:
			if input == 10 {
				return
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

func runCamera(shipMap map[[2]int]string, outputCh <-chan int64, quit <-chan bool) {
	running := true
	x := 0
	y := 0
	for running {
		select {
		case input := <-outputCh:
			fmt.Println(input)
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
		for x := 0; x <= maxX; x++ {
			fmt.Print(paintMap[[2]int{x, y}])
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
	inputCh := make(chan int64)
	outputCh := make(chan int64)
	quit := make(chan bool)

	shipMap := make(map[[2]int]string)

	go computer.ProcessIntCode(intcode, inputCh, outputCh, quit)

	runCamera2(shipMap, outputCh, quit)
	printPaintMap(shipMap)
	//:= []rune("NOT A J\nWALK\n")
	walk := []rune("NOT A J\nNOT B T\nOR T J\nNOT C T\nOR T J\nAND D J\nWALK\n")
	for _, char := range walk {
		inputCh <- int64(char)
	}
	//out := <- outputCh
	fmt.Println("DONE")

	runCamera(shipMap, outputCh, quit)
	//printPaintMap(shipMap)

}
