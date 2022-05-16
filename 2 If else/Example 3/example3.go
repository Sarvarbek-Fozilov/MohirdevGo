package main

import "fmt"

func main() {
	//HARBIY QABUL
	//Jinsi Erkak , Bo'y 180 dan baland
	//Jinsi Ayol, B'yi 160 dan baland
	//output: Salom ism, Siz  jinsdasiz, Bo'yingiz boy sm, Qabul qilindingiz!
	//output: Salom ism, Siz  jinsdasiz, Bo'yingiz boy sm, Qabul qilinmadingiz!
	var ism,jins,boy = "","",0
	fmt.Println("iltimos Ismingizni yozing")
	fmt.Scanln(&ism)
	fmt.Println("Iltimos jinsingizni kiriting!")
	fmt.Scanln(&jins)
	fmt.Println("Iltimos Bo'yingizni kiriting!")
	fmt.Scanln(&boy)

	if jins == "erkak"  && boy >=180 {
		fmt.Printf("Salom %v. Siz  %v.'siz, Bo'yingiz %v sm,\n ", ism, jins, boy)
		fmt.Println("Qabul Qilindingiz !")
	} else if jins == "erkak"  && boy < 180 {
		fmt.Printf("Salom %v. Siz  %v.'siz, Bo'yingiz %v sm,\n ", ism, jins, boy)
		fmt.Println("Qabul Qilinmadingiz !")
	} else if jins == "ayol"  && boy >= 160 {
		fmt.Printf("Salom %v. Siz  %v.'siz, Bo'yingiz %v sm,\n ", ism, jins, boy)
		fmt.Println("Qabul Qilindingiz !")
	} else if jins == "ayol"  && boy < 160 {
		fmt.Printf("Salom %v. Siz  %v.'siz, Bo'yingiz %v sm,\n ", ism, jins, boy)
		fmt.Println("Qabul Qilinmadingiz !")
	} 	
}