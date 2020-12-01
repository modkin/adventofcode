package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("2020/day1/input")
	if err != nil {
		panic(err)
	}

	for _, elem := range content {
		fmt.Println(elem)
	}

	file, err := os.Open("2020/day1/input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		foobar := strings.Split(scanner.Text(), ",")
		fmt.Println(foobar)
	}
}
