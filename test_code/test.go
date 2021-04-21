package main

import "fmt"

var num1 int = 5
var num2 = 7
var num3 float64 = 2.5

func main() {

	fmt.Println("Hello World")
	result := add(num1, num2)
	fmt.Println("The result of the addition is ", result)
	fmt.Println(multiply(num2, num3))
}
func add(x, y int) int {

	result := x + y
	return result

}
func multiply(a int, b float64) int {
	return a * int(b)
}
