package main

import "fmt"

//di golang, function adalah first class citizen
//ini berarti function bisa disimpan di variable, bisa dikirim sebagai parameter, dan bisa mengembalikan function lain
//berarti function itu sendiri bisa dikatakan sebagai tipe data

func getGoodBye(name string) string {
	return "Good Bye " + name
}
func main() {
	goodBye := getGoodBye
	misal := getGoodBye

	fmt.Println(goodBye("Rehan"))
	fmt.Println(misal("Agam"))
	fmt.Println(getGoodBye("Syarifah"))
}
