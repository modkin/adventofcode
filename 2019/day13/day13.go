package main

import (
	"adventofcode/2019/computer"
	"adventofcode/utils"
	"fmt"
	"io/ioutil"
	"strings"
)

func processIntCode(intcode []int64, input <-chan int64, output chan<- int64, quit chan<- bool) {
	memory := make([]int64, len(intcode))
	copy(memory, intcode)

	getMemAdd := func(address int64) int64 {
		for int64(len(memory)) <= address {
			memory = append(memory, 0)
		}
		return address
	}

	itrPtr := int64(0)
	relativOffset := int64(0)
	for true {
		param := memory[getMemAdd(itrPtr)+1:]
		opCode, paramMode := computer.ParseOpcode(utils.SplitInt(int(memory[getMemAdd(itrPtr)])))
		getParam := func(paramIdx int) int64 {
			mode := paramMode[paramIdx]
			switch mode {
			case 0:
				return memory[getMemAdd(param[paramIdx])]
			case 1:
				return param[paramIdx]
			case 2:
				return memory[getMemAdd(param[paramIdx]+relativOffset)]
			default:
				panic("wrong mode")
			}
		}
		getWriteAddress := func(paramIdx int) int64 {
			mode := paramMode[paramIdx]
			switch mode {
			case 0:
				return getMemAdd(param[paramIdx])
			case 2:
				return getMemAdd(param[paramIdx] + relativOffset)
			default:
				panic("wrong mode")
			}
		}

		switch opCode {
		case 1:
			memory[getWriteAddress(2)] = getParam(0) + getParam(1)
			itrPtr += 4
		case 2:
			memory[getWriteAddress(2)] = getParam(0) * getParam(1)
			itrPtr += 4
		case 3:
			//time.Sleep(100000000)
			//fmt.Println("input")
			quit <- true
			memory[getWriteAddress(0)] = <-input
			itrPtr += 2
		case 4:
			output <- getParam(0)
			itrPtr += 2
		case 5:
			if getParam(0) != 0 {
				itrPtr = getParam(1)
			} else {
				itrPtr += 3
			}
		case 6:
			if getParam(0) == 0 {
				itrPtr = getParam(1)
			} else {
				itrPtr += 3
			}
		case 7:
			if getParam(0) < getParam(1) {
				memory[getWriteAddress(2)] = 1
			} else {
				memory[getWriteAddress(2)] = 0
			}
			itrPtr += 4
		case 8:
			if getParam(0) == getParam(1) {
				memory[getWriteAddress(2)] = 1
			} else {
				memory[getWriteAddress(2)] = 0
			}
			itrPtr += 4
		case 9:
			relativOffset += getParam(0)
			itrPtr += 2
		case 99:
			quit <- false
			return

		}
	}
	return
}

func printGame(game [24][42]int) int {
	paddleX, ballX := 0, 0
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
				paddleX = x
				fmt.Print("█")
			case 4:
				ballX = x
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
	if paddleX < ballX {
		return 1
	} else if paddleX > ballX {
		return -1
	} else {
		return 0
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

	go processIntCode(intcode, inputCh, outputCh, quit)

	var display [24][42]int
	blockCount := 0
	running := true
	//finishFram := false
	init := true
	timer := 0
	move := 0
	for running {
		select {
		case x := <-outputCh:
			y := <-outputCh
			id := <-outputCh
			if x == -1 {
				fmt.Println("Score: ", id)
				if init {
					init = false
				}
			} else {
				display[y][x] = int(id)
			}
			if id == 2 {
				blockCount++
			}
			if x == 41 && y == 23 && init {

			}
			if !init && x != -1 {
				move = printGame(display)
				timer++
			}

		case waitingInput := <-quit:
			if waitingInput == true {
				//var input int
				//fmt.Scan(&input)
				//inputCh <- int64(input)
				inputCh <- int64(move)
			} else {
				running = false
			}
		}
	}
	fmt.Println(blockCount)
}
