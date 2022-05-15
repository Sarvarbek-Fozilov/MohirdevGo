package main

import "fmt"

func main() {
	a := 10
	b := 20
	
	qoshilganda := a+b
	fmt.Println(qoshilganda)

	modda := a%b
	fmt.Println(modda)

	kopayganda := a*b
	fmt.Println(kopayganda)

	jami := 0
	jami += 10  // jami = jami + 10 ---> 10
	jami -= 10  // jami = jami - 10  ---> 0
	jami *= 10  //jami = jami * 10   --> 0
	jami /= 10  // jami = jami / 10 ---> 0
	fmt.Println(jami)

	natija := 10

	natija ++ // natija = natija +1 // natija +=1
	fmt.Println(natija)

	natija-- // natija = natija - 1 // natija -=1
	fmt.Println(natija)

}