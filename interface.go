package main

import "fmt"

/*
Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
Sebuah interface berisikan definisi definisi method
Biasanya digunakan sebagai kontrak
*/

type HasName interface {
	GetName() string //kontrak
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

//fungsi yang menggunakan kontrak harus lah implementasi dari si interface
func SayHello(hasName HasName) {
	fmt.Println("Hallo", hasName.GetName())
}

//implementasi interfac di struct person
func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}
func main() {
	person := Person{
		"Person",
	}

	animal := Animal{
		"kucing",
	}

	SayHello(person)
	SayHello(animal)

	fmt.Println(person)
}
