package main

import "fmt"

//type assertion itu any handler, mengubah tipe data sesuai dengan value yang di berikan
func random() any {
	return false
}
func main() {
	result := random()
	//manual
	// resultString := result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("integer", value)
	default:
		fmt.Println("unknown", value)
	}
}
