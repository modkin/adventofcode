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

func reverse(factory map[string]Recipe, pool map[string]int, material Recipe, amount int) int {
	multiplier := amount / material.output
	totalOre := 0
	if ReqAmount, ok := material.input["ORE"]; ok {
		pool[material.outputMat] -= material.output * multiplier
		return ReqAmount * multiplier
	} else {
		for mat, InnerAmount := range material.input {
			totalOre += reverse(factory, pool, factory[mat], InnerAmount)
		}
	}
	pool[material.outputMat] -= material.output * multiplier
	return totalOre
}

func produce(factory map[string]Recipe, pool map[string]int, material Recipe, required int, poolMult int) int {
	//required -= pool[material.outputMat]
	multiplier := (required + material.output - 1) / material.output
	totalOre := 0
	if ReqAmount, ok := material.input["ORE"]; ok {
		pool[material.outputMat] += ((material.output * multiplier) - required) * poolMult
		return ReqAmount * multiplier
	} else {
		for mat, InnerAmount := range material.input {
			totalOre += produce(factory, pool, factory[mat], InnerAmount, required)
		}
	}
	pool[material.outputMat] += ((material.output * multiplier) - required) * poolMult
	return totalOre * multiplier
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	factory := make(map[string]Recipe)
	scanner := bufio.NewScanner(file)
	pool := make(map[string]int)
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

	fmt.Println(factory)
	//needed := produce(factory, pool, factory["FUEL"], 1, 1)
	running := true
	pool["FUEL"] = 1
	fmt.Println(pool)
	fmt.Println("START")
	for running {
		running = false
		for mat, amount := range pool {
			if mat != "ORE" && amount > 0 {
				running = true
				amountOfOutUp := (amount + factory[mat].output + -1) / factory[mat].output
				pool[mat] = -(amountOfOutUp*factory[mat].output - amount)
				for inputMat, inputMatAmount := range factory[mat].input {
					pool[inputMat] += inputMatAmount * amountOfOutUp
				}
			}
		}
	}
	fmt.Println("Task 14.1: ", pool["ORE"])

}
