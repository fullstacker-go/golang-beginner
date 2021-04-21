package main

import "fmt"

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, fmt.Errorf("can't divide by zero")
	}
	return x / y, nil
}

func main() {
	quotient, error := divide(5.6, 2)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Printf("%0.2f\n", quotient) // => 2.80
}
