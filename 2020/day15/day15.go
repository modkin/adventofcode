package main

import "fmt"

func playRounds(input []int, maxRounds int) int {
	spoken := make(map[int][]int)
	round, lastSpoken := 1, 0
	for _, number := range input {
		spoken[number] = []int{round}
		lastSpoken = number
		round++
	}

	speak := func(num int, r int) {
		if tmp, ok := spoken[num]; ok {
			spoken[num] = []int{r, tmp[0]}
		} else {
			spoken[num] = []int{r}
		}
	}
	for ; round <= maxRounds; round++ {
		if len(spoken[lastSpoken]) == 1 {
			speak(0, round)
			lastSpoken = 0
		} else {
			toSpeak := spoken[lastSpoken][0] - spoken[lastSpoken][1]
			speak(toSpeak, round)
			lastSpoken = toSpeak
		}

	}
	return lastSpoken
}

func main() {
	input := []int{9, 3, 1, 0, 8, 4}
	//input = []int{0,3,6}

	fmt.Println("Task 15.1:", playRounds(input, 2020))
	fmt.Println("Task 15.2:", playRounds(input, 30000000))
}
