package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("2021/day8/input")
	scanner := bufio.NewScanner(file)
	if err != nil {
		panic(err)
	}
	input := make([][]string, 0)
	output := make([][]string, 8)
	for scanner.Scan() {
		inOut := strings.Split(scanner.Text(), "|")
		output = append(output, strings.Fields(inOut[1]))
		input = append(input, strings.Fields(inOut[0]))
	}
	total := 0
	for _, i := range output {
		for _, o := range i {
			if len(o) == 2 || len(o) == 3 || len(o) == 7 || len(o) == 4 {
				total++
			}
		}
	}
	fmt.Println(total)
}
