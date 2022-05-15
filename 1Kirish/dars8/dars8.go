//INPUT haqida
/* Output biz ekranga chiqarayotgan narsalarga aytiladi.
foydalanuvchidan qiymat olib  uni 
package main o'zgaruvchiga tenglashtirishimizga  INPUT  deymiz,*/

import "fmt"
        "bufio"
		"os"
		"strconv"


func main(){
	println("HELLO GOLANG") //OUTPUT

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Son1 kiriting: ") // Ekranga Son kiriting yozuvi chiqadi undan kegin
	// readerni ReadString metodi  orqalik foydalanuvchidan  qandaydur qiymat olib uni quyidagi "son" degan o'zgaruvchiga tenglab qo'yadi. 
	son1,_ := reader.ReadString('\n')
	fmt.Println("Son2 kiriting: ")
	son,_ := reader.ReadString('\n')
	fmt.Println( son1 + son2 )

	// typeni o'zgartrib  ozgaruvchiga trnglash

	fmt.Println("Number Kiriting")
	number,_ := reader.ReadString('\n') // qiymat foydalanuvchi tomonidanm kiritiladi 
	 //ReadString metodi orqalik bu qiymat number ga  beriladi. bu qiymatni esa kegin.
	value,_ := strconv.ParseFloat(strings.TrimSpace(str),64)	// Parse orqalik float  turiga o'zgartramiz.
	//trimpsace  ni  bu yerda foydalanuvchi  qiymat kirtayotganda tab tugamasini  bosib  
	//bir necha  son kiritsa error qaytarwi  mumkin  shu vaqtda "bo'shqiymat" orqalik muammo xall boladi.	
	fmt.Println(value +100)

	//foydalanuvchidan qiymat olmoqchi bo'lsek  scanln ham bor ularni kegin o'rganamiz.

 }