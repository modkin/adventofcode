package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	fmt.Println(splitlines)
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
	fmt.Println(gammaInt)
	epsilonInt, _ := strconv.ParseInt(strings.Join(epsilon[:], ""), 2, 64)
	fmt.Println(epsilonInt)
	fmt.Println(gammaInt * epsilonInt)

}
