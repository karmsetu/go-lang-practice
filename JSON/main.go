package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"courseName"` //? `json: "courseName"` if you use this, it will throw an issue: [](https://forum.golangbridge.org/t/struct-field-tag-not-compatible-with-reflect-structtag-get-bad-syntax-for-struct-tag-pair/27903)
	Price int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string
}

func main() {
	fmt.Println(`JSON here!`)
	// EncodeJSON()
	DecodeJSON()

}

func EncodeJSON(){
	lcoCourses := []course{
		{"ReactJS", 299, "YT", "abc123", []string{"web dev", "js"}},
		{"NextJS", 199, "FB", "abc123", []string{"web dev", "js"}},
		{"AngularJS", 299, "IG", "abc123", []string{"web dev", "js"}},
		{"VueJS", 299, "X", "abc123", nil},
	}

	// package this data as JSON data
	finalJSON, err := json.MarshalIndent(lcoCourses, "","\t") //? `\t` is tab
	CheckError(err)

	fmt.Printf("%s\n", finalJSON)

}

func DecodeJSON(){
	jsonData := []byte(`{"courseName":"ReactJs","Price":"299","website":"YT","Tags":["web dev","JS"]}`)
	var lcoCourse course 
	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSOn was valid")
		json.Unmarshal(jsonData, &lcoCourse )
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("json is not valid")
	}

	//some cases where you eant to add data to a key

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)

		fmt.Printf("%#v\n", myOnlineData)

	for key, val := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is:%T\n", key, val, val)
	}
}

func CheckError(err error){
	if err != nil {
		panic(err)
	}
}