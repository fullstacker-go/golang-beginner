package main

import (
	"fmt"
	"os"
)

var (
	course string  = "Go Fundamentals" //This is explicit variable definition we mention type here
	score  int     = 60
	module float64 = 2.8
) //this are package level variables means global variables

const changeFactor = 0.3 //Declaration of constant this value cant be changed by any means.

func main() {
	student := "Tanvir" //This is implicit definition of a variable. Type string is immutable
	fmt.Println(course, score, module, student)
	fmt.Println("\nHi", student, "You are currently watching course", course)
	changeCourse(course)
	fmt.Println("\nHi", student, "You are currently watching course", course)

	changeCoursebyRef(&course)
	fmt.Println("\nHi", student, "You are currently watching course", course)
	result := score + int(module)
	fmt.Println(result)
	//Checking Constant and typecasting
	fmt.Println(changeScore(score, changeFactor))

	//Getting Environment Variables from system
	username := os.Getenv("USERNAME")
	fmt.Println(username)
	//fmt.Println(os.Environ())
	// for _, val := range os.Environ() {
	// 	fmt.Println(val)
	// }
}

//Passing By Values
func changeCourse(course string) string {
	course = "Deep Dive in Go Functions"
	//this value gets changed only for this function and not changes the value of original variable
	fmt.Println("Changing Your course to ", course)
	return course
}

//Passing Values by Reference to mutate a variable
func changeCoursebyRef(course *string) string {
	*course = "Deep Dive in Go Functions"
	//this value gets changed globally
	//because pointer to a location will mutate the value of that location by dereferencing it
	fmt.Println("Changing Your course to ", *course)
	return *course
}

func changeScore(score int, change float64) float64 {
	changedScore := float64(score) * change
	return changedScore
}
