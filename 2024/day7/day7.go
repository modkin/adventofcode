package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	expected int
	list     []int
}

func concatInts(one, two int) int {
	return utils.ToInt(strconv.Itoa(one) + strconv.Itoa(two))

}

func main() {

	file, err := os.Open("2024/day7/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	equations := make([]equation, 0)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		first := utils.ToInt(strings.Trim(fields[0], ":"))
		newEq := equation{first, make([]int, len(fields)-1)}
		for i, s := range fields[1:] {
			newEq.list[i] = utils.ToInt(s)
		}
		equations = append(equations, newEq)
	}

	fmt.Println(equations)

	sum := 0
outer:
	for _, eq := range equations {
		possibleResults := make([]int, 0)
		possibleResults = append(possibleResults, eq.list[0])
		for _, op := range eq.list[1:] {
			var newPossResulsts []int
			for _, result := range possibleResults {
				newPossResulsts = append(newPossResulsts, result+op)
				newPossResulsts = append(newPossResulsts, result*op)
				newPossResulsts = append(newPossResulsts, concatInts(result, op))
			}
			possibleResults = newPossResulsts

		}

		for _, result := range possibleResults {
			if result == eq.expected {
				sum += eq.expected
				//fmt.Println(expecedResults, possibleResults[i], possibleResults)
				continue outer
			}
		}

	}
	fmt.Println(sum)
}
