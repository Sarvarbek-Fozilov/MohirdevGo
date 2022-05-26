package main

import "fmt"

func main(){
	ism1,ism2 := swap("Ahmad","Boburjon")
	fmt.Println("Ismlar:",ism1,ism2)
}

func swap(a,b string)(string,string){
	return b,a
}