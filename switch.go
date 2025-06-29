package main

import "fmt"

func main() {
	name := "Rehan"

	switch name {
	case "Rehan":
		fmt.Println("Hello Rehan")
	case "Abit":
		fmt.Println("Hello Abit")
	case "Aqul":
		fmt.Println("Hello Aqul")
	default:
		fmt.Println("Hi, boleh kenalan?")
	}

	//switch dengan short statement
	switch len := len(name); len > 5 {
	case true:
		fmt.Println("Nama kamu terlalu panjang")
	case false:
		fmt.Println("Nama kamu pas")
	default:
		fmt.Println("Tidak ada nama")
	}

	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Nama kamu terlalu panjang")
	case length < 5:
		fmt.Println("Nama kamu terlalu pendek")
	case length == 5:
		fmt.Println("Nama kamu pas bang")
	default:
		fmt.Println("Gak ada nama")
	}
}
