package main

import "fmt"
func main(){
	//siz kirtgan sondan 0gacha bo'lgan sonlar orasda 4ga bo'linadigonlarni
	//ekranga chiqaradigon dastutr yozing.

	son:=0
	fmt.Printf("Sonni kiriting: ")
	fmt.Scanf("%d",&son)

	for ; son>0; son-=1{
		if son%4==0{
			fmt.Printf("Bo'linadi: %d\n",son)
			
		} else {
			fmt.Println("To'gri butun son kirting")

		}
	}
}