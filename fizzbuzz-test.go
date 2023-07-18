package main

import (
	"fmt"
)

func fizzBuzz(angka int) {
	if angka%3 == 0 && angka%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if angka%3 == 0 {
		fmt.Println("Fizz")
	} else if angka%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println("")
	}
}

func main() {

	for j := 1; j <= 100; j++ {
		fizzBuzz(j)

	}

}
