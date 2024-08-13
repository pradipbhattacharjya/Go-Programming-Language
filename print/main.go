package main

import "fmt"

func main() {
	age := 25
	name := "Pradip"
	height := 5.5234567

	fmt.Println("age: ", age, "name: ", height, "name: ", name)
	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %.2f\n", height)
	fmt.Printf("height is %T \n", height)
	fmt.Printf("name is %s\n",name)
}