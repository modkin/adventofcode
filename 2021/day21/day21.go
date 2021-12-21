package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
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

func wrapTrack(pos int) int {
	if pos > 10 {
		pos = pos % 10
		if pos == 0 {
			pos = 10
		}
	}
	return pos
}

func main() {
	playerWins := make(map[[5]int][2]int)
	file, err := os.Open("2021/day21/input")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")
	p1 := utils.ToInt(line[len(line)-1])
	scanner.Scan()
	line = strings.Split(scanner.Text(), " ")
	p2 := utils.ToInt(line[len(line)-1])
	fmt.Println(p1, p2)
	player1Points, player2Points := 0, 0
	player1Start, player2Start := p1, p2
	fmt.Println(player1Start, player2Start)

	dice := 1
	counter := 0
	loser := 0
	add := 0
	for {
		add, dice = roll(dice)
		counter += 3
		p1 += add
		p1 = wrapTrack(p1)
		player1Points += p1
		if player1Points >= 1000 {
			loser = player2Points
			break
		}

		add, dice = roll(dice)
		counter += 3
		p2 += add
		p2 = wrapTrack(p2)

		player2Points += p2
		if player2Points >= 1000 {
			loser = player1Points
			break
		}

	}

	fmt.Println(loser, counter)
	fmt.Println(loser * counter)

	diceList := make([]int, 0)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				diceList = append(diceList, i+j+k)
			}
		}
	}
	sort.Ints(diceList)
	//fmt.Println(diceList)

	for p1points := 20; p1points >= 0; p1points-- {
		for p2points := 20; p2points >= 0; p2points-- {
			for p1 = 1; p1 <= 10; p1++ {
				for p2 = 1; p2 <= 10; p2++ {
					for activePlayer := 1; activePlayer <= 2; activePlayer++ {
						for _, d := range diceList {
							if activePlayer == 1 {
								tmp := playerWins[[5]int{p1, p2, p1points, p2points, activePlayer}]
								if newPoints := wrapTrack(p1 + d); p1points+newPoints >= 21 {
									tmp[0]++
								} else {
									tmp2 := playerWins[[5]int{newPoints, p2, p1points + newPoints, p2points, 2}]
									tmp[0] += tmp2[0]
									tmp[1] = tmp2[1]
								}
								playerWins[[5]int{p1, p2, p1points, p2points, activePlayer}] = tmp
							} else {
								tmp := playerWins[[5]int{p1, p2, p1points, p2points, activePlayer}]
								if newPoints := wrapTrack(p2 + d); p2points+newPoints >= 21 {
									tmp[1]++
								} else {
									tmp2 := playerWins[[5]int{p1, newPoints, p1points, p2points + newPoints, 1}]
									tmp[1] += tmp2[1]
									tmp[0] += tmp2[0]
								}
								playerWins[[5]int{p1, p2, p1points, p2points, activePlayer}] = tmp
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(playerWins[[5]int{10, 10, 20, 17, 2}])
	fmt.Println(playerWins[[5]int{player1Start, player2Start, 0, 0, 1}])

}
