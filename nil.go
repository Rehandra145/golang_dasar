package main

import "fmt"

//nil hanya bisa digunakan pada map, slice, pointer, channel, interface, struct
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("Rehan")

	if data == nil {
		fmt.Println("data map masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}
