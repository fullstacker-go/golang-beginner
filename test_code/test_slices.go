package main

import "fmt"

func main() {
	// mySlice := make([]int, 2, 4)
	mySlice := []int{3, 5, 7, 9, 12, 15, 18, 20}
	myCourses := []string{"docker", "python", "go"}
	fmt.Printf("the lenth of the mySlice is %d and capacity is %d", len(mySlice), cap(mySlice))
	fmt.Printf("\nthe lenth of the myCourses is %d and capacity is %d", len(myCourses), cap(myCourses))

	sliceOfslice := mySlice[2:6]

	fmt.Println("\n", sliceOfslice)
	fmt.Printf("\nthe lenth of the sliceOfslice is %d and capacity is %d", len(sliceOfslice), cap(sliceOfslice))

	demoSlice := []int{}

	for i := 0; i < 17; i++ {
		demoSlice = append(demoSlice, i)
		fmt.Println("\n", i)
		fmt.Printf("\nThe Capacity of demoSlice is %d", cap(demoSlice))
	}
}
