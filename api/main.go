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
	
	// PerformGetRequest("/")
	// PerformGetRequest("/getit")

	// PerformPostJSONRequest("/post")

	PerformPostFormRequest("postform")
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

func PerformPostJSONRequest(path string){
url := &url.URL{
		Scheme: "http",
		Host: "localhost:8000",
		Path: path,
	};

	// fake json payload
	requestBody := strings.NewReader(`
	{
		"Name":"Boi",
		"toBe": "Good"
	}
	`)

	response, err :=http.Post(url.String(), "application/json", requestBody)
	CheckError(err)

	content, err := io.ReadAll(response.Body)
	CheckError(err)

	fmt.Println(string(content))

	defer response.Body.Close()
}

func PerformPostFormRequest(path string) {
	myUrl := &url.URL{
		Scheme: "http",
		Host: "localhost:8000",
		Path: path,
	};

	data := url.Values{}
	data.Add("firstName", "Boi")
	data.Add("lastName", "gajodhar")
	data.Add("age", "69")

	response, err :=http.PostForm(myUrl.String(), data)
	CheckError(err)

	content, _:= io.ReadAll(response.Body)

	defer response.Body.Close()
	fmt.Println(string(content))
}

func CheckError(err error){
	if err != nil {
		panic(err)
	}
}