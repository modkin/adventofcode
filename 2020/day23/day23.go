package main

import (
	"adventofcode/utils"
	"container/ring"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printRing(ring *ring.Ring) {
	for i := 0; i < ring.Len(); i++ {
		fmt.Print(ring.Value.(cup).value, " ")
		ring = ring.Next()
	}
	fmt.Println()
}

func playGame(cups *ring.Ring, rounds int) {
	dstPtr := cups.Next()
	pickedUp := make([]int, 3)
	for i := 0; i < rounds; i++ {

		cups = cups.Next()
		dstPtr = cups.Value.(cup).previousVal
		for dstPtr == cups.Next() || dstPtr == cups.Next().Next() || dstPtr == cups.Next().Next().Next() {
			dstPtr = dstPtr.Value.(cup).previousVal
		}
		pickup := cups.Unlink(3)

		for k := range pickedUp {
			pickedUp[k] = pickup.Value.(cup).value
			pickup = pickup.Next()
		}

		dstPtr.Link(pickup)
	}
}

type cup struct {
	value       int
	previousVal *ring.Ring
}

func solve() {
	input := "952316487"
	//input = "389125467"
	cupsStr := strings.Split(input, "")
	cupsInt := make([]int, len(cupsStr))
	for i, val := range cupsStr {
		cupsInt[i] = utils.ToInt(val)
	}

	cups2RingPtr := make(map[int]*ring.Ring)
	cups2Ring := ring.New(len(cupsInt))
	startCup := cups2Ring
	for _, val := range cupsInt {
		cups2Ring.Value = cup{val, nil}
		cups2RingPtr[val] = cups2Ring
		cups2Ring = cups2Ring.Next()
	}
	tmp := startCup
	for i := 0; i < len(cupsInt); i++ {
		cupNr := tmp.Value.(cup).value
		if cupNr == 1 {
			tmp.Value = cup{cupNr, cups2RingPtr[len(cupsInt)]}
		} else {
			tmp.Value = cup{cupNr, cups2RingPtr[cupNr-1]}
		}
		tmp = tmp.Next()
	}
	cups2Ring = startCup.Prev()
	playGame(cups2Ring, 100)
	//printRing(cups2Ring)
	for cups2Ring.Value.(cup).value != 1 {
		cups2Ring = cups2Ring.Next()
	}
	cups2Ring = cups2Ring.Next()
	solution := ""
	for i := 0; i < cups2Ring.Len()-1; i++ {
		solution += strconv.Itoa(cups2Ring.Value.(cup).value)
		cups2Ring = cups2Ring.Next()
	}
	fmt.Println("Task 23.1:", solution)
	if solution != "25398647" {
		fmt.Println("error")
		os.Exit(1)
	}
	cups2Ring = cups2RingPtr[cupsInt[len(cupsInt)-1]]
	for _, val := range cupsInt {
		cups2Ring.Link(cups2RingPtr[val])
		cups2Ring = cups2Ring.Next()
	}
	cups2Ring = tmp.Prev()
	cups2RingExtend := ring.New(1000000 - len(cupsInt))
	cups2Ring.Link(cups2RingExtend)
	cups2RingPtr[1].Value = cup{1, startCup.Prev()}
	cups2Ring = cups2Ring.Next()
	cups2Ring.Value = cup{len(cupsInt) + 1, cups2RingPtr[len(cupsInt)]}
	cups2Ring = cups2Ring.Next()
	for i := len(cupsInt) + 2; i <= 1000000; i++ {
		cups2Ring.Value = cup{i, cups2Ring.Prev()}
		cups2Ring = cups2Ring.Next()
	}
	cups2Ring = startCup.Prev()
	playGame(cups2Ring, 10000000)
	cups2Ring = cups2RingPtr[1]
	first := cups2Ring.Next().Value.(cup).value
	second := cups2Ring.Next().Next().Value.(cup).value
	fmt.Println("Task 23.2:", first*second)
}

func main() {
	solve()
}
