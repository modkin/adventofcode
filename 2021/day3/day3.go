package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getCommonBits(splitlines [][]string) ([]string, []string) {
	gamma := make([]string, 0)
	epsilon := make([]string, 0)
	for i := 0; i < len(splitlines[0]); i++ {
		ones := 0
		zeros := 0
		for _, splitline := range splitlines {
			if splitline[i] == "0" {
				zeros++
			} else {
				ones++
			}
		}
		if ones > zeros {
			gamma = append(gamma, "1")
			epsilon = append(epsilon, "0")
		} else if ones == zeros {
			gamma = append(gamma, "1")
			epsilon = append(epsilon, "0")
		} else {
			gamma = append(gamma, "0")
			epsilon = append(epsilon, "1")
		}
	}
	return gamma, epsilon
}

func main() {
	file, err := os.Open("2021/day3/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	splitlines := make([][]string, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		splitlines = append(splitlines, line)
	}
	gamma := make([]string, 0)
	epsilon := make([]string, 0)
	for i := 0; i < len(splitlines[0]); i++ {
		ones := 0
		zeros := 0
		for _, splitline := range splitlines {
			if splitline[i] == "0" {
				zeros++
			} else {
				ones++
			}
		}
		if ones > zeros {
			gamma = append(gamma, "1")
			epsilon = append(epsilon, "0")
		} else {
			gamma = append(gamma, "0")
			epsilon = append(epsilon, "1")
		}
	}
	gammaInt, _ := strconv.ParseInt(strings.Join(gamma[:], ""), 2, 64)
	epsilonInt, _ := strconv.ParseInt(strings.Join(epsilon[:], ""), 2, 64)
	fmt.Println(gammaInt * epsilonInt)

	cleanup := make([][]string, 0)
	for _, splitline := range splitlines {
		cleanup = append(cleanup, splitline)
	}
	for i := 0; i < len(splitlines[0]); i++ {
		gamma, epsilon = getCommonBits(cleanup)
		currentBit := gamma[i]
		newCleanup := make([][]string, 0)
		for _, splitline := range cleanup {
			if splitline[i] == currentBit {
				newCleanup = append(newCleanup, splitline)
			}
		}
		cleanup = newCleanup
	}
	fmt.Println(cleanup)
	oxygenGenRating, _ := strconv.ParseInt(strings.Join(cleanup[0][:], ""), 2, 64)
	fmt.Println("OXY", oxygenGenRating)

	cleanup = make([][]string, 0)
	for _, splitline := range splitlines {
		cleanup = append(cleanup, splitline)
	}
	for i := 0; i < len(splitlines[0]); i++ {
		gamma, epsilon = getCommonBits(cleanup)
		currentBit := epsilon[i]
		newCleanup := make([][]string, 0)
		for _, splitline := range cleanup {
			if splitline[i] == currentBit {
				newCleanup = append(newCleanup, splitline)
			}
		}
		cleanup = newCleanup
		if len(cleanup) == 1 {
			break
		}
	}
	fmt.Println(cleanup)
	co2scrubrating, _ := strconv.ParseInt(strings.Join(cleanup[0][:], ""), 2, 64)
	fmt.Println("CO2", co2scrubrating)
	fmt.Println(oxygenGenRating * co2scrubrating)
}
