package main

import "fmt"

func main(){
	count:=0

	fmt.Printf("Iltmos sondi kiriting")
	fmt.Scanf("%d",&count)
	for i :=1; i<count; i++{
		if i%2==0{
			fmt.Printf("Juft son: %d\n",i)
		} else  {
			fmt.Printf("Toq son: %d\n",i)
		}
	}
	from:=0
	fmt.Printf("Iltmos sondi kiriting")
	fmt.Scanf("%d",&from)

	for i:=from; i>=0; i--{
		fmt.Printf("Son: %d\n",i)

	}
}