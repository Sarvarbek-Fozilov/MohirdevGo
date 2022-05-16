package main

import "fmt"
func main(){

	/*
	BAHOLASH TIZMIM
	80-100 --> 5
	60-80 ---> 4
	40-60 ---> 3
	40 dan past --> failed
	*/

	 /*  
	     || ->OR YOKi -->  + 

	fmt.Println( true || true) // 1 + 1 = 1
	fmt.Println(true || false) // 1 + 0 =1
	fmt.Println( false || true) // 0 + 1 = 1
	fmt.Println(false|| false) // 0 + 0 =1

	      && --> AND --> VA --> *

	fmt.Println( true && true) // 1 * 1 = 1
	fmt.Println(true && false) // 1 * 0 =0
	fmt.Println( false && true) // 0 * 1 = 0
	fmt.Println(false && false) // 0 * 0 =0

	*/

	var ball int

	fmt.Println("Ballni kiriting: ")
	fmt.Scanf("%d",&ball)

	if ball >100{
		fmt.Println("Tabriklaymiz Grant!!")
	} else if ball >=80 && ball<=100 {
		fmt.Println("Baho: 5 !")
	} else if ball >=60 && ball<80 {
		fmt.Println("Baho: 4 !")
	} else if ball >=40 && ball <60 {
		fmt.Println("baho: 3 !")
	} else {
		fmt.Println("Failed !!!")
	}
}