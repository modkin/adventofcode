package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filename := "./testInput1"
	content, err := ioutil.ReadFile(filename)
	ylist := strings.Split(string(content), "\n")
	xlist := strings.Split(ylist[0], "")
	xmax, ymax := len(xlist), len(ylist)

	fmt.Println(xmax, " ", ymax)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	starmap := make([][]bool, xmax)
	for i, _ := range starmap {
		starmap[i] = make([]bool, ymax)
	}
	fmt.Println(starmap)
	ycoord := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for xcoord, elem := range line {
			if elem == "#" {
				/// switch coords to use [x][y] indexing later on
				starmap[ycoord][xcoord] = true
			}
		}
		ycoord++
	}
	fmt.Println(starmap)
	fmt.Println(starmap[0])
}
