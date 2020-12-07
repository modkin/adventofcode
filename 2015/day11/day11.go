package main

import "fmt"

func passwordValid(password []byte) bool {
	straight, noIOL, double := false, true, false
	first := '0'
	for i, b := range password {
		if i < len(password)-3 {
			if int(password[i])-int(password[i+1]) == -1 && int(password[i+1])-int(password[i+2]) == -1 {
				straight = true
			}
		}
		if b == 'i' || b == 'o' || b == 'l' {
			noIOL = false
		}
		if i < len(password)-1 {
			if password[i]-password[i+1] == 0 {
				if first == '0' {
					first = int32(password[i])
				} else {
					if first != int32(password[i]) {
						double = true
					}
				}
			}
		}
	}
	return straight && noIOL && double
}
func main() {
	password := []byte("cqjxjnds")
	fmt.Println(passwordValid(password))
	for {
		pos := len(password) - 1
		for {
			password[pos]++
			if password[pos] == 'z'+1 {
				password[pos] = 'a'
				pos--
			} else {
				break
			}
		}
		if passwordValid(password) {
			break
		}
	}
	fmt.Println(string(password[:]))
}
