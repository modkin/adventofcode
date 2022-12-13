package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type recList struct {
	val  int
	list []recList
}

func main() {
	file, err := os.Open("2022/day13/input")
	if err != nil {
		panic(err)
	}

	//grid := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	rightOrder := 0
	counter := 0

outer:
	for scanner.Scan() {
		counter++
		if scanner.Text() == "" {
			scanner.Scan()
		}
		left := strings.Split(scanner.Text(), ",")
		scanner.Scan()
		right := strings.Split(scanner.Text(), ",")
		//leftLevel := 0
		//rightLevel := 0
		for i := 0; ; i++ {
			if i == len(left) {
				rightOrder += counter
				continue outer
			}
			if i == len(right) {
				continue outer
			}
			leftCut := strings.Trim(left[i], "[")
			leftCut = strings.Trim(leftCut, "]")
			rightCut := strings.Trim(strings.Trim(right[i], "["), "]")
			if leftCut == "" && rightCut == "" {
				if strings.Count(left[i], "[") < strings.Count(right[i], "[") {
					rightOrder += counter
					continue outer
				}
				if strings.Count(left[i], "[") > strings.Count(right[i], "[") {
					continue outer
				}
			}
			if leftCut == "" {
				rightOrder += counter
				continue outer
			}
			if rightCut == "" {
				rightOrder += counter
				continue outer
			}
			leftNum := utils.ToInt(leftCut)
			rightNum := utils.ToInt(rightCut)
			if leftNum < rightNum {
				rightOrder += counter
				continue outer
			} else if rightNum < leftNum {
				continue outer
			}
		}
	}

	fmt.Println(rightOrder)
}
