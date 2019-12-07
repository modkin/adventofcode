package main

import (
	"adventofcode/utils"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type inputWrap struct {
	input []int
}

/// paraMode assumed to be filled with 0s
func compute(opcode int, param []int, paramMode []int, memory []int, itrPtr *int, input *inputWrap) (ret int) {
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

	switch opcode {
	case 1:
		memory[param[2]] = input1 + input2
		*itrPtr += 4
	case 2:
		memory[param[2]] = input1 * input2
		*itrPtr += 4
	case 3:
		memory[param[0]] = input.input[0]
		//pop first element
		input.input = input.input[1:]
		*itrPtr += 2
	case 4:
		if paramMode[0] == 0 {
			ret = memory[param[0]]
		} else if paramMode[0] == 1 {
			ret = param[0]
		}
		*itrPtr += 2
	case 5:
		if input1 != 0 {
			*itrPtr = input2
		} else {
			*itrPtr += 3
		}
	case 6:
		if input1 == 0 {
			*itrPtr = input2
		} else {
			*itrPtr += 3
		}
	case 7:
		if input1 < input2 {
			memory[param[2]] = 1
		} else {
			memory[param[2]] = 0
		}
		*itrPtr += 4
	case 8:
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

func processIntCode(intcode []int, input *inputWrap) (outputs []int) {
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

func checkValid(input [5]int, offset int) bool {
	counter := [5]int{0, 0, 0, 0, 0}
	for i := 0; i < 5; i++ {
		counter[input[i]-offset]++
	}
	for i := 0; i < 5; i++ {
		if counter[i] > 1 {
			return false
		}
	}
	return true
}

func generatePSS() (pss [][5]int) {
	for i1 := 0; i1 < 5; i1++ {
		for i2 := 0; i2 < 5; i2++ {
			for i3 := 0; i3 < 5; i3++ {
				for i4 := 0; i4 < 5; i4++ {
					for i5 := 0; i5 < 5; i5++ {
						tmp := [5]int{i1, i2, i3, i4, i5}
						if checkValid(tmp, 0) {
							pss = append(pss, [5]int{i1, i2, i3, i4, i5})
						}
					}
				}
			}
		}
	}
	return
}

func generatePSS2() (pss [][5]int) {
	for i1 := 5; i1 < 10; i1++ {
		for i2 := 5; i2 < 10; i2++ {
			for i3 := 5; i3 < 10; i3++ {
				for i4 := 5; i4 < 10; i4++ {
					for i5 := 5; i5 < 10; i5++ {
						tmp := [5]int{i1, i2, i3, i4, i5}
						if checkValid(tmp, 5) {
							pss = append(pss, [5]int{i1, i2, i3, i4, i5})
						}
					}
				}
			}
		}
	}
	return
}

func task1() int {
	content, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	contentString := strings.Split(string(content), ",")
	intcode := make([]int, len(contentString))
	intcodeCopy := make([]int, len(contentString))
	for pos, elem := range contentString {
		intcode[pos] = utils.ToInt(elem)
	}
	copy(intcodeCopy, intcode)
	pss := generatePSS()
	maxThruster := -math.MaxInt32
	var pssMax [5]int
	var inputs [5]inputWrap
	for _, code := range pss {
		for i, c := range code {
			inputs[i] = inputWrap{
				input: []int{c},
			}
		}
		inputs[0].input = append(inputs[0].input, 0)
		for ampNr := 0; ampNr < 5; ampNr++ {
			outputs := processIntCode(intcode, &inputs[ampNr])
			if ampNr < 4 {
				inputs[ampNr+1].input = append(inputs[ampNr+1].input, outputs[0])
			} else {
				if outputs[0] > maxThruster {
					maxThruster = outputs[0]
					pssMax = code
				}
			}

		}
	}
	fmt.Println(pssMax)
	return maxThruster
}

func task2() int {
	content, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	contentString := strings.Split(string(content), ",")
	intcode := make([]int, len(contentString))
	var ampIntcodes [5][]int
	for pos, elem := range contentString {
		intcode[pos] = utils.ToInt(elem)
	}
	for i := 0; i < 5; i++ {
		ampIntcodes[i] = make([]int, len(contentString))
		copy(ampIntcodes[i], intcode)
	}
	pss := generatePSS2()
	maxThruster := -math.MaxInt32
	var pssMax [5]int
	var inputs [5]inputWrap
	indexAmp := [5]int{0, 0, 0, 0}
	var lastOutput int
	for _, code := range pss {
		indexAmp = [5]int{0, 0, 0, 0}
		for i := 0; i < 5; i++ {
			copy(ampIntcodes[i], intcode)
		}
		for i, c := range code {
			inputs[i] = inputWrap{
				input: []int{c},
			}
		}
		inputs[0].input = append(inputs[0].input, 0)
		running := true
		for running {
			for ampNr := 0; ampNr < 5; ampNr++ {
				fmt.Println(inputs)
				for true {
					opCode, paramMode := parseOpcode(splitInt(ampIntcodes[ampNr][indexAmp[ampNr]]))
					//if opCode == 3 && len(inputs[ampNr].input) == 0 {
					//	break
					//}
					if opCode == 99 {
						fmt.Println("Program Finished", ampNr)
						running = false
						break
					}
					ret := compute(opCode, ampIntcodes[ampNr][indexAmp[ampNr]+1:], paramMode, ampIntcodes[ampNr], &indexAmp[ampNr], &inputs[ampNr])

					if ret != -1 {
						if ampNr < 4 {
							lastOutput = ret
							inputs[ampNr+1].input = append(inputs[ampNr+1].input, ret)
							break
						} else {
							lastOutput = ret
							inputs[0].input = append(inputs[0].input, ret)
							break
						}
					}
					//fmt.Println("index: ", index)
				}

			}
		}
		if lastOutput > maxThruster {
			maxThruster = lastOutput
			pssMax = code
		}
	}

	fmt.Println(pssMax)
	return maxThruster
}

func main() {
	fmt.Println("Task 7.1: ", task1())
	fmt.Println("Task 7.2: ", task2())
}
