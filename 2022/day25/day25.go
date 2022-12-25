package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var snafuMap = map[string]int{"2": 2, "1": 1, "0": 0, "-": -1, "=": -2}

func convertToDecimal(in []string) int {
	ret := 0
	value := 1
	for i := len(in) - 1; i >= 0; i-- {
		ret += snafuMap[in[i]] * value
		value *= 5
	}
	return ret
}

func convertToSnafu(in int) []string {
	value := 1
	counter := 1
	for value < in {
		value *= 5
		counter++
	}
	counter--
	value /= 5
	tmpNumbers := make([]int, counter)
	snafuNumbers := make([]string, counter)
	for i := 0; i < len(snafuNumbers); i++ {
		tmpNumbers[i] = in / value
		in = in % value
		value /= 5
	}
	for i := counter - 1; i >= 0; i-- {
		dec := tmpNumbers[i]
		if dec <= 2 {
			snafuNumbers[i] = strconv.Itoa(dec)
		} else if dec == 4 {
			tmpNumbers[i-1] += 1
			snafuNumbers[i] = "-"
		} else if dec == 3 {
			tmpNumbers[i-1] += 1
			snafuNumbers[i] = "="
		}
	}
	return snafuNumbers
}

func main() {
	file, err := os.Open("2022/day25/input")

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	numbers := make([][]string, 0)

	for scanner.Scan() {
		numbers = append(numbers, strings.Split(scanner.Text(), ""))
	}

	sum := 0
	for _, number := range numbers {
		sum += convertToDecimal(number)
	}
	//fmt.Println(sum)
	fmt.Println("Day 25.1:", strings.Join(convertToSnafu(sum), ""))
}
