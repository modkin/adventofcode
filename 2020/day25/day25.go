package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
)

func calcHash(subjectNumber int, loopSize int) int {
	current := 1
	for i := 0; i < loopSize; i++ {
		current *= subjectNumber
		current = current % 20201227
	}
	return current
}

func main() {
	scanner := bufio.NewScanner(utils.OpenFile("2020/day25/input"))
	scanner.Scan()
	cardPublic := utils.ToInt(scanner.Text())
	scanner.Scan()
	doorPublic := utils.ToInt(scanner.Text())
	cardLoopSize, doorLoopSize := 0, 0

	testLoopSize := 1
	hash := 1
	for cardLoopSize == 0 || doorLoopSize == 0 {
		hash *= 7
		hash = hash % 20201227
		//hash := calcHash(7, testLoopSize)
		if hash == cardPublic {
			cardLoopSize = testLoopSize
		}
		if hash == doorPublic {
			doorLoopSize = testLoopSize
		}
		testLoopSize++
	}
	fmt.Println(cardLoopSize, doorLoopSize)
	encyrptionKey := calcHash(doorPublic, cardLoopSize)
	fmt.Println(encyrptionKey)
}
