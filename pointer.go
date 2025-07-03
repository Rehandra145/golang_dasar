package main

import "fmt"

type Address struct {
	City, Province, Region string
}

func main() {
	//pass by value
	// address1 := Address{"Padang", "Sumatra Barat", "Indonesia"}
	// address2 := address1 //copied

	//pass by reference
	address1 := Address{"Padang", "Sumatra Barat", "Indonsia"}
	address2 := &address1

	address1.City = "Bukittinggi"
	fmt.Println(address1)
	fmt.Println(address2)
}
