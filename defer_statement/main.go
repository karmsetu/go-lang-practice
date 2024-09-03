// ? https://go.dev/tour/flowcontrol/12
package main

import "fmt"

func main()  {
	defer fmt.Println("world") //? LIFO:  this statement goes to the end of the program
	fmt.Println("hello")

	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three") //? LIFO

	myDefer()

}

func myDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}