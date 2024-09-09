package main

import (
	"fmt"
	"net/http"

	"github.com/karmsetu/maongo-api/router"
)

func main() {
	fmt.Println("MONGO")
	r := router.Router()
	fmt.Printf("http.ListenAndServe(\":4000\", r): %v\n", http.ListenAndServe(":4000", r))
	// fmt.Println("")
}