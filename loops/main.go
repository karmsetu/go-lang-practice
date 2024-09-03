package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in golang")
	days := []string{"Sunday", "monday", "Tuesday", ""}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println((days[d]))
	} // ? basic 

	for i := range days {
		fmt.Println((days[i]))
	} //? index by range

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}// index and value derived from array

	rougueValue := 1
	for rougueValue < 10 { //? a while loop

		if rougueValue == 2 {
			goto lco //? jumps to the statement
		}

		if rougueValue == 8 {
			break
		}

		if rougueValue == 6 {
			rougueValue++ //! pretty imp, otherwise it will break the program
			continue
		}

		fmt.Println("value is", rougueValue)
		rougueValue++
	}

	lco:
		fmt.Println("jumping...")

}