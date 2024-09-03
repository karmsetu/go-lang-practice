package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://127.0.0.1:5500/web_requests/index.html" //! make sure to use live server 

func main() {
	
	fmt.Println("LCO web request")
	response, err := http.Get(url)
	if(err != nil) {
		panic(err)
	}
	
	fmt.Println(response)

	defer response.Body.Close() //! very imp
	dataBytes, err :=ioutil.ReadAll(response.Body)
	if(err != nil) {
		panic(err)
	}
	fmt.Println(string(dataBytes))
}

