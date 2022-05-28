package main

// DEFER

import "fmt"
  
var isConnected=false



func main(){
	fmt.Println("Connection status: ",isConnected)
	dbProcessing()
	fmt.Println("Connection status: ", isConnected)
}

func dbProcessing(){
	connect()
	fmt.Println("Deffering disconnected !")
	defer disconnect()
	fmt.Println("Connnection status: ",isConnected)
	fmt.Println("BLAH BLAH BLAH !!!")
}

func connect(){
	isConnected =true
	fmt.Println("Connected to Database !!")
}

func disconnect(){
	isConnected=false
	fmt.Println("Disconnected !!!")
}

