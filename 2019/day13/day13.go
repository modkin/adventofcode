package main

import (
	"adventofcode/2019/computer"
	"adventofcode/utils"
	"fmt"
	"io/ioutil"
	"strings"
)

func printGame(game [24][42]int) {

	for y := 0; y < 24; y++ {
		for x := 0; x < 42; x++ {
			switch game[y][x] {
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Print("|")
			case 2:
				fmt.Print("#")
			case 3:
				fmt.Print("_")
			case 4:
				fmt.Print("*")
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

	go computer.ProcessIntCode(intcode, inputCh, outputCh, quit)

	var display [24][42]int
	blockCount := 0
	running := true
	//finishFram := false
	init := true
	timer := 0
	for running {
		select {
		case x := <-outputCh:
			y := <-outputCh
			id := <-outputCh
			if x == -1 {
				fmt.Println("Score: ", id)
				//inputCh <- 0
			} else {
				display[y][x] = int(id)
			}
			if id == 2 {
				blockCount++
			}
			if x == 41 && y == 23 && init {
				init = false
			}
			if !init {
				timer++
				if timer%2 == 0 {
					printGame(display)
					var input int
					fmt.Scan(&input)
					inputCh <- int64(input)
				}

			}

		case <-quit:
			running = false
		}
	}
	fmt.Println(blockCount)
}
