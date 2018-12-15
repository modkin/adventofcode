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
		fmt.Println(update)
	}
	tree = append(tree[:start], tree[(start+metalen+2):]...)
	return update
}

func metaDataSum() int {

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
	fmt.Println(data)
	return processTree(data, 0)

}

func Task1() {
	fmt.Println(metaDataSum())
}
