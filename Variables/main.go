package main

import "fmt"

func main() {
	var username string = "pradip"
	fmt.Println(username)
	fmt.Printf("Variable is of type:  %T  \n", username)


	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type:  %T  \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type:  %T  \n", smallVal)

	var smallFloat float64 = 255.45556255567885
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type:  %T  \n", smallFloat)

	//default values and some aliases

	var anotherVariable int 
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)
}
