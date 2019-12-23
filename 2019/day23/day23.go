package main

import (
	"adventofcode/2019/computer"
	"adventofcode/utils"
	"fmt"
	"io/ioutil"
	"math"
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
	var nat [2]int64

	firstY := true
	lastY := int64(math.MaxInt64)
	running := true
	idleCounter := 0
	for running {
		allIdle := true
		for i := 0; i < 50; i++ {
			time.Sleep(10000)
			select {
			case address := <-network[i].output:
				allIdle = false
				x := <-network[i].output
				y := <-network[i].output
				if address == 255 {
					nat[0] = x
					nat[1] = y
					if firstY {
						fmt.Println("Task 23.1: ", y)
						firstY = false
					}
				} else {
					newMessage := [2]int64{x, y}
					messages[address] = append(messages[address], newMessage)
				}
			default:
				if len(messages[i]) > 0 {
					message := messages[i][0]
					messages[i] = messages[i][1:]
					network[i].input <- message[0]
					network[i].input <- message[1]
					allIdle = false
				} else {
					network[i].input <- -1
				}
			case <-network[i].quit:
				fmt.Println("is this supposed to happen?")
				running = false
			}
		}
		if idleCounter > 100 {
			if nat[1] == lastY {
				fmt.Println("Task 23.2: ", nat[1])
				running = false
				break
			}
			lastY = nat[1]
			network[0].input <- nat[0]
			network[0].input <- nat[1]
			idleCounter = 0
		}
		if allIdle {
			idleCounter++
		} else {
			idleCounter = 0
		}
	}
}
