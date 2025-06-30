package main

import "fmt"

func getHello(name string) string { // string disini berarti penanda dia bakalan return string
	return "Hello" + name
}
func main() {

	//ini disimmpan ke variable dulu karena di function getHello itu dia return string
	//bukan langsung print
	//jadi harus disimpan dulu ke variable baru bisa di print
	result := getHello("Rehan")
	fmt.Println(result)

	//atau bisa langsung dipanggil tanpa disimpan ke variable
	fmt.Println(getHello("Rehan"))
}
