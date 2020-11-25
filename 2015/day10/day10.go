package main

import "fmt"

func main() {
	start := []int{1, 3, 2, 1, 1, 3, 1, 1, 1, 2}
	//start = []int{2, 1}
	for i := 0; i < 40; i++ {
		current := start[0]
		counter := 1
		nextStep := make([]int, 0)
		for j := 1; j < len(start); j++ {
			if current != start[j] {
				nextStep = append(nextStep, []int{counter, current}...)
				counter = 1
				current = start[j]
			} else {
				counter++
			}
			if j == len(start)-1 {
				nextStep = append(nextStep, []int{counter, current}...)
			}
		}
		start = nextStep
	}
	fmt.Println("Task 10.1:", len(start))
}
