package main

import "fmt"
func main(){
	//Fibonachi sonlari
	son:=0
	fmt.Println("Iltmos on kiriting")
	fmt.Scanf("%d",&son)
	  x,y,z :=1,1,0
	 for i := 0; i < son; i++ {
		 z=x+y
		 x=y
		 y=z
		 fmt.Printf("%d ",z)
		 
	 }
}