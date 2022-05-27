package main

import "fmt"

func main() {
	fmt.Println("Natija: ",add(11,22))

	f1 := example("Sarvarbek")
	f1()

	example("John")()

}

var add = func(n1,n2 int)(result int) {
	result =n1 +n2
	return
}

var example = func(name string) func() {
	return func() {
		fmt.Println("Salom",name)
	}
}