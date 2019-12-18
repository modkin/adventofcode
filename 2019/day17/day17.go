package main

import (
	"adventofcode/2019/computer"
	"adventofcode/utils"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func rotate90right(vec [2]int) (ret [2]int) {
	ret[0] = -1 * vec[1]
	ret[1] = vec[0]
	return
}

func rotate90left(vec [2]int) (ret [2]int) {
	ret[0] = vec[1]
	ret[1] = -1 * vec[0]
	return
}

func printPaintMap(paintMap map[[2]int]string) {
	maxX, maxY := math.MinInt32, math.MinInt32
	for pos, _ := range paintMap {
		if pos[0] > maxX {
			maxX = pos[0]
		}

		if pos[1] > maxY {
			maxY = pos[1]
		}
	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			fmt.Print(paintMap[[2]int{x, y}])
		}
		fmt.Println()
	}
}

func findItersections(paintMap map[[2]int]string) (result int) {
	maxX, maxY := math.MinInt32, math.MinInt32
	for pos, _ := range paintMap {
		if pos[0] > maxX {
			maxX = pos[0]
		}

		if pos[1] > maxY {
			maxY = pos[1]
		}
	}
	for y := 1; y <= maxY-1; y++ {
		for x := 1; x < maxX-1; x++ {
			if paintMap[[2]int{x, y}] == "#" {
				if paintMap[[2]int{x - 1, y}] == "#" && paintMap[[2]int{x + 1, y}] == "#" && paintMap[[2]int{x, y - 1}] == "#" && paintMap[[2]int{x, y + 1}] == "#" {
					paintMap[[2]int{x, y}] = "O"
					result += y * x
				}
			}
		}
	}
	return
}

func runCamera(shipMap map[[2]int]string, outputCh <-chan int64, quit <-chan bool) {
	running := true
	x := 0
	y := 0
	for running {
		select {
		case input := <-outputCh:
			if input == 10 {
				y++
				x = 0
			} else {
				shipMap[[2]int{x, y}] = string(rune(input))
				x++
			}
		case <-quit:
			running = false
		}
	}
}

func findPath(shipMap map[[2]int]string) string {
	var pos [2]int

	step := func(start [2]int, direction [2]int) [2]int {
		return [2]int{start[0] + direction[0], start[1] + direction[1]}
	}
	for coords, elem := range shipMap {
		/// this is specific for my input it might be not working with others if the robo is not facing up
		if elem == "^" {
			pos = coords
		}
	}
	fmt.Println("Start: ", pos)
	path := "R,"

	dir := [2]int{1, 0}
	for {
		steplength := 0
		for shipMap[step(pos, dir)] == "#" {
			pos = step(pos, dir)
			steplength++
		}
		path += fmt.Sprint(steplength, ",")
		if shipMap[step(pos, rotate90left(dir))] == "#" {
			path += "L,"
			dir = rotate90left(dir)
		} else if shipMap[step(pos, rotate90right(dir))] == "#" {
			path += "R,"
			dir = rotate90right(dir)
		} else {
			break
		}
	}
	return path
}

func main() {
	content, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	contentString := strings.Split(string(content), ",")
	intcode := make([]int64, len(contentString))
	for pos, elem := range contentString {
		intcode[pos] = utils.ToInt64(elem)
	}
	inputCh := make(chan int64)
	outputCh := make(chan int64)
	quit := make(chan bool)

	shipMap := make(map[[2]int]string)

	go computer.ProcessIntCode(intcode, inputCh, outputCh, quit)

	runCamera(shipMap, outputCh, quit)
	path := findPath(shipMap)
	task1 := findItersections(shipMap)
	//printPaintMap(shipMap)
	fmt.Println("Task 17.1: ", task1)

	maxX, maxY := math.MinInt32, math.MinInt32
	for pos, _ := range shipMap {
		if pos[0] > maxX {
			maxX = pos[0]
		}
		if pos[1] > maxY {
			maxY = pos[1]
		}
	}
	printPaintMap(shipMap)
	fmt.Println(path)
	A := "R,8,L,10,R,8,R,12,R,8,L,8,L,12"
	B := "L,12,L,10,L,8"
	C := "R,8,L,10,R,8"
	path = strings.Replace(path, A, "A", -1)
	path = strings.Replace(path, B, "B", -1)
	path = strings.Replace(path, C, "C", -1)
	fmt.Println(path)
	mainProg := strings.Trim(path, ",")
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(mainProg)

	intcode[0] = 2
	go computer.ProcessIntCode(intcode, inputCh, outputCh, quit)
	//for i := 0; i < maxX*maxY; i++ {
	//	<-outputCh
	//}
	//runCamera(shipMap,outputCh,quit)

	Arune := append([]rune(A), '\n')
	Brune := append([]rune(B), '\n')
	Crune := append([]rune(C), '\n')
	mainProgramm := append([]rune(mainProg), '\n')
	functions := append(append(Arune, Brune...), Crune...)
	videoFeed := []rune("n\n")
	var total []rune
	total = append(append(mainProgramm, functions...), videoFeed...)
	fmt.Println(total)
	counter := 0
	running := true
	for running {
		select {
		case <-outputCh:
		case <-quit:
			running = false
		case inputCh <- int64(total[counter]):
			fmt.Println("DEBUG: ", total[counter])
			counter++
			if counter == len(total) {
				running = false
			}
		}
	}

	shipMap = make(map[[2]int]string)

	runCamera(shipMap, outputCh, quit)
	findItersections(shipMap)
	//printPaintMap(shipMap)
}
