package main

import "fmt"

func main() { //? main acts as entry point to the program
	fmt.Println("main function")

	greeter()
	// func greeter2()  {//? can't write function in a function
	// fmt.Println("Another method")
	// }

	data := add(2,1)
	fmt.Println(data)

	manyData := proAdder(1,2,3,4,5,6,7,8)
	fmt.Println(manyData)

}

//? yes you can create anonymous or ify funcs

func add(num1 int, num2 int) int{
	return num1 + num2
}

func proAdder(values ...int) int { //? `...` are variadic funcitons
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func greeter() {
	fmt.Println("Namastey from golang")
}