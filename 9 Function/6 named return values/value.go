package main

import "fmt"

func main(){
	// * NAMED RETURN VALUES
	bir , ikki :=split(17)
	fmt.Println("Bir: ",bir , "Ikki: ",ikki)
	l,r := sum(11,22,33,44,55,66)
	fmt.Println("Lenght: ",l,"Result: ",r)

	
}

func split(sum int)(x,y int){
	x = sum +10
	y = sum -x
	return
}

func sum(nums ...int)(length,result int){
	result = 0
	for _, num := range nums {
		result += num
	}
	length = len(nums)
	return
}