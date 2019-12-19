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

	//shipMap := make(map[[2]int]string)

	//var beam [200][200]int

	counter := 0
	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			inputCh := make(chan int64)
			outputCh := make(chan int64)
			quit := make(chan bool)
			go computer.ProcessIntCode(intcode, inputCh, outputCh, quit)
			inputCh <- int64(x)
			inputCh <- int64(y)
			out := <-outputCh
			if out == 1 {
				counter++
			}

		}
	}
	fmt.Println(counter)

}
