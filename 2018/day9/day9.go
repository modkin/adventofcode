package day9

import "fmt"

func runMarbleGame(players int, lastMarble int) int {
	scores := make([]int, players)
	circle := []int{0, 2, 1}
	currentIdx := 1
	for i := 3; i <= lastMarble; i++ {
		if i%23 != 0 {
			currentIdx = currentIdx + 2
			if currentIdx > len(circle) {
				currentIdx = 1
			}
			circle = append(circle[:currentIdx], append([]int{i}, circle[currentIdx:]...)...)
		} else {
			fmt.Println(i)
			currentIdx = currentIdx - 7
			if currentIdx < 0 {
				currentIdx = len(circle) + currentIdx
			}
			scores[i%players] += i + circle[currentIdx]
			if currentIdx == len(circle)-1 {
				circle = circle[:len(circle)-1]
				currentIdx = 0
			} else {
				circle = append(circle[:currentIdx], circle[currentIdx+1:]...)
			}
		}

		//fmt.Println(circle, currentIdx, circle[currentIdx])
	}
	max := 0
	for _, elem := range scores {
		if elem > max {
			max = elem
		}
	}
	return max
}

func Task1() {
	fmt.Println(runMarbleGame(478, 71240))
}

func Task2() {
	fmt.Println(runMarbleGame(478, 7124000))

}

func Tests() {
	fmt.Println(runMarbleGame(9, 25))
	fmt.Println(runMarbleGame(10, 1618))
	fmt.Println(runMarbleGame(13, 7999))
	fmt.Println(runMarbleGame(17, 1104))
	fmt.Println(runMarbleGame(21, 6111))
}
