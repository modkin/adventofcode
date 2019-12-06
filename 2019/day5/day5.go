package main

import (
	"adventofcode/utils"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

/// paraMode assumed to be filled with 0s
func compute(opcode int, param []int, paramMode []int, memory []int, itrPtr *int, input int) (ret int) {
	ret = -1
	var input1, input2 int
	if paramMode[0] == 0 {
		input1 = memory[param[0]]
	} else if paramMode[0] == 1 {
		input1 = param[0]
	}
	if opcode != 3 && opcode != 4 {
		if paramMode[1] == 0 {
			input2 = memory[param[1]]
		} else if paramMode[1] == 1 {
			input2 = param[1]
		}
	}

	if opcode == 1 {
		memory[param[2]] = input1 + input2
		*itrPtr += 4
	} else if opcode == 2 {
		memory[param[2]] = input1 * input2
		*itrPtr += 4
	} else if opcode == 3 {
		memory[param[0]] = input
		*itrPtr += 2
	} else if opcode == 4 {
		ret = memory[param[0]]
		*itrPtr += 2
	} else if opcode == 5 {
		if input1 != 0 {
			*itrPtr = input2
		} else {
			*itrPtr += 3
		}
	} else if opcode == 6 {
		if input1 == 0 {
			*itrPtr = input2
		} else {
			*itrPtr += 3
		}
	} else if opcode == 7 {
		if input1 < input2 {
			memory[param[2]] = 1
		} else {
			memory[param[2]] = 0
		}
		*itrPtr += 4
	} else if opcode == 8 {
		if input1 == input2 {
			memory[param[2]] = 1
		} else {
			memory[param[2]] = 0
		}
		*itrPtr += 4
	}
	return
}

func splitInt(in int) []int {
	count := 0
	copyIn := in
	for in != 0 {
		in /= 10
		count++
	}
	//mt.Println("count", count)

	output := make([]int, count)
	for i := 0; i < count; i++ {
		output[i] = copyIn / int(math.Pow(10, float64(count-1-i))) % 10
	}
	return output
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
	//fmt.Println("Opcode ", opcode, " param ", param)
	return opcode, param
}

func processIntCode(intcode []int, input int) (outputs []int) {
	index := 0
	for true {
		opCode, paramMode := parseOpcode(splitInt(intcode[index]))
		if opCode == 99 {
			//fmt.Println("Program Finished")
			return
		}
		ret := compute(opCode, intcode[index+1:], paramMode, intcode, &index, input)
		if ret != -1 {
			outputs = append(outputs, ret)
		}
		//fmt.Println("index: ", index)
	}
	return
}

func task1() int {
	content, err := ioutil.ReadFile("2019/day5/input")
	if err != nil {
		panic(err)
	}
	contentString := strings.Split(string(content), ",")
	intcode := make([]int, len(contentString))
	for pos, elem := range contentString {
		intcode[pos] = utils.ToInt(elem)
	}
	outputs := processIntCode(intcode, 1)
	return outputs[len(outputs)-1]
}

func task2() int {
	content, err := ioutil.ReadFile("2019/day5/input")
	if err != nil {
		panic(err)
	}
	contentString := strings.Split(string(content), ",")
	intcode := make([]int, len(contentString))
	for pos, elem := range contentString {
		intcode[pos] = utils.ToInt(elem)
	}
	outputs := processIntCode(intcode, 5)
	return outputs[len(outputs)-1]
}

func main() {
	fmt.Println("Task 5.1: ", task1())
	fmt.Println("Task 5.2: ", task2())
}
