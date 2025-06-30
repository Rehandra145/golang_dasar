package main

import "fmt"

type Blacklist func(string) bool

//anonymous function adalah function tanpa nama

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

//menggunakan anonymous function
func main() {

	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("Rehan", blacklist)

	//atau bisa juga
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
