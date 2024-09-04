package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://jsonplaceholder.typicode.com/todos/?q=boi"

func main() {
	fmt.Println("welcome to handleing URLS in go")

	// parsing
	result,_ := url.Parse(myUrl)

	/*
	type URL struct {
	Scheme      string
	Opaque      string    // encoded opaque data
	User        *Userinfo // username and password information
	Host        string    // host or host:port (see Hostname and Port methods)
	Path        string    // path (relative paths may omit leading slash)
	RawPath     string    // encoded path hint (see EscapedPath method)
	OmitHost    bool      // do not emit empty host (authority)
	ForceQuery  bool      // append a query ('?') even if RawQuery is empty
	RawQuery    string    // encoded query values, without '?'
	Fragment    string    // fragment for references, without '#'
	RawFragment string    // encoded fragment hint (see EscapedFragment method)
}
	*/
	qParams := result.Query()
	fmt.Println(qParams["q"])

	partsOfURL := &url.URL{
		Scheme: "https",
		Host: "jsonplaceholder.typicode.com",
		Path: "/todo",
	}

	anotherUrl := partsOfURL.String()
	fmt.Println(anotherUrl)
}