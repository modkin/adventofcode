package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(utils.OpenFile("2020/day22/input"))
	deck1 := make([]int, 0)
	deck2 := make([]int, 0)
	player1 := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "Player 1:" || line == "" {
			continue
		} else if line == "Player 2:" {
			player1 = false
			continue
		}
		if player1 {
			deck1 = append(deck1, utils.ToInt(line))
		} else {
			deck2 = append(deck2, utils.ToInt(line))
		}
	}
	fmt.Println(deck1)
	fmt.Println(deck2)

	for len(deck1) > 0 && len(deck2) > 0 {
		if deck1[0] > deck2[0] {
			deck1 = append(deck1[1:], deck1[0], deck2[0])
			deck2 = deck2[1:]
		} else {
			deck2 = append(deck2[1:], deck2[0], deck1[0])
			deck1 = deck1[1:]
		}
	}
	fmt.Println(deck1)
	fmt.Println(deck2)
	score := 0
	for i, elem := range deck1 {
		score += (len(deck1) - i) * elem
	}
	fmt.Println(score)
}
