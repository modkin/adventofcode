package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type getCurrentBitFunction func([][]string) string

func getGammaEpsilon(splitlines [][]string) [2]string {
	gammaString := ""
	for i := 0; i < len(splitlines[0]); i++ {
		verticalLine := ""
		for _, s := range splitlines {
			verticalLine += s[i]
		}
		if float64(strings.Count(verticalLine, "1")) >= float64(len(splitlines))/2 {
			gammaString += "1"
		} else {
			gammaString += "0"
		}
	}
	epsilonString := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(gammaString, "1", "X"), "0", "1"), "X", "0")
	return [2]string{gammaString, epsilonString}
}

func getCxyCO2(splitlines [][]string, fn getCurrentBitFunction) int64 {
	cleanup := make([][]string, 0)
	for _, splitline := range splitlines {
		cleanup = append(cleanup, splitline)
	}
	for i := 0; i < len(splitlines[0]) && len(cleanup) != 1; i++ {
		currentBit := string(fn(cleanup)[i])
		idx := 0
		for _, line := range cleanup {
			if line[i] == currentBit {
				cleanup[idx] = line
				idx++
			}
		}
		cleanup = cleanup[:idx]
	}
	result, _ := strconv.ParseInt(strings.Join(cleanup[0][:], ""), 2, 64)
	return result
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
	gammaEpsilon := getGammaEpsilon(splitlines)
	gamma, _ := strconv.ParseInt(gammaEpsilon[0], 2, 32)
	epsilon, _ := strconv.ParseInt(gammaEpsilon[1], 2, 32)
	fmt.Println("Day 3.1:", gamma*epsilon)

	getGamma := func(input [][]string) string {
		return getGammaEpsilon(input)[0]
	}
	getEpsilon := func(input [][]string) string {
		return getGammaEpsilon(input)[1]
	}
	fmt.Println("Day 3.2:", getCxyCO2(splitlines, getEpsilon)*getCxyCO2(splitlines, getGamma))
}
