package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	input := []byte("abcdef")
	//salt := 1
	//for true {
	salt := 609043
		input = append(input, []byte()
	hash := md5.Sum(input)
	//}
	fmt.Println(hash)
	fmt.Println(hex.EncodeToString(hash[:]))
}
