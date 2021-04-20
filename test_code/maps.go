package main

import "fmt"

func main() {

	Performance := make(map[string]int)

	Performance["Hindi"] = 65

	studentData := map[string]int{
		"Go": 60,
		"C":  45,
		"JS": 75,
	}

	fmt.Println(Performance)
	fmt.Println(studentData)
	studentData["Go"] = 70     //Updating value of map as key exists in map
	studentData["Python"] = 80 //Inserting values into map as key provided not exists
	fmt.Println(studentData)
	delete(Performance, "Hindi") //Deleting an key from map will also delete the value
	fmt.Println(Performance)
	display(studentData)

}

func display(myMap map[string]int) {
	for i, val := range myMap {
		fmt.Println(i, val)
	}
}
