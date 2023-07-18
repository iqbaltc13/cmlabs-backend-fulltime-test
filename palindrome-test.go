package main

import (
	"fmt"
	"strings"
)

func isPalindrome(sentence string) interface{} {
	fmt.Println(sentence)
	if sentence == balikString(sentence) {
		return true
	}
	return false
}

func balikString(sentence string) (result string) {
	for _, v := range sentence {
		result = string(v) + result
	}
	return result
}

func main() {
	var sentence string
	fmt.Print("Masukkan tulisan: ")
	fmt.Scanln(&sentence)
	if isPalindrome(strings.ToUpper(sentence)) == true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
