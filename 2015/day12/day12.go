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
	}
	return sum
}

func read2(data interface{}) int {
	sum := 0
	switch v := data.(type) {
	case []interface{}:
		for _, elem := range v {
			sum += read2(elem)
		}
	case map[string]interface{}:
		noRed := true
		for _, value := range v {
			switch v1 := value.(type) {
			case string:
				if v1 == "red" {
					noRed = false
				}
			}
		}
		if noRed {
			for _, value := range v {
				sum += read2(value)
			}
		}
	case float64:
		sum += int(v)
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
	fmt.Println("Task 12.1: ", read(data))
	fmt.Println("Task 12.2: ", read2(data))
	//fmt.Println(data.(type))
}
