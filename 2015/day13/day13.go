package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func permutate(allPermutations [][]string, names []string, k int, last string) {
	names = utils.CopyStringSlice(names)
	if k == 1 {
		if names[len(names)-1] == last {
			for idx, permutation := range allPermutations {
				if permutation == nil {
					allPermutations[idx] = utils.CopyStringSlice(names)
					break
				}
			}
		}
		return
	}
	// Generate permutations for kth swapped with each k-1 initial
	for i := 0; i < k; i += 1 {
		permutate(allPermutations, utils.CopyStringSlice(names), k-1, last)
		// Swap choice dependent on parity of k (even or odd)
		if k%2 == 1 {
			names[i], names[k-1] = names[k-1], names[i] // zero-indexed, the kth is at k-1
		} else {
			names[0], names[k-1] = names[k-1], names[0]
		}

	}
}

func main() {
	file, err := os.Open("2015/day13/input.txt")
	if err != nil {
		panic(err)
	}

	happinessMap := make(map[string]map[string]int)
	number := regexp.MustCompile(` \d* `)
	var allNames []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		happy := utils.ToInt(number.FindString(line))
		name := lineSplit[0]
		against := strings.Trim(lineSplit[len(lineSplit)-1], ".")
		if _, ok := happinessMap[name]; !ok {
			happinessMap[name] = make(map[string]int)
		}
		if strings.Contains(line, "gain") {
			happinessMap[name][against] = happy
		} else {
			happinessMap[name][against] = -happy
		}
	}
	for name := range happinessMap {
		allNames = append(allNames, name)
	}
	fmt.Println(allNames)
	combination := utils.Factorial(len(allNames) - 1)

	fmt.Println("number of combinations: ", combination)
	allPermutations := make([][]string, combination)
	//for i := 0; i < combination; i++ {
	//	allPermutations[i] = utils.CopyStringSlice(allNames)
	//}

	permutate(allPermutations, allNames, len(allNames), allNames[len(allNames)-1])

	fmt.Println(allPermutations)

	mostHappy := 0
	for _, seating := range allPermutations {
		seatingHappy := 0
		for i, person := range seating {
			left := i - 1
			if left == -1 {
				left = len(seating) - 1
			}
			right := i + 1
			if right == len(seating) {
				right = 0
			}
			seatingHappy += happinessMap[person][seating[left]]
			seatingHappy += happinessMap[person][seating[right]]
		}
		if seatingHappy > mostHappy {
			mostHappy = seatingHappy
		}
	}
	fmt.Println(mostHappy)

}
