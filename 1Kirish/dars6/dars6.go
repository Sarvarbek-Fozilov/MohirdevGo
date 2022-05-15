//Type Changing

package main 

import "strconv"

func main() {
	s:="1"   //s bu yerda  string  turiga kiradi.
    /* (  println(s + 10) ) bunday xolatda  kod ishlamaydi. 
	sababi  s  string  10 esa int  turga kiradi  
	ularni  ishlatish uchun esa ikkkisini ham int  turiga o'tkazish kerak
	 s ni intga o'tkazish uchun strconv degan kutubxonani ishlatib  
	 s ni boshqa bir nomga "Atoi "bilan tenglab olamiz. */

	 raqam,_ := strconv.Atoi(s)
	 println( raqam + 10)
	 // int  turini  float   turiga o'zgartrish  ancha  osonroq
	 n:=11
	 println(float32(n))

	 // Aksincha float qiymatini int turiga o'zgartrishda esa uint   deb nomlangan funksyadan foydalanadi.

	 w := 12.9
	 println(uint(w))  //UNSIGNED INT
}