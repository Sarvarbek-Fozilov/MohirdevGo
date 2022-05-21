package main

import "fmt"

func main(){
	//BASIC ARRAY
	myArray :=[3] int{}
	fmt.Println(myArray)
	myArray[0] = 111
	myArray[1] = 222
	myArray[2] = 333
	fmt.Printf("Son 1: %d\n", myArray[0])
	fmt.Printf("Son 2: %d\n", myArray[1])
	fmt.Printf("Son 3: %d\n", myArray[2])

	var  colors [4] string
	fmt.Println(colors)
	colors[0] = "One"
	colors[1] = "Two"
	colors[2] = "Three"
	colors[3] = "Four"

	fmt.Println(colors)
	fmt.Println(colors[3])
	// fmt.Println(colors[4])

	var numbers =[5] int {12,45,23,76,83}
	fmt.Println("Array length: ",len(numbers))
}