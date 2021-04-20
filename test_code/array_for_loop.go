package main

import "fmt"

func main() {
	intArray := [5]int{10, 20, 30, 40, 50}
	arr := [...]int{10, 12, 14, 56, 66, 23} //array initialization with ellipses.dynamic length
	fmt.Println(arr)
	fmt.Println("\n---------------Example 1--------------------\n")
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}

	fmt.Println("\n---------------Example 2--------------------\n")
	for index, element := range intArray {
		fmt.Println(index, "=>", element)

	}

	fmt.Println("\n---------------Example 3--------------------\n")
	for _, value := range intArray {
		fmt.Println(value)
	}

	j := 0
	fmt.Println("\n---------------Example 4--------------------\n")
	for range intArray {
		fmt.Println(intArray[j])
		j++
	}
}
