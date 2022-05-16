package main

import "fmt"

func main(){
	// Juft yoki toq son aniqlaydigon dastur

	son:=0
	fmt.Println("Juft yoki Toq ekanligini aniqlamoqchi bo'lgan soningizni kiriting")
	fmt.Scanf("%d",&son)

	if son%2 ==0 {
		fmt.Println("Kiritgan soningiz Juft !")
	} else {
		fmt.Println("Kiritgan soningz Toq son !!")
	}
}