package main

import (
	"adventofcode/2019/computer"
	"adventofcode/utils"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
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
	quit := make(chan bool)

	go computer.ProcessIntCode(intcode, inputCh, outputCh, quit)

	reader := bufio.NewReader(os.Stdin)

	///you need cake boulder antenna coin
	for true {
		time.Sleep(10000000)
		select {
		case input := <-outputCh:
			fmt.Print(string(input))
		default:
			text, _ := reader.ReadString('\n')
			toRune := []rune(text)
			for _, elem := range toRune {
				inputCh <- int64(elem)
			}
		}
	}
}
