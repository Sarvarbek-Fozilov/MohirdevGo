package main

import (
	"bytes"
	"fmt"
	"runtime"

	"time"
)

func main(){
	runtime.GOMAXPROCS(1)
	go alphabet()
	time.Sleep(time.Second)
}

func alphabet(){
	for l:= byte('a'); l<=byte('j'); l++{
		go fmt.Println(string(l))

	}
}