package main

import "fmt"
func main(){
	son1:=12
	son2:=10
	var p=&son1
	fmt.Println(son1==son2)
	fmt.Println(son2==*p)

}
