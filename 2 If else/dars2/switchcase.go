package main

import "fmt"
func main(){
    //Switch
	/*fmt.Println("1.Balans\n2.SMS\n3.MB")
	var tanlov int
	fmt.Println("Tanlang")
	fmt.Scanf("%d",&tanlov)
	switch tanlov{
	case 1:
		fmt.Println("Balansigizda: 34$ mavjud!")
		break
	case 2:
		fmt.Println("Sizda 6900ta sms mavjud!")
		break
	case 3:
		fmt.Println("Sizda 7896 mb mavjud!")
		break
	default:
		fmt.Println("Notogri amal tanlandi")
					
	} */
	var ism string
	fmt.Println("Ismingizni kiriting!")
	fmt.Scanf("%v",&ism)
	switch ism {
	case "Anvar":
		fmt.Println("Salom Anvar!")
	case"behruz":
		fmt.Println("Salom Behruz !")
	case "Temur":
		fmt.Println("Salom Temur !")
	default:
		fmt.Println("Bunaqa Ism Mavjud Emas !")				
	}
}