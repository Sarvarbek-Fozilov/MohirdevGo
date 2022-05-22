package main

import (
	"fmt"
	
)

func main(){
	states:= make(map[string]string)
	fmt.Println("States: ", states)
	states["TAS"] = "Tashkent"
	states["KAS"] = "Kashqadaryo"
	states ["AND"] ="ANDIJAN"

	fmt.Println("States: ",states)

	fmt.Println("State one: ",states["TAS"])
	fmt.Println("State two: ",states["XOR"]) // Map ning ichidagi malum bir qiymatni tanlab olish

	sam:= states["KAS"] // states ni ichdagi  kas qiymatini sam degan ozgaruvchiga tenglashtrayapti
    fmt.Println("KAS: ",sam)

	delete(states,"AND") // map ning ichdagi qiymatini o'chirish
	fmt.Println("States: ",states)

	for k,v :=range states{
		fmt.Printf("Key: %v. Value: %v\n",k,v) // state ning  ichga kirib key va value larini ekranga chiqazadi  agar istasak  faqatgina key yoki valuelarini ekranga chiqarsa bo'ladi
		myMap:=make(map[int]string)
		myMap[11] = "O'n bir"
		myMap[22] ="Yigirma Ikki"
		myMap[33] = "O'ttiz Uch"
		myMap[44] = "Qirq"
		myMap[55] = "Ellik Besh"

	keys := make([]string, len(states))

	i:=0
	for _,v= range states {
		keys[i] = v
		i++
	}
	fmt.Println("KEY SLICE: ",keys)

	}
}