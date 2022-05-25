package main

import "fmt"

func main(){
	message := "First Message"
	changeMessage(&message)
	fmt.Println(message)
}

func changeMessage(msg *string){
	*msg = "Second Message"
}