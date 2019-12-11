package computer

import "adventofcode/utils"

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

func ProcessIntCode(intcode []int64, input <-chan int64, output chan<- int64, closeOutput bool) {
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
		opCode, paramMode := parseOpcode(utils.SplitInt(int(memory[getMemAdd(itrPtr)])))
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
			if closeOutput {
				close(output)
			}
			return

		}
	}
	return
}
