package main

import "fmt"

func main(){
	countries :=[...] string{"US","UK","UZBEKSITAN"}
	for i, v := range countries{
		fmt.Println("VALUE: ",v)

	}
	numbers := [...]int {11,22,33,44,55,66,77,88,99}
	for _,v :=range numbers  {
		if v%2 ==0{
			fmt.Println("Array qiymatlatri ichidagi juft sonlar: ",v)
		}
	}


}