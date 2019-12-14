package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Recipe struct {
	input     map[string]int
	outputMat string
	output    int
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	factory := make(map[string]Recipe)
	scanner := bufio.NewScanner(file)
	pool := make(map[string]int64)
	for scanner.Scan() {
		recipe := strings.Split(scanner.Text(), " => ")
		outEntry := strings.Split(recipe[1], " ")
		outQuant, _ := strconv.Atoi(outEntry[0])
		//outputMap := map[string]int{outEntry[1] : outQuant}
		inputMap := make(map[string]int)
		for _, in := range strings.Split(recipe[0], ", ") {
			inEntry := strings.Split(in, " ")
			inQuant, _ := strconv.Atoi(inEntry[0])
			inputMap[inEntry[1]] = inQuant
		}
		newRecipe := Recipe{
			input:     inputMap,
			outputMat: outEntry[1],
			output:    outQuant,
		}
		factory[outEntry[1]] = newRecipe
		pool[outEntry[1]] = 0
	}

	getOre := func(fuel int64) int64 {
		running := true
		pool["FUEL"] = fuel
		for running {
			running = false
			for mat, amount := range pool {
				if mat != "ORE" && amount > 0 {
					running = true
					amountOfOutUp := (amount + int64(factory[mat].output) + -1) / int64(factory[mat].output)
					pool[mat] = -(amountOfOutUp*int64(factory[mat].output) - amount)
					for inputMat, inputMatAmount := range factory[mat].input {
						pool[inputMat] += int64(inputMatAmount) * amountOfOutUp
					}
				}
			}
		}
		return pool["ORE"]
	}

	//needed := produce(factory, pool, factory["FUEL"], 1, 1)
	ore1 := getOre(1)
	fmt.Println("Task 14.1: ", ore1)

	maxOre := int64(1_000_000_000_000)
	step := maxOre
	for fuel := maxOre / ore1; ; fuel += step {
		for mat, _ := range pool {
			pool[mat] = 0
		}
		ore := getOre(fuel)
		fuel += step
		fmt.Println("fuel ", fuel)
		fmt.Println(ore)
		if ore > maxOre {
			if step == 1 {
				fmt.Println("Task 14.2: ", fuel-1)
				break
			} else {
				fmt.Println("smaller")
				fuel -= step
				step /= 10
			}
		}

	}
}
