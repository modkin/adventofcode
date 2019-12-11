package main

import (
	"adventofcode/2019/computer"
	"adventofcode/utils"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func rotate90left(vec [2]int) (ret [2]int) {
	ret[0] = -1 * vec[1]
	ret[1] = vec[0]
	return
}

func rotate90right(vec [2]int) (ret [2]int) {
	ret[0] = vec[1]
	ret[1] = -1 * vec[0]
	return
}

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

	runRobot := func(start int, print bool) int {

		pos := [2]int{0, 0}
		direction := [2]int{0, 1}
		paintMap := make(map[[2]int]int)
		inputCh := make(chan int64, 1)
		outputCh := make(chan int64)
		quit := make(chan bool)

		go computer.ProcessIntCode(intcode, inputCh, outputCh, quit)

		paintMap[pos] = start
		//inputCh <- int64(paintMap[pos])
		counter := 0
		painted := make(map[[2]int]bool)
		running := true
		for running {
			if color, ok := paintMap[pos]; ok {
				inputCh <- int64(color)
			} else {
				paintMap[pos] = 0
				inputCh <- int64(0)
			}
			select {
			case val := <-outputCh:
				if painted[pos] {
					paintMap[pos] = int(val)
				} else if !painted[pos] && val == 1 {
					painted[pos] = true
					paintMap[pos] = 1
					counter++
				} else if !painted[pos] && val == 0 {
					paintMap[pos] = 0
				} else {
					fmt.Println("Something wrong")
				}

				rotate := <-outputCh
				if rotate == 0 {
					direction = rotate90left(direction)
				} else if rotate == 1 {
					direction = rotate90right(direction)
				}
				pos[0] += direction[0]
				pos[1] += direction[1]

			case <-quit:
				running = false
			}
		}
		if print == true {
			printPaintMap(painted)
		}
		return counter
	}
	runRobot(0, false)
	fmt.Println("Task 11.1", runRobot(0, false))
	fmt.Println("Task 11.2")
	runRobot(1, true)

}
