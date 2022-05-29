package main

import (
	
	"fmt"
	"sync"
)

/*


func summadd(s []int,c chan int) {
	sum :=0
	for _,v:=range s{
		sum+= v
	}
	c <- sum
}

*/

var wg =sync.WaitGroup{}
func main(){
	/*
	s := [] int {1,2,3,4,5,6}
	c:= make(chan int)

	go summadd(s[:],c)
	go summadd(s[: len(s)/2],c)

	x,y := <-c,<-c
	fmt.Println(x,y,x+y)

	*/
	 channel := make(chan string)
	 wg.Add(2)

	 go sortAlphabet(channel)
	 go printAlphabet(channel)

	 wg.Wait()

}

func sortAlphabet(ch chan string){
	for l := byte('a'); l<= byte('z');l++{
		ch<-string(l)
	}
	wg.Done()
	
}
func printAlphabet(ch chan string){
	for l := byte('a'); l<= byte('z');l++{
		fmt.Println(<-ch)
	}
	wg.Done()
}	