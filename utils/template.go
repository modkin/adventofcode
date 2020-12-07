package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
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

	scanner := bufio.NewScanner(OpenFile("2020/day1/input"))
	for scanner.Scan() {
		foobar := strings.Split(scanner.Text(), ",")
		fmt.Println(foobar)
	}
}
