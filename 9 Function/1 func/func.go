package main

import "fmt"

// * Functions

func main() {
 sayHi()
 sayMyName("Sarvarbek")
 fmt.Println(add(10, 20, ""))
 natija := add(100, 200, "")
 fmt.Println("natija: ", natija)

}

func sayHi() {
 fmt.Println("Hello !!!")
}

func sayMyName(name string) {
 fmt.Println("Hi", name)
}
func add(n1, n2 int, ism string) int {
 fmt.Println(n1 + n2)
 return n1 + n2
}