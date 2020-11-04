package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {

	input := []byte("ckczppom")
	salt := 1
	first := false
	second := false
	for true {
		testInput := append(input, []byte(strconv.Itoa(salt))...)
		hash := md5.Sum(testInput)
		if hash[0] == byte(0) && hash[1] == byte(0) && hash[2] < uint8(16) && !first {

			fmt.Println("Task 4.1:", salt)
			//fmt.Println(hex.EncodeToString(hash[:]))
			first = true
		}
		if hash[0] == byte(0) && hash[1] == byte(0) && hash[2] == byte(0) && !second {

			fmt.Println("Task 4.2:", salt)
			//fmt.Println(hex.EncodeToString(hash[:]))
			second = true
		}
		if first && second {
			break
		}
		salt += 1
	}
}
