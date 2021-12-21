package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func roll(dice int) (int, int) {
	points := 0
	for i := 0; i < 3; i++ {
		points += dice
		dice++
		if dice > 100 {
			dice = dice % 100
		}
	}
	return points, dice
}

func main() {
	file, err := os.Open("2021/day21/input")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")
	player1 := utils.ToInt(line[len(line)-1])
	scanner.Scan()
	line = strings.Split(scanner.Text(), " ")
	player2 := utils.ToInt(line[len(line)-1])
	fmt.Println(player1, player2)
	player1Points, player2Points := 0, 0

	dice := 1
	counter := 0
	loser := 0
	add := 0
	for {
		fmt.Println(dice, player1Points, player2Points)
		add, dice = roll(dice)
		counter += 3
		player1 += add
		if player1 > 10 {
			player1 = player1 % 10
			if player1 == 0 {
				player1 = 10
			}
		}
		player1Points += player1
		if player1Points >= 1000 {
			loser = player2Points
			break
		}

		add, dice = roll(dice)
		counter += 3
		player2 += add
		if player2 > 10 {
			player2 = player2 % 10
			if player2 == 0 {
				player2 = 10
			}
		}

		player2Points += player2
		if player2Points >= 1000 {
			loser = player1Points
			break
		}

	}

	fmt.Println(loser, counter)
	fmt.Println(loser * counter)
}
