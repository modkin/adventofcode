package day7

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type worker struct {
	timer  int
	letter rune
}

func Solve() {
	file, err := os.Open("day7/day7-input.txt")
	if err != nil {
		panic(err)
	}

	nodeMap := make(map[rune][]rune)
	var depSlice []rune
	depMap := make(map[rune]bool)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		coords := strings.Split(word, " ")
		dep := rune([]rune(coords[1])[0])
		node := rune([]rune(coords[7])[0])
		nodeMap[node] = append(nodeMap[node], dep)
		depMap[dep] = true
		depMap[node] = true
	}
	for _, elem := range nodeMap {
		sort.Slice(elem, func(i, j int) bool {
			return i < j
		})
	}

	for key, _ := range depMap {
		depSlice = append(depSlice, key)
	}
	sort.Slice(depSlice, func(i, j int) bool {
		return depSlice[i] < depSlice[j]
	})

	for len(depSlice) != 0 {
		for depIdx, letter := range depSlice {
			if nodeMap[letter] == nil || len(nodeMap[letter]) == 0 {
				fmt.Print(string(letter))
				depSlice = append(depSlice[0:depIdx], depSlice[depIdx+1:]...)
				for keys, _ := range nodeMap {
					for idx, elem := range nodeMap[keys] {
						if elem == letter {
							nodeMap[keys] = append(nodeMap[keys][0:idx], nodeMap[keys][idx+1:]...)
							break
						}
					}
				}
				break
			}
		}
	}
	fmt.Println()
}

func Task1() {
	Solve()
}

func empty(workers []worker) bool {
	for _, task := range workers {
		if task.letter != 0 {
			return false
		}
	}
	return true
}

func Task2() {
	file, err := os.Open("day7/day7-input.txt")
	if err != nil {
		panic(err)
	}

	nodeMap := make(map[rune][]rune)
	var depSlice []rune
	depMap := make(map[rune]bool)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		coords := strings.Split(word, " ")
		dep := rune([]rune(coords[1])[0])
		node := rune([]rune(coords[7])[0])
		nodeMap[node] = append(nodeMap[node], dep)
		depMap[dep] = true
		depMap[node] = true
	}
	for _, elem := range nodeMap {
		sort.Slice(elem, func(i, j int) bool {
			return i < j
		})
	}

	for key, _ := range depMap {
		depSlice = append(depSlice, key)
	}
	sort.Slice(depSlice, func(i, j int) bool {
		return depSlice[i] < depSlice[j]
	})

	totaltime := 0
	timecost := 60
	workerNum := 5
	workers := make([]worker, workerNum)
	for {
		var remove []int
		for depIdx, letter := range depSlice {
			if nodeMap[letter] == nil || len(nodeMap[letter]) == 0 {
				for workerNum, worker := range workers {
					if worker.timer < 1 && worker.letter == 0 {
						workers[workerNum].timer = int(letter) - 64 + timecost
						workers[workerNum].letter = letter
						//fmt.Println("Start: ", totaltime, string(letter))
						remove = append(remove, depIdx)
						break
					}
				}
			}
		}
		for offset, removeIdx := range remove {
			depSlice = append(depSlice[0:removeIdx-offset], depSlice[removeIdx-offset+1:]...)
		}
		//fmt.Println(workers)
		for idx, _ := range workers {
			workers[idx].timer--
		}
		totaltime++
		for workIdx, worker := range workers {
			if worker.timer < 1 && worker.letter != 0 {
				//fmt.Println("Finish: ", totaltime, string(worker.letter))
				for keys, _ := range nodeMap {
					for idx, elem := range nodeMap[keys] {
						if elem == worker.letter {
							nodeMap[keys] = append(nodeMap[keys][0:idx], nodeMap[keys][idx+1:]...)
							break
						}
					}
				}
				workers[workIdx].letter = 0
			}
		}
		if len(depSlice) == 0 && empty(workers) {
			fmt.Println(totaltime)
			return
		}
	}
}
