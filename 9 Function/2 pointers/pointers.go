package main

import "fmt"

// * Pointers

func main() {
 number := 1
 example(&number)                //  number bu shunchaki o'zgaruvchi uni pastdagi example funksyasida ishlata olmaymiz ,ishlatish uchun esa uning adresini  examplega jonatish orqalik uning qiymatidan foydalanishmiz mumkn
 fmt.Println("Number: ", number) // pastdagi example orqali o'zgartrgan qiymatimizni chiqaradi

 /*
  var p *int
  i := 42
  p = &i
  fmt.Println(p)  // pinterning ramdagi manzilini ekranga chiqaradi
  fmt.Println(*p) // pointerning o'zlashtrgan qiymatini chiqaradi
 */

}

func example(son *int) {
 fmt.Println("Adresi:", son)
 fmt.Println("Son:", *son)
 *son = 23 // Bunda numberni ichga kirib qiymatni o'zgartradi
}