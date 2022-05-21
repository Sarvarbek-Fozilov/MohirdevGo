package main

import "fmt"

func main(){
	myArray:= [...] int {11,22,33,44,55,66,77,88,99}
	//9
	fmt.Println(myArray[len(myArray)-1])
	// myArray[9]=14 i cant change

	myCars:=[...] string{"Lombarghini","Ferrari","Mustang"}
	fmt.Println(len(myCars))
	myCars[0]="Porche"
	fmt.Println(myCars)

	yigindi:=0

	for i := 0; i < len(myArray); i++ {
		fmt.Println("Son: ",myArray[i])
		yigindi +=myArray[i]
		
	}
	fmt.Println("Yig'indi: ",yigindi)

}