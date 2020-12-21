package main

import (
	"adventofcode/utils"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(utils.OpenFile("2020/day21/input"))
	allergenMap := make(map[string][][]string)
	allIngredients := make(map[string]int)
	ingredientToAllgergenMap := make(map[string]map[string]int)
	toxicIngredients := make([]string, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "(")
		ingredients := strings.Split(strings.TrimSpace(line[0]), " ")
		allergen := strings.Split(strings.TrimSuffix(strings.TrimPrefix(line[1], "contains "), ")"), ", ")
		for _, ag := range allergen {
			allergenMap[ag] = append(allergenMap[ag], ingredients)
		}
		for _, ing := range ingredients {
			allIngredients[ing]++
		}
		for _, ing := range ingredients {
			if _, ok := ingredientToAllgergenMap[ing]; !ok {
				ingredientToAllgergenMap[ing] = make(map[string]int)
			}
			for _, alg := range allergen {
				ingredientToAllgergenMap[ing][alg]++
			}
		}
	}
	smallerAlgIngMap := make(map[string]map[string]bool)
	for allergen := range allergenMap {
		max := 0
		for _, value := range ingredientToAllgergenMap {
			if value[allergen] > max {
				max = value[allergen]
			}
		}
		possibleIng := make([]string, 0)
		for ing, value := range ingredientToAllgergenMap {
			if value[allergen] == max {
				possibleIng = append(possibleIng, ing)
			}
		}
		smallerAlgIngMap[allergen] = make(map[string]bool)
		for _, elem := range possibleIng {
			smallerAlgIngMap[allergen][elem] = true
		}
	}
	for i := 0; i < 8; i++ {
		for alg, ings := range smallerAlgIngMap {
			if len(ings) == 1 {
				for oAlg, oIngs := range smallerAlgIngMap {
					if alg != oAlg {
						for key := range ings {
							delete(oIngs, key)
						}
					}
				}
			}
		}
	}

	sum := 0
	for ing, amount := range allIngredients {
		if !utils.SliceContains(toxicIngredients, ing) {
			sum += amount
		}
	}
	fmt.Println("Task 21.1:", sum)
	fmt.Print("Task 21.2: ")
	printStr := fmt.Sprint(smallerAlgIngMap)
	tmp := strings.Split(strings.TrimPrefix(printStr, "map["), " ")
	for i, elem := range tmp {
		fmt.Print(strings.Split(strings.Split(elem, "[")[1], ":")[0])
		if i != len(tmp)-1 {
			fmt.Print(",")
		}
	}
}
