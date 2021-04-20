package main

import "fmt"

type courseMeta struct {
	author string
	level  string
	course string
	score  int
	rating float64
}

func main() {

	//var goFundamentals courseMeta

	//goFundamentals := new(courseMeta)

	goFundamentals := courseMeta{ //Creating new instance of type courseMeta
		author: "Tanvir Shaikh",
		level:  "Beginner",
		score:  9,
		rating: 4.5,
	}
	fmt.Println("The author of this course is", goFundamentals.author)
	modify(&goFundamentals) //referencing gofundamentals as pointer
	fmt.Println(goFundamentals)
	//Accessing single value from a struct courseMeta
	fmt.Println("The author of this course is", goFundamentals.author)

}

func modify(myStruct *courseMeta) {
	myStruct.author = "Peter Parker"
}
