package main

import (
	"adventofcode/utils"
	"fmt"
	"io/ioutil"
	"strings"
)

/// paraMode assumed to be filled with 0s
func compute(opcode int, paramMode []int, memory *[]int64, itrPtr *int64, relativOffset *int64, input <-chan int64, output chan<- int64) {
	param := (*memory)[*itrPtr+1:]

	//getMemAdd := func(address int64) int64 {
	//	if address >= int64(len((*memory))) {
	//		for int64(len((*memory))) <= address {
	//			(*memory) = append(*memory, 0)
	//		}
	//	}
	//	return address
	//}

	getAdressPtr := func(address int64) *int64 {
		if address >= int64(len((*memory))) {
			for int64(len((*memory))) <= address {
				(*memory) = append(*memory, 0)
			}
		}
		return &(*memory)[address]
	}

	getParam := func(paramIdx int64) *int64 {
		mode := paramMode[paramIdx]
		switch mode {
		case 0:
			return getAdressPtr(param[paramIdx])
		case 1:
			return &param[paramIdx]
		case 2:
			return getAdressPtr(param[paramIdx] + *relativOffset)
		default:
			panic("wrong mode")
		}
	}

	switch opcode {
	case 1:
		*getAdressPtr(param[2]) = *getParam(0) + *getParam(1)
		*itrPtr += 4
	case 2:
		*getAdressPtr(param[2]) = *getParam(0) * *getParam(1)
		*itrPtr += 4
	case 3:
		*getAdressPtr(param[2]) = <-input
		*itrPtr += 2
	case 4:
		output <- *getParam(0)
		*itrPtr += 2
	case 5:
		if *getParam(0) != 0 {
			*itrPtr = *getParam(1)
		} else {
			*itrPtr += 3
		}
	case 6:
		if *getParam(0) == 0 {
			*itrPtr = *getParam(1)
		} else {
			*itrPtr += 3
		}
	case 7:
		if *getParam(0) < *getParam(1) {
			*getAdressPtr(param[2]) = 1
		} else {
			*getAdressPtr(param[2]) = 0
		}
		*itrPtr += 4
	case 8:
		if *getParam(0) == *getParam(1) {
			*getAdressPtr(param[2]) = 1
		} else {
			*getAdressPtr(param[2]) = 0
		}
		*itrPtr += 4
	case 9:
		*relativOffset += *getParam(0)
		*itrPtr += 2
	}

	return
}

func parseOpcode(input []int) (int, []int) {
	opcode := input[len(input)-1]
	if len(input) >= 2 {
		opcode += 10 * input[len(input)-2]
	}
	var param []int
	for i := len(input) - 3; i >= 0; i-- {
		param = append(param, input[i])
	}
	for len(param) < 3 {
		param = append(param, 0)
	}
	return opcode, param
}

func processIntCode(intcode []int64, input <-chan int64, output chan<- int64) {
	ownIntcode := make([]int64, len(intcode))
	copy(ownIntcode, intcode)
	relativOffset := int64(0)
	index := int64(0)
	for true {
		opCode, paramMode := parseOpcode(utils.SplitInt(int(ownIntcode[index])))
		if opCode == 99 {
			close(output)
			return
		}
		compute(opCode, paramMode, &ownIntcode, &index, &relativOffset, input, output)
	}
	return
}

func task1(intcode []int64) []int64 {

	inputCh := make(chan int64)
	outputCh := make(chan int64)

	var output []int64

	go processIntCode(intcode, inputCh, outputCh)

	inputCh <- 1
	//loop til channel is closed
	for out := range outputCh {
		output = append(output, out)
	}
	return output
}

//func task2(intcode []int) int64 {
//	return 1
//}

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
	output1 := task1(intcode)
	fmt.Println("Task 7.1: ", output1)
	//fmt.Println("Task 7.2: ", task2(intcode))
}
