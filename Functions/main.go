package main

import "fmt"

func simpleFunction() {
	fmt.Println("simple function")
}

func add(a, b int) (int) {
	return a + b
}
func multiply(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("we are learing function in Golang")
	simpleFunction()
	ans := add(3, 4)
	fmt.Println("add of two number is :", ans)

	data := multiply(3, 4)
	fmt.Println("multiply of two number is :", data)
}
