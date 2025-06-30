package main

import "fmt"

func sayHelloTo(name string, alamat string) {
	fmt.Println("Hello", name, "dari", alamat)
}

func main() {
	sayHelloTo("Rehan", "Agam")
}
