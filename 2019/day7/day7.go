package main

import (
	"adventofcode/utils"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

/// paraMode assumed to be filled with 0s
func compute(opcode int, paramMode []int, memory []int, itrPtr *int, input <-chan int, output chan<- int) {
	param := memory[*itrPtr+1:]

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
		memory[param[0]] = <-input
		*itrPtr += 2
	case 4:
		if paramMode[0] == 0 {
			output <- memory[param[0]]
		} else if paramMode[0] == 1 {
			output <- param[0]
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

func processIntCode(intcode []int, input <-chan int, output chan<- int, halt chan<- bool) {
	ownIntcode := make([]int, len(intcode))
	copy(ownIntcode, intcode)
	index := 0
	for true {
		opCode, paramMode := parseOpcode(utils.SplitInt(ownIntcode[index]))
		if opCode == 99 {
			halt <- true
			return
		}
		compute(opCode, paramMode, ownIntcode, &index, input, output)
	}
	return
}

func heapPermutation(input []int) (ouput [][]int) {
	var generate func(int, []int)
	generate = func(k int, work []int) {
		if k == 0 {
			newWork := append(work[:0:0], work...)
			ouput = append(ouput, newWork)
			return
		}
		for i := 0; i < len(input); i++ {
			generate(k-1, work)
			if k%2 == 0 {
				work[0], work[k-1] = work[k-1], work[0]
			} else {
				work[i], work[k-1] = work[k-1], work[i]
			}
		}
	}
	generate(len(input), input)
	return
}

func task1(intcode []int) int {
	maxThruster := -math.MaxInt32
	var pssMax []int
	channels := [...]chan int{make(chan int), make(chan int), make(chan int), make(chan int), make(chan int, 1)}

	for _, code := range heapPermutation([]int{0, 1, 2, 3, 4}) {
		halt := make(chan bool)
		for ampNr := 0; ampNr < 5; ampNr++ {
			go processIntCode(intcode, channels[(ampNr+4)%5], channels[ampNr], halt)
		}
		for i, c := range code {
			channels[i] <- c
		}
		channels[4] <- 0
		for i := 0; i < 5; i++ {
			<-halt
		}
		output := <-channels[4]
		if output > maxThruster {
			maxThruster = output
			pssMax = code
		}
	}
	fmt.Println(pssMax)
	return maxThruster
}

func task2(intcode []int) int {
	var ampIntcodes [5][]int
	for i := 0; i < 5; i++ {
		ampIntcodes[i] = make([]int, len(intcode))
		copy(ampIntcodes[i], intcode)
	}
	maxThruster := -math.MaxInt32
	///0: A->B, 1: B->C, 3: C->D, 4: D->E, 5: E->A
	channels := [...]chan int{make(chan int, 2), make(chan int, 2), make(chan int, 2), make(chan int, 2), make(chan int, 2)}
	var lastOutput int
	for _, code := range heapPermutation([]int{5, 6, 7, 8, 9}) {
		halt := make(chan bool)
		for i := 0; i < 5; i++ {
			copy(ampIntcodes[i], intcode)
		}
		for i, c := range code {
			channels[i] <- c
		}
		channels[4] <- 0
		for ampNr := 0; ampNr < 5; ampNr++ {
			go processIntCode(ampIntcodes[ampNr], channels[(ampNr+4)%5], channels[ampNr], halt)
		}
		for i := 0; i < 5; i++ {
			<-halt
		}
		lastOutput = <-channels[4]
		if lastOutput > maxThruster {
			maxThruster = lastOutput
		}
	}
	return maxThruster
}

func main() {
	content, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	contentString := strings.Split(string(content), ",")
	intcode := make([]int, len(contentString))
	for pos, elem := range contentString {
		intcode[pos] = utils.ToInt(elem)
	}
	fmt.Println("Task 7.1: ", task1(intcode))
	fmt.Println("Task 7.2: ", task2(intcode))
}
