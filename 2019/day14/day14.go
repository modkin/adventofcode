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
	if amount, ok := material.input["ORE"]; ok {
		pool[material.outputMat] -= amount * multiplier
		return amount * multiplier
	} else {
		for mat, amount := range material.input {
			totalOre += reverse(factory, pool, factory[mat], amount)
		}
	}
	pool[material.outputMat] -= amount * multiplier
	return totalOre
}

func produce(factory map[string]Recipe, pool map[string]int, material Recipe, required int) int {
	//required -= pool[material.outputMat]
	multiplier := (required + material.output - 1) / material.output
	totalOre := 0
	if amount, ok := material.input["ORE"]; ok {
		pool[material.outputMat] += (material.output * multiplier) % required
		return amount * multiplier
	} else {
		for mat, amount := range material.input {
			totalOre += produce(factory, pool, factory[mat], amount)
		}
	}
	pool[material.outputMat] += (material.output * multiplier) % required
	return totalOre
}

func main() {
	file, err := os.Open("./testInput")
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

	//fmt.Println(factory)
	needed := produce(factory, pool, factory["FUEL"], 1)
	fix := true
	tooMuch := 0
	fmt.Println(pool)
	for fix {
		fix = false
		for mat, amount := range pool {
			if amount > factory[mat].output {
				fix = true
				tooMuch += reverse(factory, pool, factory[mat], amount)
			}
		}
	}
	fmt.Println("Need ", needed)
	fmt.Println(pool)
	fmt.Println(needed - tooMuch)

}
