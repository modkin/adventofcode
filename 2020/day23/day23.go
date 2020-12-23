package main

import (
	"adventofcode/utils"
	"container/ring"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type currentCup struct {
	pos   int
	value int
}

func getDestination(currentCup int, cups *ring.Ring) int {
	sorted := make([]int, 0)
	for i := 0; i < cups.Len(); i++ {
		sorted = append(sorted, cups.Value.(int))
		cups = cups.Next()
	}
	sort.Ints(sorted)
	for i, elem := range sorted {
		if elem == currentCup {
			if i == 0 {
				return sorted[len(sorted)-1]
			} else {
				return sorted[i-1]
			}
		}
	}
	fmt.Println("ERROR")
	return math.MaxInt64
}

func printRing(ring *ring.Ring) {
	for i := 0; i < ring.Len(); i++ {
		fmt.Print(ring.Value.(int), " ")
		ring = ring.Next()
	}
	fmt.Println()
}

func main() {
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
	for i := 0; i < 100; i++ {

		cups = cups.Next()
		//printRing(cups)
		currentVal := cups.Value.(int)
		//fmt.Println(currentVal)
		pickup := cups.Unlink(3)
		//printRing(cups)
		destinationCup := getDestination(currentVal, cups)
		for cups.Value.(int) != destinationCup {
			cups = cups.Next()
		}
		cups.Link(pickup)
		for cups.Value.(int) != currentVal {
			cups = cups.Next()
		}
	}
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
	fmt.Println(solution)
}
