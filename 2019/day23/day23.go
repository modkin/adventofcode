package main

import (
	"adventofcode/2019/computer"
	"adventofcode/utils"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type NIC struct {
	input  chan int64
	output chan int64
	quit   chan bool
}

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

	var network [50]NIC

	for x := 0; x < 50; x++ {
		newNIC := NIC{
			input:  make(chan int64),
			output: make(chan int64),
			quit:   make(chan bool),
		}
		network[x] = newNIC
	}
	for x := 0; x < 50; x++ {
		go computer.ProcessIntCode(intcode, network[x].input, network[x].output, network[x].quit)
		network[x].input <- int64(x)
	}

	var messages [50][][2]int64

	running := true
	for running {
		for i := 0; i < 50; i++ {
			time.Sleep(100000)
			select {
			case address := <-network[i].output:
				x := <-network[i].output
				y := <-network[i].output
				if address == 255 {
					fmt.Println(y)
					running = false
					break
				}
				newMessage := [2]int64{x, y}
				messages[address] = append(messages[address], newMessage)
			default:
				if len(messages[i]) > 0 {
					message := messages[i][0]
					messages[i] = messages[i][1:]
					network[i].input <- message[0]
					network[i].input <- message[1]
				} else {
					network[i].input <- -1
				}
			case <-network[i].quit:
				fmt.Println("is this supposed to happen?")
				running = false
			}
		}
	}
}
