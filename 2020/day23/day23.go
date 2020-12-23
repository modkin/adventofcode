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

func playGame(cups *ring.Ring, rounds int) {

	searchNext := 0
	start := time.Now()
	pickedUp := make([]int, 3)
	for i := 0; i < rounds; i++ {
		searchDest := 0
		cups = cups.Next()
		currentVal := cups.Value.(int)
		pickup := cups.Unlink(3)
		//destinationCup := getOptimizedDestination(currentVal, pickup, )
		destinationCup := 0

		for i := range pickedUp {
			pickedUp[i] = pickup.Value.(int)
			pickup = pickup.Next()
		}
		current := currentVal
		for {
			current--
			if current == 0 {
				current = cups.Len() + pickup.Len()
			}
			if !utils.IntSliceContains(pickedUp, current) {
				destinationCup = current
				break
			}
		}
		for cups.Value.(int) != destinationCup {
			cups = cups.Prev()
			searchDest++
		}

		cups.Link(pickup)
		for cups.Value.(int) != currentVal {
			cups = cups.Next()
			searchNext++
		}
		if i%50 == 0 {
			fmt.Println(i, searchDest, searchNext, time.Now().Sub(start), destinationCup)
			start = time.Now()
		}
	}
}

func main() {
	input := "952316487"
	input = "389125467"
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

	cups2 := make([]int, 1000000)
	cups2Ring := ring.New(1000000)
	for i := 0; i < len(cups2); i++ {
		if i < len(cupsInt) {
			cups2[i] = cupsInt[i]
			cups2Ring.Value = cupsInt[i]
		} else {
			cups2[i] = i + 1
			cups2Ring.Value = i + 1
		}
		cups2Ring = cups2Ring.Next()
	}
	cups2Ring = cups2Ring.Prev()
	start := time.Now()
	playGame(cups2Ring, 1000*1000)
	fmt.Println(time.Now().Sub(start))
	for cups2Ring.Value.(int) != 1 {
		cups2Ring = cups2Ring.Next()
	}
	first := cups2Ring.Next().Value.(int)
	second := cups2Ring.Next().Next().Value.(int)
	fmt.Println(first * second)
}
