package main

import (
	"adventofcode/2019/computer"
	"adventofcode/utils"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"sync"
)

/// paraMode assumed to be filled with 0s

func heapPermutation(input []int) (ouput [][]int) {
	var generate func(int, []int)
	generate = func(k int, work []int) {
		if k == 0 {
			newWork := append(work[:0:0], work...)
			ouput = append(ouput, newWork)
			return
		}
		for i := 0; i < len(input); i++ {
			generate(k-1, work)
			if k%2 == 0 {
				work[0], work[k-1] = work[k-1], work[0]
			} else {
				work[i], work[k-1] = work[k-1], work[i]
			}
		}
	}
	generate(len(input), input)
	return
}

func task1(intcode []int64) int {
	maxThruster := int64(-math.MaxInt64)
	var pssMax []int
	channels := [...]chan int64{make(chan int64), make(chan int64), make(chan int64), make(chan int64), make(chan int64)}

	for _, code := range heapPermutation([]int{0, 1, 2, 3, 4}) {
		for ampNr := 0; ampNr < 5; ampNr++ {
			go computer.ProcessIntCode(intcode, channels[(ampNr+4)%5], channels[ampNr])
		}
		for i, c := range code {
			channels[i] <- int64(c)
		}
		channels[4] <- 0

		output := <-channels[4]
		if output > maxThruster {
			maxThruster = output
			pssMax = code
		}
	}
	fmt.Println(pssMax)
	return int(maxThruster)
}

func task2(intcode []int64) int {
	var ampIntcodes [5][]int64
	for i := 0; i < 5; i++ {
		ampIntcodes[i] = make([]int64, len(intcode))
		copy(ampIntcodes[i], intcode)
	}
	maxThruster := int64(-math.MaxInt64)
	///0: A->B, 1: B->C, 3: C->D, 4: D->E, 5: E->A
	channels := [...]chan int64{make(chan int64), make(chan int64), make(chan int64), make(chan int64), make(chan int64, 1)}
	var lastOutput int64
	for _, code := range heapPermutation([]int{5, 6, 7, 8, 9}) {
		var wg sync.WaitGroup
		for i := 0; i < 5; i++ {
			copy(ampIntcodes[i], intcode)
		}
		for ampNr := 0; ampNr < 5; ampNr++ {
			wg.Add(1)
			go func(ampNr int) {
				computer.ProcessIntCode(ampIntcodes[ampNr], channels[(ampNr+4)%5], channels[ampNr])
				wg.Done()
			}(ampNr)
		}
		for i, c := range code {
			channels[i] <- int64(c)
		}

		channels[4] <- 0
		wg.Wait()
		lastOutput = <-channels[4]
		if lastOutput > maxThruster {
			maxThruster = lastOutput
		}
	}
	return int(maxThruster)
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
	fmt.Println("Task 7.1: ", task1(intcode))
	fmt.Println("Task 7.2: ", task2(intcode))
}
