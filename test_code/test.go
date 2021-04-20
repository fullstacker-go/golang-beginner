package main

import "fmt"

var num1 int = 5
var num2 = 7

func main() {

	fmt.Println("Hello World")
	add(num1, num2)
}
func add(x, y int) int {

	result := x + y
	return result

}
