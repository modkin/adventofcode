package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

// best == highest
func handPoints(hand []string) (int, int) {
	handType := 0
	var intHand []int
	for _, s := range hand {
		intHand = append(intHand, toValue(s))
	}
	sort.Ints(intHand)
	if strings.Count(strings.Join(hand, ""), string(hand[0])) == 5 {
		handType = 7
	} else if strings.Count(strings.Join(hand, ""), string(hand[0])) == 4 || strings.Count(strings.Join(hand, ""), string(hand[1])) == 4 {
		handType = 6
	} else if intHand[0] == intHand[1] && intHand[2] == intHand[3] && intHand[3] == intHand[4] {
		handType = 5
	} else if intHand[0] == intHand[1] && intHand[1] == intHand[2] && intHand[3] == intHand[4] {
		handType = 5
	} else if intHand[0] == intHand[1] && intHand[1] == intHand[2] {
		handType = 4
	} else if intHand[1] == intHand[2] && intHand[2] == intHand[3] {
		handType = 4
	} else if intHand[2] == intHand[3] && intHand[3] == intHand[4] {
		handType = 4
	} else if intHand[0] == intHand[1] && intHand[2] == intHand[3] {
		handType = 3
	} else if intHand[1] == intHand[2] && intHand[3] == intHand[4] {
		handType = 3
	} else if intHand[0] == intHand[1] && intHand[3] == intHand[4] {
		handType = 3
	} else if intHand[0] == intHand[1] || intHand[1] == intHand[2] || intHand[2] == intHand[3] || intHand[3] == intHand[4] {
		handType = 2
	} else {
		handType = 1
	}
	return handType, toValue(hand[0])*100000000 + toValue(hand[1])*1000000 + toValue(hand[2])*10000 + toValue(hand[3])*100 + toValue(hand[4])
}

func findBestJoker(hand []string) int {
	cards := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	maxPoints := 0
	for _, card := range cards {
		joined := strings.Join(hand, "")
		replaced := strings.ReplaceAll(joined, "J", card)
		handType, _ := handPoints(strings.Split(replaced, ""))
		if handType > maxPoints {
			maxPoints = handType
		}
	}
	return maxPoints
}

func toValue(in string) int {
	order := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	order = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}
	for i, i2 := range order {
		if i2 == in {
			return i + 2
		}
	}
	return math.MaxInt
}

func main() {
	file, err := os.Open("2023/day7/input")
	if err != nil {
		panic(err)
	}

	//order := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

	scanner := bufio.NewScanner(file)

	//races := make([][]int, 0)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}

	//bets := make(map[string]int)
	var bets []int
	var ranks []int
	var ref []int
	var hands []string
	for _, line := range lines {
		split := strings.Fields(line)
		//bets[split[0]] = utils.ToInt(split[1])
		bets = append(bets, utils.ToInt(split[1]))
		hands = append(hands, split[0])
		points, hihgest := handPoints(strings.Split(split[0], ""))
		points = findBestJoker(strings.Split(split[0], ""))
		ref = append(ref, 10000000000*points+hihgest)
		ranks = append(ranks, 10000000000*points+hihgest)
	}

	sort.Ints(ranks)
	totalWinning := 0
	rankSum := 0
	for rank, value := range ranks {
		for betIdx, value2 := range ref {
			if value2 == value {
				totalWinning += (rank + 1) * bets[betIdx]
				rankSum += (betIdx + 1)
				fmt.Println(hands[betIdx], value, rank+1, bets[betIdx])
			}
		}
	}
	fmt.Println(len(ranks))
	fmt.Println(len(bets))
	fmt.Println(rankSum)

	fmt.Println("Day 7.1:", totalWinning)
}
