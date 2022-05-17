package main

import "time"

func main() {
	 i:=0
	 sum:=0

	 for i< 10 {
		 println("I: ",i)
		 sum += i
		 i ++
		 time.Sleep(500 * time.Millisecond)
	 }
	 println("YIG'INDI: ",sum)
}