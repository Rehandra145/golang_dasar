package main

import "fmt"

//ga pake rekursive
func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

//pake rekursive
func recursiveFactorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursiveFactorial(value-1)
	}
}

func main() {
	fmt.Println(recursiveFactorial(20))
	fmt.Println(factorialLoop(5))
}
