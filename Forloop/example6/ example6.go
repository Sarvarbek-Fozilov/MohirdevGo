package main

import "fmt"
func main(){
	//karra jadvali
	 for i := 1; i < 10; i++ {
		 for j:= 1; j < 10; j++ {
			 fmt.Printf("%d X %d = %d\t",j,i,j*i)
			 
		 }	 
		 fmt.Printf("\n")
	 }
}