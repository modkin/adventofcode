package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func read(data interface{}) int {
	sum := 0
	switch v := data.(type) {
	case []interface{}:
		for _, elem := range v {
			sum += read(elem)
		}
	case map[string]interface{}:
		for _, value := range v {
			sum += read(value)
		}
	case float64:
		sum += int(v)
	default:
		fmt.Println(sum)
	}
	return sum
}

func main() {
	jsonByte, err := ioutil.ReadFile("2015/day12/input")
	if err != nil {
		panic(err)
	}

	var data interface{}

	err = json.Unmarshal([]byte(jsonByte), &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(read(data))
	//fmt.Println(data.(type))
}
