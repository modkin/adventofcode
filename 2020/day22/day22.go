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
	for card := range deck1 {
		ret = append(ret, strconv.Itoa(card))
	}
	for card := range deck2 {
		ret = append(ret, strconv.Itoa(card))
	}
	return strings.Join(ret, "")
}

func playGame(deck1 []int, deck2 []int, alreadyPlayed map[string]bool) (bool, []int) {
	//hash := hashGame(deck1,deck2)
	//if _,ok := alreadyPlayed[hash]; ok {
	//	return true, deck1
	//} else {
	//	alreadyPlayed[hash] = true
	//}

	playerOneWon := false
	for len(deck1) > 0 && len(deck2) > 0 {
		playerOneWon = false
		if deck1[0] <= len(deck1[1:]) && deck2[0] <= len(deck2[1:]) {
			playerOneWon, _ = playGame(utils.CopyIntSlice(deck1[1:]), utils.CopyIntSlice(deck2[1:]), alreadyPlayed)
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
	}
	if playerOneWon {
		return playerOneWon, deck1
	} else {
		return playerOneWon, deck2
	}

}

func main() {
	scanner := bufio.NewScanner(utils.OpenFile("2020/day22/testinput"))
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

	playedGames := make(map[string]bool)
	_, winningDeck := playGame(deck1, deck2, playedGames)
	fmt.Println(deck1)
	fmt.Println(deck2)
	score := 0
	for i, elem := range winningDeck {
		score += (len(winningDeck) - i) * elem
	}
	fmt.Println("Task 22.1", score)
}
