package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func inRange(first int, second int, check int) bool {
	for i := first; i <= second; i++ {
		if check == i {
			return true
		}
	}
	return false
}

func main() {

	file, err := os.Open("2022/day4/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	rangePairs := make([][2][2]int, 0)

	for scanner.Scan() {
		re := regexp.MustCompile(`\d*`)
		dgts := re.FindAllString(scanner.Text(), -1)
		fp := [2]int{utils.ToInt(dgts[0]), utils.ToInt(dgts[1])}
		sp := [2]int{utils.ToInt(dgts[2]), utils.ToInt(dgts[3])}
		rangePairs = append(rangePairs, [2][2]int{fp, sp})
	}

	counter := 0
	for _, pair := range rangePairs {
		if pair[1][0] >= pair[0][0] && pair[1][1] <= pair[0][1] {
			counter++
		} else if pair[0][0] >= pair[1][0] && pair[0][1] <= pair[1][1] {
			counter++
		}
	}

	fmt.Println("Day 3.1:", counter)

	counter2 := 0
outer:
	for _, pair := range rangePairs {
		for i := pair[0][0]; i <= pair[0][1]; i++ {
			if inRange(pair[1][0], pair[1][1], i) {
				counter2++
				continue outer
			}
		}
	}
	fmt.Println("Day 3.1:", counter2)

}
