package main

import (
	"adventofcode/utils"
	"container/ring"
	"fmt"
	"strings"
	"time"
)

func printRing(ring *ring.Ring) {
	for i := 0; i < ring.Len(); i++ {
		fmt.Print(ring.Value.(cup).value, " ")
		ring = ring.Next()
	}
	fmt.Println()
}

func playIntGame(cups []int, rounds int) {
	amountOfCups := len(cups)
	start := time.Now()
	searchDest := 0
	for i := 0; i < rounds; i++ {
		pickup := utils.CopyIntSlice(cups[1:4])
		cups = append(cups[:1], cups[4:]...)
		destinationCup := 0
		tmp := cups[0]
		for {
			tmp--
			if tmp == 0 {
				tmp = amountOfCups
			}
			if !utils.IntSliceContains(pickup, tmp) {
				destinationCup = tmp
				break
			}
		}
		destinationIdx := 0
		for cups[destinationIdx] != destinationCup {
			destinationIdx++
			searchDest++
		}

		tmpCups := append(cups[0:destinationIdx], pickup...)
		cups = append(tmpCups, cups[destinationIdx:]...)
		cups = append(cups[1:], cups[len(cups)-1])
		if i%100 == 0 {
			fmt.Println(i, searchDest, time.Now().Sub(start), destinationCup)
			start = time.Now()
			searchDest = 0
		}
	}
}

func playGame(cups *ring.Ring, rounds int, output bool) {
	dstPtr := cups.Next()

	start := time.Now()
	totalCups := cups.Len()
	pickedUp := make([]int, 3)
	searchDest := 0
	searchNext := 0
	for i := 0; i < rounds; i++ {

		cups = cups.Next()
		currentVal := cups.Value.(cup).value
		dstPtr = cups.Value.(cup).previousVal
		for dstPtr == cups.Next() || dstPtr == cups.Next().Next() || dstPtr == cups.Next().Next().Next() {
			dstPtr = dstPtr.Value.(cup).previousVal
		}
		pickup := cups.Unlink(3)
		destinationCup := 0

		for k := range pickedUp {
			pickedUp[k] = pickup.Value.(cup).value
			pickup = pickup.Next()
		}
		current := currentVal
		for {
			current--
			if current == 0 {
				current = totalCups
			}
			if !utils.IntSliceContains(pickedUp, current) {
				destinationCup = current
				break
			}
		}
		for dstPtr.Value.(cup).value != destinationCup {
			dstPtr = dstPtr.Next()
			searchDest++
		}

		dstPtr.Link(pickup)
		//dstPtr = dstPtr.Next().Next().Next()
		//if destinationCup%16 == 12 {
		//	dstPtr = dstPtr.Next().Next().Next().Next()
		//}

		for cups.Value.(cup).value != currentVal {
			cups = cups.Next()
			searchNext++
		}
		if output {
			if searchDest > 0 || searchNext > 0 {

				fmt.Println(i, searchDest, searchNext, time.Now().Sub(start), destinationCup)
				start = time.Now()
				searchDest = 0
				searchNext = 0
			}
		}
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
	//playGame(cups2Ring, 100, true)
	//printRing(cups2Ring)
	//for cups2Ring.Value.(cup).value != 1 {
	//	cups2Ring = cups2Ring.Next()
	//}
	//cups2Ring = cups2Ring.Next()
	//solution := ""
	//for i := 0; i < cups2Ring.Len()-1; i++ {
	//	solution += strconv.Itoa(cups2Ring.Value.(cup).value)
	//	cups2Ring = cups2Ring.Next()
	//}
	//fmt.Println("Task 23.1:", solution)
	//if solution != "25398647" {
	//	fmt.Println("error")
	//	os.Exit(1)
	//}
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
	start := time.Now()
	playGame(cups2Ring, 10000000, true)
	fmt.Println(time.Now().Sub(start))
	for cups2Ring.Value.(cup).value != 1 {
		cups2Ring = cups2Ring.Next()
	}
	first := cups2Ring.Next().Value.(cup).value
	second := cups2Ring.Next().Next().Value.(cup).value
	fmt.Println(first, second)
	fmt.Println(first * second)
}

func main() {
	solve()
}
