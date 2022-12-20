package main

import (
	"adventofcode/utils"
	"bufio"
	"container/ring"
	"fmt"
	"math"
	"os"
)

func searchIdx(list []int, tar int) int {
	for i, i2 := range list {
		if i2 == tar {
			return i
		}
	}
	return math.MaxInt
}

func printRing(r *ring.Ring, ints []int) {
	n := r.Len()

	fmt.Println()
	for j := 0; j < n; j++ {
		fmt.Print(ints[r.Value.(int)])
		fmt.Print(" ")
		r = r.Next()
	}
	fmt.Println()
}

func main() {
	file, err := os.Open("2022/day20/input")

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	rotations := 1
	initlist := make([]int, 0)

	for scanner.Scan() {
		initlist = append(initlist, utils.ToInt(scanner.Text()))
	}
	myRing := ring.New(len(initlist))

	for i, _ := range initlist {
		myRing.Value = i
		myRing = myRing.Next()
	}
	//printRing(myRing, initlist)

	doRotation := func() {
		for i := 0; i < len(initlist); i++ {
			rotation := initlist[i]
			if rotation == 0 {
				//fmt.Println(value)
				continue
			}
			//fmt.Println(value)
			for myRing.Next().Value != i {
				myRing = myRing.Next()
			}

			ringEle := myRing.Unlink(1)
			myRing = myRing.Move(rotation)
			myRing.Link(ringEle)

			///printRing(myRing, initlist)
		}
	}
	for i := 0; i < rotations; i++ {
		doRotation()
	}
	for initlist[myRing.Value.(int)] != 0 {
		myRing = myRing.Next()
	}
	sum := 0
	for _, i := range []int{1000, 1000, 1000} {
		myRing = myRing.Move(i)
		sum += initlist[myRing.Value.(int)]
	}
	fmt.Println(sum)
}
