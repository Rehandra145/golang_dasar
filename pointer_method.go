package main

/*
	Sama seperti sebelumnya, karna ini value yang di kirim lewat parameter itu bersifat pass by value,
	maka di parameter di functionnya itu harus menggunakan pointer jika value yang disebut di parameter itu
	ingin diubah.
	untuk method ini sangat disarankan untuk selalu memakai pointer
*/

import "fmt"

type Man struct {
	Name string
}

func (man *Man) married() {
	man.Name = "Mr." + man.Name
}
func main() {
	pria := Man{"Rehan"}
	pria.married()

	fmt.Println(pria)
}
