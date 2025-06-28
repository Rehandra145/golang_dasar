package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var d = 5
	var e = 2
	var c = a + b - d*e
	fmt.Println(c)

	//augmented assignment
	a += 5 // a = a + 5
	fmt.Println(a)
	b -= 2 // b = b - 2
	fmt.Println(b)
	c *= 2 // c = c * 2
	fmt.Println(c)
	d /= 2 // d = d / 2
	fmt.Println(d)
	e %= 2 // e = e % 2
	fmt.Println(e)
}
