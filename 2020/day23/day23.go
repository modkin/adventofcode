package main

import (
	"adventofcode/utils"
	"container/ring"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func printRing(ring *ring.Ring) {
	for i := 0; i < ring.Len(); i++ {
		fmt.Print(ring.Value.(int), " ")
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

func playGame(cups *ring.Ring, rounds int) {
	cups.Next()
	previousPtr := cups
	currentSafe := cups
	dstPtr := cups
	cups.Prev()

	start := time.Now()
	totalCups := cups.Len()
	pickedUp := make([]int, 3)
	searchDest := 0
	searchNext := 0
	for i := 0; i < rounds; i++ {

		cups = cups.Next()
		currentVal := cups.Value.(int)
		pickup := cups.Unlink(3)
		//destinationCup := getOptimizedDestination(currentVal, pickup, )
		destinationCup := 0

		for k := range pickedUp {
			pickedUp[k] = pickup.Value.(int)
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
		currentSafe = cups
		//currentSafeAdr := &cups
		//cups = cups.Move(-i)
		//cups = previousPtr
		for dstPtr.Value.(int) != destinationCup {
			dstPtr = dstPtr.Next()
			searchDest++
		}
		previousPtr = cups

		dstPtr.Link(pickup)
		dstPtr = dstPtr.Next().Next().Next()
		if destinationCup%16 == 12 {
			dstPtr = dstPtr.Next().Next().Next().Next()
		}
		//cups = currentSafe
		//cups2 := *currentSafeAdr
		if false {
			fmt.Println(currentSafe)
			fmt.Println(previousPtr)
		}
		for cups.Value.(int) != currentVal {
			cups = cups.Next()
			searchNext++
		}
		if searchDest > 0 || searchNext > 0 {

			fmt.Println(i, searchDest, searchNext, time.Now().Sub(start), destinationCup)
			start = time.Now()
			searchDest = 0
			searchNext = 0
		}
	}
}

func solve() {
	input := "952316487"
	//input = "389125467"
	cupsStr := strings.Split(input, "")
	cupsInt := make([]int, len(cupsStr))
	for i, val := range cupsStr {
		cupsInt[i] = utils.ToInt(val)
	}

	cups := ring.New(len(cupsInt))
	for _, val := range cupsInt {
		cups.Value = val
		cups = cups.Next()
	}
	cups = cups.Prev()
	//printRing(cups)
	//playGame(cups, 100)
	printRing(cups)
	for cups.Value.(int) != 1 {
		cups = cups.Next()
	}
	cups = cups.Next()
	solution := ""
	for i := 0; i < cups.Len()-1; i++ {
		solution += strconv.Itoa(cups.Value.(int))
		cups = cups.Next()
	}
	fmt.Println("Task 23.1:", solution)
	//playIntGame(cupsInt, 100)
	//fmt.Println(cupsInt)

	cups2RingPtr := make([]*ring.Ring, 1000*1000)
	cups2 := make([]int, 1000000)
	cups2Ring := ring.New(1000000)
	for i := 0; i < 1000000; i++ {
		if i < len(cupsInt) {
			cups2[i] = cupsInt[i]
			cups2Ring.Value = cupsInt[i]
		} else {
			cups2[i] = i + 1
			cups2Ring.Value = i + 1
		}
		cups2RingPtr[i] = cups2Ring
		cups2Ring = cups2Ring.Next()
	}
	cups2Ring = cups2Ring.Prev()
	start := time.Now()
	playGame(cups2Ring, 100)
	//playIntGame(cups2, 1000*1000)
	fmt.Println(time.Now().Sub(start))
	for cups2Ring.Value.(int) != 1 {
		cups2Ring = cups2Ring.Next()
	}
	first := cups2Ring.Next().Value.(int)
	second := cups2Ring.Next().Next().Value.(int)
	fmt.Println(first * second)
}

func main() {
	solve()
}
