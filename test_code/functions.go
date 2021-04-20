package main

import (
	"fmt"
	"strings"
)

func main() {
	num1 := 12
	num2 := 14
	name := "tanVir"
	course := "go funDamentals"

	bestFinish := bestIplFinishes(12, 34, 54, 15, 20, 67, 44)
	fmt.Println(bestFinish)

	printHello()

	fmt.Println(add(num1, num2))

	fmt.Println(changeCase(name, course))

	func() {
		fmt.Println("Self Calling Function")
	}() //Self calling Annonymous function
}

// simple function

func printHello() {
	fmt.Println("Hello!")
}

//Function with parameters and single return value

func add(x, y int) int {

	return x + y
}

//Function with multiple return values

func changeCase(s1, s2 string) (str1, str2 string) {
	s1 = strings.Title(s1)
	s2 = strings.ToUpper(s2)
	return s1, s2
}

//Variadic Functions
//These functions are used when we dont know exactly
//how many parameters will be passed to function

func bestIplFinishes(finishes ...int) int {
	best := finishes[0]

	for _, val := range finishes {
		if val > best {
			best = val
		}
	}
	return best
}
