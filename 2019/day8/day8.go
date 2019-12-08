package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func countDigitInLayer(layer []string, digit string) (count int) {
	for _, elem := range layer {
		count += strings.Count(elem, digit)
		//fmt.Println(elem, ": ", strings.Count(elem, digit))
	}
	return
}

func main() {
	content, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	width := 25
	entriesPerLayer := 6
	layers := len(content) / entriesPerLayer / width
	pixels := make([][]string, layers)
	for i, _ := range pixels {
		pixels[i] = make([]string, entriesPerLayer)
	}

	for layer := 0; layer < layers; layer++ {
		for entries := 0; entries < entriesPerLayer; entries++ {
			layerStart := layer * (len(content) / layers)
			start := layerStart + entries*width
			pixels[layer][entries] = string(content[start : start+width])
		}
	}
	//fmt.Println("Layers: ", len(pixels))

	minCount := math.MaxInt32
	var minLayer int
	for i, elem := range pixels {
		layerCount := countDigitInLayer(elem, strconv.Itoa(0))
		//fmt.Println("Layer: ", i, ": ", layerCount)
		//fmt.Println(elem)
		if layerCount < minCount {
			minCount = layerCount
			minLayer = i
		}
	}
	//fmt.Println("Min layer: ", minLayer)
	ones := countDigitInLayer(pixels[minLayer], strconv.Itoa(1))
	twos := countDigitInLayer(pixels[minLayer], strconv.Itoa(2))
	fmt.Println("Task 8.1", ones*twos)

	picture := make([][]string, entriesPerLayer)
	for i, _ := range picture {
		picture[i] = make([]string, width)
	}

	for _, layer := range pixels {
		for i, row := range layer {
			for j, pix := range row {
				if picture[i][j] == "" {
					switch pix {
					//1
					case 48:
						picture[i][j] = " "
					case 49:
						picture[i][j] = "#"
					}
				}
			}
		}
	}
	for _, line := range picture {
		fmt.Println(line)
	}

}
