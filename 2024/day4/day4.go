package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printStringSlice(input [][]string) {
	for _, i2 := range input {
		fmt.Println(i2)
	}
}

func search(input [][]string) int {
	ylen := len(input) - 3
	xlen := len(input[0]) - 3
	totalNum := 0
	for x := 3; x < ylen; x++ {
		for j := 3; j < xlen; j++ {

			//fmt.Println(x, j, input[x][j])
			north := input[x][j] + input[x][j-1] + input[x][j-2] + input[x][j-3]
			northEast := input[x][j] + input[x+1][j-1] + input[x+2][j-2] + input[x+3][j-3]
			east := input[x][j] + input[x+1][j] + input[x+2][j] + input[x+3][j]
			southEast := input[x][j] + input[x+1][j+1] + input[x+2][j+2] + input[x+3][j+3]
			south := input[x][j] + input[x][j+1] + input[x][j+2] + input[x][j+3]
			southWest := input[x][j] + input[x-1][j+1] + input[x-2][j+2] + input[x-3][j+3]
			west := input[x][j] + input[x-1][j] + input[x-2][j] + input[x-3][j]
			northwest := input[x][j] + input[x-1][j-1] + input[x-2][j-2] + input[x-3][j-3]
			tries := []string{north, northEast, east, southEast, south, southWest, west, northwest}
			fmt.Println(tries)
			for _, try := range tries {
				if try == "XMAS" {
					totalNum += 1
				}
			}
		}
		fmt.Println(totalNum)

	}
	return totalNum
}

func searchx_mas(input [][]string) int {
	ylen := len(input) - 3
	xlen := len(input[0]) - 3
	totalNum := 0
	for x := 3; x < ylen; x++ {
		for j := 3; j < xlen; j++ {

			if input[x][j] == "A" {
				if (input[x+1][j+1] == "M" && input[x-1][j-1] == "S") || (input[x+1][j+1] == "S" && input[x-1][j-1] == "M") {
					if (input[x+1][j-1] == "M" && input[x-1][j+1] == "S") || (input[x+1][j-1] == "S" && input[x-1][j+1] == "M") {
						totalNum += 1
					}
				}
			}
		}
		fmt.Println(totalNum)

	}
	return totalNum
}

func main() {
	file, err := os.Open("2024/day4/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var words [][]string

	first := true

	var tmp []string
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "")
		if first {

			for i := 0; i < 3; i++ {
				tmp = make([]string, len(split)+6)
				words = append(words, tmp)
			}
			first = false
		}
		boundary := make([]string, 3)
		split = append(boundary, split...)
		split = append(split, boundary...)
		words = append(words, split)

	}
	for i := 0; i < 3; i++ {
		words = append(words, tmp)
	}

	printStringSlice(words)
	fmt.Println(search(words))
	fmt.Println("Day 4.2:", searchx_mas(words))
}
