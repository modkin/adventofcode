package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func hashGame(deck1 []int, deck2 []int) string {
	ret := make([]string, 0)
	for _, card := range deck1 {
		ret = append(ret, strconv.Itoa(card))
	}
	ret = append(ret, "x")
	for _, card := range deck2 {
		ret = append(ret, strconv.Itoa(card))
	}
	return strings.Join(ret, "")
}

func playGame(deck1 []int, deck2 []int, playRecursive bool) (bool, []int) {
	//fmt.Println(len(allGames))
	alreadyPlayed := make(map[string]bool)

	playerOneWon := false
	for len(deck1) > 0 && len(deck2) > 0 {
		playerOneWon = false
		if deck1[0] <= len(deck1[1:]) && deck2[0] <= len(deck2[1:]) && playRecursive {
			playerOneWon, _ = playGame(utils.CopyIntSlice(deck1[1:deck1[0]+1]), utils.CopyIntSlice(deck2[1:deck2[0]+1]), playRecursive)
		} else {
			if deck1[0] > deck2[0] {
				playerOneWon = true
			}
		}
		if playerOneWon {
			deck1 = append(deck1[1:], deck1[0], deck2[0])
			deck2 = deck2[1:]
		} else {
			deck2 = append(deck2[1:], deck2[0], deck1[0])
			deck1 = deck1[1:]
		}
		hash := hashGame(deck1, deck2)
		if _, ok := alreadyPlayed[hash]; ok {
			return true, deck1
		} else {
			alreadyPlayed[hash] = true
		}
	}
	if len(deck1) > len(deck2) {
		return true, deck1
	} else {
		return false, deck2
	}

}

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
	_, winningDeck := playGame(deck1, deck2, false)
	score := 0
	for i, elem := range winningDeck {
		score += (len(winningDeck) - i) * elem
	}
	fmt.Println("Task 22.1", score)

	_, winningDeck = playGame(deck1, deck2, true)
	score = 0
	for i, elem := range winningDeck {
		score += (len(winningDeck) - i) * elem
	}
	fmt.Println("Task 22.2", score)
}
