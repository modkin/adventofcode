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
	inputCh := make(chan int64, 1)
	outputCh := make(chan int64, 10)

	quit := make(chan bool, 1)
	inputCh <- 1
	computer.ProcessIntCode(intcode, inputCh, outputCh, quit)

	fmt.Println("Task 7.1: ", <-outputCh)

	intcode = make([]int64, len(contentString))
	for pos, elem := range contentString {
		intcode[pos] = utils.ToInt64(elem)
	}
	inputCh <- 2
	computer.ProcessIntCode(intcode, inputCh, outputCh, quit)

	fmt.Println("Task 7.2: ", <-outputCh)
}
