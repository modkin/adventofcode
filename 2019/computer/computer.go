package computer

import "adventofcode/utils"

func compute(opcode int, paramMode []int, memory []int, itrPtr *int, input <-chan int, output chan<- int) {
	param := memory[*itrPtr+1:]

	getParam := func(paramIdx int) int {
		mode := paramMode[paramIdx]
		switch mode {
		case 0:
			return memory[param[paramIdx]]
		case 1:
			return param[paramIdx]
		default:
			panic("wrong mode")
		}
	}

	switch opcode {
	case 1:
		memory[param[2]] = getParam(0) + getParam(1)
		*itrPtr += 4
	case 2:
		memory[param[2]] = getParam(0) * getParam(1)
		*itrPtr += 4
	case 3:
		memory[param[0]] = <-input
		*itrPtr += 2
	case 4:
		output <- getParam(0)
		*itrPtr += 2
	case 5:
		if getParam(0) != 0 {
			*itrPtr = getParam(1)
		} else {
			*itrPtr += 3
		}
	case 6:
		if getParam(0) == 0 {
			*itrPtr = getParam(1)
		} else {
			*itrPtr += 3
		}
	case 7:
		if getParam(0) < getParam(1) {
			memory[param[2]] = 1
		} else {
			memory[param[2]] = 0
		}
		*itrPtr += 4
	case 8:
		if getParam(0) == getParam(1) {
			memory[param[2]] = 1
		} else {
			memory[param[2]] = 0
		}
		*itrPtr += 4
		//case 9:
		//	*relativOffset += getParam(0 )
		//	*itrPtr += 2
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

func ProcessIntCode(intcode []int, input <-chan int, output chan<- int) {
	ownIntcode := make([]int, len(intcode))
	copy(ownIntcode, intcode)
	index := 0
	for true {
		opCode, paramMode := parseOpcode(utils.SplitInt(ownIntcode[index]))
		if opCode == 99 {
			return
		}
		compute(opCode, paramMode, ownIntcode, &index, input, output)
	}
	return
}
