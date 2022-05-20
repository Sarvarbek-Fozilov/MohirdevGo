package main

import "fmt"
func main(){
	son:=0
	fmt.Print("Iltimos Son Kiriting")
	fmt.Scanf("%d",&son)

	yigindi:=0
	faktorial:=1
	for i :=1; i <= son; i++ {
		yigindi +=i
		faktorial *=i
		
	}
	fmt.Printf("Yig'indi: %d\n",yigindi)
	fmt.Printf("Faktorial: %d\n",faktorial)
}