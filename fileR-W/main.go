package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("hello")
	content := "hello"
	writeFile(content,"./txt.txt" )
	readFile("./txt.txt")
}

func writeFile(content string, location string){
	file, err := os.Create(location)

	if err != nil {
		panic(err) //? shut down program exec
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("lenght is:", length)

	defer file.Close()
}

func readFile(location string){
	dataByte, err := os.ReadFile(location)

	if err != nil {
		panic(err) //? shut down program exec
	}
	fmt.Println(dataByte)
	fmt.Println(string(dataByte))
	
}