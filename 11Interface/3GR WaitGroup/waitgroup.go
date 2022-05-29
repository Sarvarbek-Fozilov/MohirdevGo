package main

import (
	"fmt"
	"sync"
)

var wg=sync.WaitGroup{}
func main(){
	wg.Add(2)
	go salomBer()
	go korishgunchaDe()
	wg.Wait()
	
}

func salomBer(){
	fmt.Println("Salom!!")
	wg.Done()
}

func korishgunchaDe(){
	fmt.Println("Ko'rishguncha !!")
	wg.Done()
}