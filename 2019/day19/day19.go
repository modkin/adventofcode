package main

import (
	"adventofcode/2019/computer"
	"adventofcode/utils"
	"fmt"
	"io/ioutil"
	"strings"
)

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

	//shipMap := make(map[[2]int]string)

	counter := 0
	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			inputCh := make(chan int64)
			outputCh := make(chan int64)
			quit := make(chan bool)
			go computer.ProcessIntCode(intcode, inputCh, outputCh, quit)
			inputCh <- int64(x)
			inputCh <- int64(y)
			out := <-outputCh
			if out == 1 {
				counter++
			}

		}
	}
	fmt.Println("Task 18.1: ", counter)

	possibleStarts := make(map[[2]int]bool)
	var beam [10000][10000]int
	stop := false
	xLower, xUpper := 0, 100
	for y := 0; y < 10000; y++ {
		if stop {
			break
		}
		for x := 0; x < 10000; x++ {
			if x < xLower || x > xUpper {
				continue
			}
			inputCh := make(chan int64)
			outputCh := make(chan int64)
			quit := make(chan bool)
			go computer.ProcessIntCode(intcode, inputCh, outputCh, quit)
			inputCh <- int64(x)
			inputCh <- int64(y)
			out := <-outputCh
			beam[x][y] = int(out)
			if x > 5 && out == 1 && beam[x-1][y] == 0 {
				xLower = x - 5
			}
			if out == 1 && beam[x+1][y] == 0 {
				xUpper = x + 5
			}
			if x >= 99 {
				if beam[x-99][y] == 1 {
					possibleStarts[[2]int{x - 99, y}] = true
				}
			}
			if _, ok := possibleStarts[[2]int{x, y - 00}]; ok {
				fmt.Println(x, " ", y)
				stop = true
				break
			}
			//fmt.Print(out)
		}
		//fmt.Println()
	}
}
