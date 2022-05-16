package main

import "fmt"

func main(){
	//Bi biriga qoldiqsiz  bo'linishi yoki bo'linmasligini aniqlaydigon dastur yozish.
    kattaSon, kichikSon := 0,0
	fmt.Println("Katta sondi kiriting:")
	fmt.Scanf("%d" , &kattaSon)
	fmt.Println("Kichik sondi kiriting:")
	fmt.Scanf("%d", &kichikSon)

	if kattaSon%kichikSon == 0 {
		fmt.Println("Bo'linadi !!!")
	} else {
		fmt.Println("Bolinmaydi !!!!")
	}
}