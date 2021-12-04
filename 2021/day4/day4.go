package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type board struct {
	numbers [5][5]int
	marked  [5][5]int
}

func markNumber(b *board, number int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.numbers[i][j] == number {
				b.marked[i][j] = 1
			}
		}
	}
}

func checkWinner(b board) bool {
	for i := 0; i < 5; i++ {
		row, col := 0, 0
		for j := 0; j < 5; j++ {
			row += b.marked[i][j]
			col += b.marked[j][i]
		}
		if row == 5 || col == 5 {
			return true
		}
	}
	return false
}

func boardScore(b board) int {
	score := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.marked[i][j] == 0 {
				score += b.numbers[i][j]
			}
		}
	}
	return score
}

func main() {
	file, err := os.Open("2021/day4/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	boards := make([]board, 0)
	numbers := make([]int, 0)
	scanner.Scan()
	firstLine := strings.Split(scanner.Text(), ",")
	for _, s := range firstLine {
		numbers = append(numbers, utils.ToInt(s))
	}

	for scanner.Scan() {
		scanner.Text()
		newBoard := board{}
		for i := 0; i < 5; i++ {
			scanner.Scan()
			line := strings.Fields(scanner.Text())
			for i2, s := range line {
				newBoard.numbers[i][i2] = utils.ToInt(s)
			}
		}
		boards = append(boards, newBoard)
	}
	numBoards := len(boards)
	winners := make([]int, numBoards)
	taskOneNotDone := true
outer:
	for _, number := range numbers {
		for bi := range boards {
			markNumber(&boards[bi], number)
		}
		for bi, b := range boards {

			if checkWinner(b) {
				//fmt.Println("Winner", bi, "num", number)
				if taskOneNotDone {
					fmt.Println("Day 4.1:", boardScore(b)*number)
					taskOneNotDone = false
				}
				winners[bi] = 1
				winnerSum := 0
				for _, winner := range winners {
					winnerSum += winner
				}
				if winnerSum == numBoards {
					fmt.Println("Day 4.2:", boardScore(b)*number)
					break outer
				}
			}
		}

	}
}
