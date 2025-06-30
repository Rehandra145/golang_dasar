package main

import "fmt"

//bisa juga func di param nya itu di aliaskan menggunakan type declaration
//type Filter func(string)string

//bentuk fungsi ini bisa diubah menjadi
//func sayFillteredHello(name string, filter Filter){} => hasilnya akan sama saja
func sayFillteredHello(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Halo", filteredName)
}

func filter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	filterName := filter

	sayFillteredHello("Rehan", filterName)
	sayFillteredHello("Anjing", filter)
}
