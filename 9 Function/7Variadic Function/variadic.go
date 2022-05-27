package main

import "fmt"
func main(){
	natija:=yigindiHisobla(12,23,34,45,56,67,78,89,)
	fmt.Println("Natija: ",natija)
}

func yigindiHisobla(sonlar ...int)(yigindi int) {
	for _,son:=range sonlar {
		yigindi+=son
	}
	return
}