package main

import "fmt"

func main() {
	fmt.Println("We are learning Array in Golang")

	// var name[7]string

	// name[0] = "pradip"
	// name[2] = "subha"

	// fmt.Println("Names of Person is :", name)
    // fmt.Println("Length of name array is :", len(name))
	// fmt.Println("value of name at 2 index is :", name[2])


	// var numbers = [5]int{1, 2, 3, 4, 5}
	// fmt.Println("Number is :", numbers)

	// fmt.Println("Length of NUmbers array is :", len(numbers))

	var name[5]string
	name[2] = "Subha"
	name[0] = "Pradip"
	fmt.Println("name is :", name)
	fmt.Printf("name Array is %q\n", name)

}