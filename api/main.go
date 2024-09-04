package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const PORT = 8000

func main() {
	
	PerformGetRequest("/")
	PerformGetRequest("/getit")
}

func PerformGetRequest(path string) {

	anotherUrl := &url.URL{
		Scheme: "http",
		Host: "localhost:8000",
		Path: path,
	};
	

	response, err := http.Get(anotherUrl.String())
	CheckError(err)
	
	fmt.Println("Response is:",response.Status)

	dataByte, err :=io.ReadAll(response.Body)
	CheckError(err)

	var responseString strings.Builder

	responseString.Write(dataByte) 
	fmt.Println(string(dataByte))
	fmt.Println(responseString)

	defer response.Body.Close()
}

func CheckError(err error){
	if err != nil {
		panic(err)
	}
}