package main

import "fmt"

func main() {
  x, y := 5, 10

  fmt.Println("Before SWAP:")
  fmt.Println("x =", x)
  fmt.Println("y =", y)
  swap(&x, &y)
  fmt.Println("After SWAP:")
  fmt.Println("x =", x)
  fmt.Println("y =", y)
}

func swap(x *int, y *int) {
  *x, *y = *y, *x
}

/*

package main

import "fmt"

 func main(){
	  son1,son2 := swap (23,12)
	  fmt.Println("Sonlar: ",son1,son2)
 }

 func swap( a,b int)(int , int){
	 return b,a

 }

 */
