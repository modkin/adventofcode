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
	inputCh := make(chan int64)
	outputCh := make(chan int64)

	var output []int64

	go computer.ProcessIntCode(intcode, inputCh, outputCh)

	inputCh <- 1
	//loop til channel is closed
	for out := range outputCh {
		output = append(output, out)
	}
	fmt.Println("Task 7.1: ", output)
	//fmt.Println("Task 7.2: ", task2(intcode))
}
