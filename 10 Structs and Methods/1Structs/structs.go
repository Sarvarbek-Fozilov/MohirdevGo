package main

import (
	"fmt"
	"strconv"
)

type Human struct{
	Name string
	Surname string
	Age  int
}

// CONSTRUCTORS

func Newhuman() *Human{
	n :=new(Human)
	return n
}

// PARAMETRIC CONTRSUCTORS
func NewHumanwithParams(name,surname string, age int) *Human {
	n:= new(Human)
	n.Name =name
	n.Surname =surname
	n.Age =age
	return n
}

func main(){
	h1:= Human{Name: "John", Surname: "Wick",Age:24}
	h2:= Human{Name: "Bernard", Surname: "Emanuel",Age:29}
	data1 := h1.Name + " " + h1.Surname+ " " + " " + strconv.Itoa(h1.Age)
	fmt.Println(data1)
	data2 := h2.Name + " " + h2.Surname+ " " + " " + strconv.Itoa(h2.Age)
	fmt.Println(data2)
	var h3 = new(Human)
	h3.Name ="Sarvarbek"
	h3.Surname ="Fozilov"
	h3.Age =101
	fmt.Println(h3)

	h4 := new(Human)
	fmt.Println(h4)

	h5 := NewHumanwithParams("Asosrbek","Naimjonov",65)
	fmt.Println(h5)


}