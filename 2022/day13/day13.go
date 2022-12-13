package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// return the index AFTER the closing ']'
func findClosing(input string) int {
	counter := 0
	for idx, str := range input {
		if str == '[' {
			counter++
		} else if str == ']' {
			counter--
		}
		if counter == 0 {
			return idx
		}
	}
	return math.MaxInt
}

type recList struct {
	val  int
	list []recList
}

func compareInt(left int, right int) int {
	if left < right {
		return 1
	} else if right < left {
		return 2
	} else if right == left {
		return 3
	} else {
		fmt.Println("ERROR")
		return math.MaxInt
	}
}

// return
// 1 : right order
// 2 : wrong order
// 3 continue
func compare(left string, right string) int {
	if left == "" && right != "" {
		return 1
	}
	if right == "" && left != "" {
		return 2
	}
	if right == "" && left == "" {
		return 3
	}
	var leftNum, rightNum int
	var err error
	leftIsNum := false
	if leftNum, err = strconv.Atoi(strings.Split(left, ",")[0]); err == nil {
		leftIsNum = true
	}
	rightIsNum := false
	if rightNum, err = strconv.Atoi(strings.Split(right, ",")[0]); err == nil {
		rightIsNum = true
	}
	if leftIsNum && rightIsNum {
		if ret := compareInt(leftNum, rightNum); ret != 3 {
			return ret
		}
	}
	leftIsList := false
	var leftSub, rightSub string
	if left[0] == '[' {
		leftIsList = true
		leftSub = left[1:findClosing(left)]
		//left = left[strings.Index(left, "]")+2:]
	}
	rightIsList := false
	if right[0] == '[' {
		rightIsList = true
		rightSub = right[1:findClosing(right)]
		//right = right[strings.Index(right, "]")+2:]
	}
	if leftIsList && rightIsList {
		ret := compare(leftSub, rightSub)
		if ret != 3 {
			return ret
		}
	}
	if leftIsNum && rightIsList {
		ret := compare(strconv.Itoa(leftNum), rightSub)
		if ret != 3 {
			return ret
		}
	}
	if leftIsList && rightIsNum {
		ret := compare(leftSub, strconv.Itoa(rightNum))
		if ret != 3 {
			return ret
		}
	}
	nextLeft := strings.Join(strings.Split(left, ",")[1:], ",")
	nextRight := strings.Join(strings.Split(right, ",")[1:], ",")
	if nextLeft == "" && nextRight == "" {
		return 3
	}
	if len(nextLeft) > 0 && nextLeft[0] == '[' {
		if findClosing(nextLeft) == len(nextLeft) {
			nextLeft = nextLeft[1 : len(nextLeft)-1]
		}
	}

	if len(nextRight) > 0 && nextRight[0] == '[' {
		if findClosing(nextRight) == len(nextRight) {
			nextRight = nextRight[1 : len(nextRight)-1]
		}
	}
	//if left == nextLeft && nextRight == right {
	//	return 3
	//}

	return compare(nextLeft, nextRight)
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

	for scanner.Scan() {
		counter++
		if scanner.Text() == "" {
			scanner.Scan()
		}
		left := scanner.Text()
		scanner.Scan()
		right := scanner.Text()
		comp := compare(left[1:len(left)-1], right[1:len(right)-1])
		if comp == 1 {
			rightOrder += counter
		}
	}

	fmt.Println(rightOrder)
}
