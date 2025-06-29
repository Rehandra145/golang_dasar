package main

import "fmt"

func main() {
	nama := "abit"

	if nama == "Rehan" {
		fmt.Println("Halo Rehan")
	} else if nama == "Rehan" {
		fmt.Println("Halo Rehan")
	} else {
		fmt.Println("Hi, boleh kenalan?")
	}

	//short statement
	if length := len(nama); length > 5 {
		fmt.Println("Nama kamu terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
