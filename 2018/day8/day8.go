package day8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processTree(tree []int, start int) int {
	childs := tree[start]
	metalen := tree[start+1]
	update := 0
	for i := 0; i < childs; i++ {
		update += processTree(tree, start+2)
	}
	for i := 0; i < metalen; i++ {
		update += tree[start+2+i]
	}
	tree = append(tree[:start], tree[(start+metalen+2):]...)
	return update
}

func processTree2(tree []int, start int) int {
	childs := tree[start]
	childsSum := make([]int, childs)
	metalen := tree[start+1]
	update := 0
	if childs == 0 {
		for i := 0; i < metalen; i++ {
			update += tree[start+2+i]
		}
	} else {
		for i := 0; i < childs; i++ {
			childsSum[i] = processTree2(tree, start+2)
		}

		for i := 0; i < metalen; i++ {
			idx := tree[start+2+i] - 1
			if idx < len(childsSum) {
				update += childsSum[idx]
			}
		}
	}
	tree = append(tree[:start], tree[(start+metalen+2):]...)
	return update
}

func getData() []int {

	file, err := os.Open("day8/day8-input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	content := strings.Split(line, " ")
	var data []int
	for _, elem := range content {
		number, _ := strconv.Atoi(elem)
		data = append(data, number)
	}
	return data
}

func Task1() {
	data := getData()
	fmt.Println(processTree(data, 0))
	data = getData()
	fmt.Println(processTree2(data, 0))
}
