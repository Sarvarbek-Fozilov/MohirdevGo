//Varieble  turlari  declare and  assigment  nmaligi bilan tanishamiz.

package main

func main() {
	//VAR TYPE -> STRING, NUMBER, BOOLEAN
	var habar string = "Salom O'quvchi  habar 1" // DECLARE ->> O'zgaruvchi  bir qatorda berilsa
	/* o'qilishi  habar deb nomlangan o'zgaruvchi uning turi  string  qiymati esa  "salom O'quvchi habar 1"*/
	var habar1 string
	habar1 = "Salom O'quvchi habar 2" // ASSIGMENT -->> O'zgaruvchi nomi va turi alohida qiymati   Alohida berilsa.
	println(habar)
	println(habar1)

	var xat = "Bu Xat O'zgaruvchisi" // O'zgaruvchilarni turini keltirmasdan ham yozish mumkin  bunda ham kod ishlaydi. U o'zi  avtomatik  turini  aniqlab oladi.
	println(xat)

	var a, b, c int  // BUTUN  SON -->. 3,5,15,78,78347834.-5,-8
	println(a, b, c) /* O'zgaruvchini  qiymatini  kiritmasdan ham ishlatsa bo'ladi .
	bu vaqtda  agar u string  turida bo'lsa hech nima qaytarmaydi  agar  int  turida bo'lsa ekranga 0  qiymati chiqadi */
	var x, y, z float32 // QOLDIQ SON -->> 34.23,2.8,-13.6,3345454.5
	println(x, y, z)

	var s1, s2 string = "Udevs", "Epam" /// Bir nechta o'zgaruvchilarni bir vaqtda  ishlatish mumkin,  bunda s1 ga Udevs  qiymat  berilayapti  s2 ga Epam qiymati.
	println(s1, s2)

	var n1, n2, n3 int = 11, 22, 33
	println(n1 + n2 + n3) // int  turida ham bir vaqtni o'zida bir qancha o'zgaruvchi chaqirib ishlatish mumkin.

	var f1, f2, f3 float64 = 11.2, 22.1, 33.3 // qoldiq  sonlarda hisoblash vaqtda eng tog'ri javobni chiqarib beradi.
	println(f1 + f2 + f3)

	// BOOLEON//

	var erkakMi bool = true
	// booleon da ikkta natija  bor  true  yoki false
	println(erkakMi)

	var q, w, e, r = 11, 22.8, "salom", true // KARMA  yoki  MIXIN
	/*  bir vaqtni o'zida etibor bergan bo'lsez 4 xil turdegi o'zgaruvchini kiritdik
	q= 11 ga ( int) w= 22.8 (fload) e="salom" (string) r = true (booleon)
	manashunday mixin  farmatida ham ishlatsa bo'adi.
	Bita satrda o'zgaruvchini turidan qatiy nazar foydalanish mumkin.*/
	println(q, w, e, r)

	//Yanada oson yo'li //

	/* O'zgaruvchi kirtiayotganda uni turin ham var  deb nomlanga kalit so'zini ham ishlatmasdan foydalansa bo'ladi.
	var  raqam int = 44  bilan  raqam := 44  bir xil manoga ega
	bunda go tili o'zi uning type ni va  varekanligi  := bilan sezib oladi.
	:= bu usluda KARMA  xolatida ham ishlatsab   bo'ladi.
	*/

	raqam := 44
	println(raqam)

	/* func maindan avval  o'zgaruvchi  chaqirishimiz mumkin faqatgina sharti uni  quyidagi  xolatdagini ihslatsih mumkin
	var name string = "Sarvarbek"
	 uni qisqa usulda ishlatib bo'lmaydi var ni ham qiymatni ham type ni ham keltrish kerak  bu o'zgarmas qoida
	  bunday  func maindan oldin yozish   "GLOBAl"  deb nomlanadi. */

	//agar berilayotgan qiymat  kop bolsa  unda quydagicha ham  ishlatsa  bo'ladi.
	var  name1,name2,name3 = "Hello world hello world","lorem ipsum  saidabdullah","b","c"
	// yuqoridagi holatda ishlatilsa kod juda ham hunuk  xolatga kelib qoladi
	// buni quyidagicha ishlatsa bo'ladi.

	var (
		ism1 = "a"
		ism2 = "b"
		ism3 = "c"
		ism4 = "f"

	)
	println( ism1,ism2,ism3,ism4)
}
