package main

import (
	"fmt"
	"math/rand"
	"strings"
)
func main(){
	//fmt
	fmt.Println("Hello Gophers!!")// fmt  packages import boldi. bu bir marta import bolganda qayta qayta compile bo'lmaydi

	//Rand --> tasodifiy random
	fmt.Println("random son:",rand.Intn(10))

	//Strings

	fmt.Println(strings.Contains("Test","es1")) // kirtlgan qiymat ichdan es1 ni qidradi va bor bolsa ture  yoki false qaytaradi
	fmt.Println(strings.Count("test","t"))//kirtlgan qiymatinng ichda nechta t borlgni sanab beradi
	fmt.Println(strings.HasPrefix("Saud@gmail.com","saud"))// boshidagi strinni qidradi
	fmt.Println(strings.HasSuffix("Meningfayllarim.tar","tar"))//oxrdagi sozni qidradi
	fmt.Println(strings.Index("Hello ","o"))// string ichdagi  qiymatdan ushbu O ni nechanchi indexda ekanligni aytadi
	
}