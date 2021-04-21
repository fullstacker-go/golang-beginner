package main

import "fmt"

func main() {
	course := 0

	switch course {

	case 0:
		fmt.Println("You selected group 0")
	case 2:
		fmt.Println("You selected group 2")
	case 3:
		fmt.Println("You selected group 3")
	default:
		fmt.Println("You are not among any group")
	}
	var age int

	fmt.Println("Enter your age")
	fmt.Scanln(&age)
	fmt.Println("You chose", age)

	switch {
	case age < 18:
		fmt.Printf("You can Drive after %d years", 18-age)
	case age > 18:
		fmt.Println("You can Drive now")
	case age == 18:
		fmt.Println("..wait....")

	}
}
