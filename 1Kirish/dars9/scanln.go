// SCANLN  va TIME  package lari haqia   tanishamiz.

package main

import (
	"fmt"
	"time"
)

func main(){
	 var ism string // Declare
	 /* Scanln  deb nomlangan funksiya input  hisoblanib  foydalanuvchidan  malumot qabul qilib olishda ishlatiladi.*/


	 fmt.Scanln(&ism)
	 println("Salom",ism)
	 // yuqoridagini O'qilishi scanln orqalik foydalanuvchidan input qabul qilib olib,
	 // uni  "ism" ga  tenglab ekranga  chiqarayapmiz.

	 t := time.Date(2022,time.December,13,22,33,45,678,time.UTC)
	 fmt.Println(t)
	 // yuqoridagini o'qilishi, Date ni ichda( avval yili,oy,kun,soat,minut,sekund,milli secund)
	 
	 fmt.Println(time.Now()) // Xozrgi  vaqtni  ekranga chiqarib beradi/.
}