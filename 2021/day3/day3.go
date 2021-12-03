package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type getCurrentBitFunction func([][]string) int64

func getGammaEpsilon(splitlines [][]string) [2]int64 {
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
	gamma, _ := strconv.ParseInt(gammaString, 2, 32)
	epsilon := gamma ^ (1<<len(splitlines[0]) - 1)
	return [2]int64{gamma, epsilon}
}

func getCxyCO2(splitlines [][]string, fn getCurrentBitFunction) int64 {
	cleanup := make([][]string, 0)
	for _, splitline := range splitlines {
		cleanup = append(cleanup, splitline)
	}
	for i := 0; i < len(splitlines[0]) && len(cleanup) != 1; i++ {
		currentBit := string(strconv.FormatInt(fn(cleanup), 2)[i])
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
	file, err := os.Open("2021/day3/testinput")
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
	fmt.Println("Day 3.1:", gammaEpsilon[0]*gammaEpsilon[1])

	getGamma := func(input [][]string) int64 {
		return getGammaEpsilon(input)[0]
	}
	getEpsilon := func(input [][]string) int64 {
		return getGammaEpsilon(input)[1]
	}
	fmt.Println(getCxyCO2(splitlines, getEpsilon))
	fmt.Println(getCxyCO2(splitlines, getGamma))

	cleanup := make([][]string, 0)
	for _, splitline := range splitlines {
		cleanup = append(cleanup, splitline)
	}
	for i := 0; i < len(splitlines[0]) && len(cleanup) != 1; i++ {
		gamma := getGammaEpsilon(cleanup)[0]
		currentBit := string(strconv.FormatInt(gamma, 2)[i])
		idx := 0
		for _, line := range cleanup {
			if line[i] == currentBit {
				cleanup[idx] = line
				idx++
			}
		}
		cleanup = cleanup[:idx]
	}
	fmt.Println(cleanup)
	co2scrubrating, _ := strconv.ParseInt(strings.Join(cleanup[0][:], ""), 2, 64)
	fmt.Println("CO2", co2scrubrating)
}
