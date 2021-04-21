package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	go func() {

		defer wg.Done() //defering function execution untill first one is done
		fmt.Println("Entering Goroutine One....")
		time.Sleep(5 * time.Second)
		fmt.Println("Executed One successfully..")

	}()
	wg.Add(1) //adding function to wait group
	go funct2()
	wg.Wait() //waiting till execution of above function
}

func funct2() {

	fmt.Println("Entering Goroutine Two....")
	fmt.Println("Executed Two successfully..")

}
