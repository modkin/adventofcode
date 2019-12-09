package main

import (
	"adventofcode/2019/computer"
	"adventofcode/utils"
	"fmt"
	"strings"
	"testing"
)

func TestTask1(t *testing.T) {
	test1 := "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"
	test2 := "1102,34915192,34915192,7,4,7,99,0"
	//test3 := "104,1125899906842624,99"

	test1check := []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}

	inputCh := make(chan int64)
	outputCh := make(chan int64, len(test1check))

	var output []int64

	test1Split := strings.Split(test1, ",")
	intcode1 := make([]int64, len(test1Split))
	for pos, elem := range test1Split {
		intcode1[pos] = utils.ToInt64(elem)
	}

	go computer.ProcessIntCode(intcode1, inputCh, outputCh)

	//inputCh <- 1
	//loop til channel is closed
	for _, _ := range test1check {
		output = append(output, <-outputCh)
	}

	for i, _ := range intcode1 {
		if test1check[i] != output[i] {
			t.Errorf("Wrong!!")
		}
	}
	output = make([]int64, 0)
	test2Split := strings.Split(test2, ",")
	intcode2 := make([]int64, len(test2Split))
	for pos, elem := range test2Split {
		intcode2[pos] = utils.ToInt64(elem)
	}
	go computer.ProcessIntCode(intcode2, inputCh, outputCh)
	for out := range outputCh {
		output = append(output, out)
	}
	if len(fmt.Sprint(output[0])) != 16 {
		t.Error("Wrong: ", output[0])
	}
}
