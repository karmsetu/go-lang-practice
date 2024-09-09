package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	// "time"
)

var wg sync.WaitGroup

func main() {
// go greeter("Shourya") //? `go`
// greeter("World")
websiteList := []string{
	"https://www.golang-book.com/",
	"https://www.youtube.com",
	"https://www.google.co.in/",
	"https://github.com/",
}

for _, web := range websiteList {
	go getStatusCode(web)
	wg.Add(1)
}

wg.Wait()

}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		// time.Sleep(3 * time.Millisecond)
// 		fmt.Printf("Hello %v!", s)
// 	}
// }

func getStatusCode(endpoint string)  {
	defer wg.Done()
	result, err := http.Get(endpoint)
	CheckIsError(err)
	fmt.Printf("%d status code for %s \n", result.StatusCode, endpoint)
}

func CheckIsError(err error){
	if err != nil {
		log.Fatal(err)
	}
}


