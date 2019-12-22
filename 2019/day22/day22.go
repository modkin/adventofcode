package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cut(n int, deck []int) []int {
	if n > 0 {
		cut := deck[0:n]
		deck = append(deck[n:], cut...)
	} else if n < 0 {
		cut := deck[len(deck)+n:]
		deck = append(cut, deck[0:len(deck)+n]...)
	}
	return deck
}

func dealWithInc(inc int, deck []int) []int {
	newDeck := make([]int, len(deck))
	pos := 0
	for _, elem := range deck {
		newDeck[pos] = elem
		pos += inc
		pos = pos % len(deck)
	}
	return newDeck
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	var deck []int
	for i := 0; i < 10007; i++ {
		deck = append(deck, i)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if line[0] == "cut" {
			deck = cut(utils.ToInt(line[1]), deck)
		} else if line[0] == "deal" && line[1] == "into" {
			deck = utils.ReverseIntSlice(deck)
		} else if line[0] == "deal" && line[1] == "with" {
			deck = dealWithInc(utils.ToInt(line[3]), deck)
		} else {
			fmt.Println("ERROR")
		}
	}
	for i, card := range deck {
		if card == 2019 {
			fmt.Println("Task 22.1: ", i)
		}
	}
}
