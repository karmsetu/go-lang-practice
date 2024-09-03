package main

import "fmt"

type User struct {
	Name   string //? capital letters signifies that the prop is exportable
	Email  string
	oneAge int //! non returnable(private)
	Status bool
}

func main() {
	boi := User{Name: "boi", Email: "boi@h.com", oneAge: 4, Status: true}
	fmt.Println(boi)
	boi.NewMail("boi")
	fmt.Println(boi.Email)
	boi.GetStatus()
}

func (u User) GetStatus() {
	fmt.Println("Is user avtive: ", u.Status)
}

func (u User) NewMail(newEmail string){
	u.Email = newEmail
	fmt.Println("Email of this user is: ", u.Email)
} //! problem with this func is that u is the copy of the User struct not the original hence, it doesn't change the boi var