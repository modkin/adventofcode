package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func inRange(first int, second int, check int) bool {
	for i := first; i <= second; i++ {
		if check == i {
			return true
		}
	}
	return false
}

func main() {

	file, err := os.Open("2022/day4/input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	rangePairs := make([][2][2]int, 0)

	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), ",")
		fp := strings.Split(tmp[0], "-")
		fp1 := utils.ToInt(fp[0])
		fp2 := utils.ToInt(fp[1])
		fptmp := [2]int{fp1, fp2}
		sp := strings.Split(tmp[1], "-")
		sp1 := utils.ToInt(sp[0])
		sp2 := utils.ToInt(sp[1])
		sptmp := [2]int{sp1, sp2}
		rangePairs = append(rangePairs, [2][2]int{fptmp, sptmp})

	}

	counter := 0
	for _, pair := range rangePairs {
		if pair[1][0] >= pair[0][0] && pair[1][1] <= pair[0][1] {
			counter++
		} else if pair[0][0] >= pair[1][0] && pair[0][1] <= pair[1][1] {
			counter++
		}
	}

	fmt.Println("Day 3.1:", counter)

	counter2 := 0
outer:
	for _, pair := range rangePairs {
		for i := pair[0][0]; i <= pair[0][1]; i++ {
			if inRange(pair[1][0], pair[1][1], i) {
				counter2++
				continue outer
			}
		}
		//for i := pair[1][0]; i < pair[1][1]; i++ {
		//	if inRange(pair[0][1], pair[0][1], i) {
		//		counter2++
		//		continue outer
		//	}
		//}
	}
	fmt.Println("Day 3.1:", counter2)

}
