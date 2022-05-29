package main

import (
	"fmt"
	"strconv"
)

type CarFace interface {
	Run() bool
	Stop() bool
	Information() string

}

// * INTERFUNCTION

func CarExecute(c CarFace){
	fmt.Println("Car Info: " + c.Information())

	message :=""
	isRun:=c.Run()
	if isRun{
		message = "Ishlayapti...."
	} else {
		message="Ishlamayapti !!!"
	}
	fmt.Println(message)

	msg :=""
	isStop:=c.Stop()
	if isStop{
		message = "To'xtagan!!."
	} else {
		message="To'xtamagan...."
	}
	fmt.Println(msg)
}

type Car struct {
	Brand, Model, Color string
	Speed               int
	Price               float64
}

type SpecialProduction struct {
	Special bool
}

// * FERRARI STRUCT ---

type Ferrari struct {
	Car
	SpecialProduction
}

func (_ Ferrari) Run() bool{
	return true;
}
func (_ Ferrari) Stop() bool{
	return true;
}
func (f Ferrari) Information() string{
	r:= "Brandi : "  + f.Brand + "\n" + "Rangi:   "+ f.Color + "\n" + "Modeli:  " + f.Model + "\n" + "Tezligi: " + strconv.Itoa(f.Speed) + "\n" + "Narxi:   " + strconv.FormatFloat(f.Price, 'f',-1,64) + "Million $\n"
	a:= "Yes"
	if f.Special {
		fmt.Println("Maxsus ishlab chiqariladimi: ",a)
	}

	return r
}

// *BMW STRUCT ---

type BMW struct {
	Car
	SpecialProduction
}

func (_ BMW) Run() bool{
	return true;
}
func (_ BMW) Stop() bool{
	return true;
}
func (b BMW) Information() string{
	r:= "Brandi : "  + b.Brand + "\n" + "Rangi:   "+ b.Color + "\n" + "Modeli:  " + b.Model + "\n" + "Tezligi: " + strconv.Itoa(b.Speed) + "\n" + "Narxi:   " + strconv.FormatFloat(b.Price, 'f',-1,64) + "Million $\n"
	a:= "Yes"
	if b.Special {
		fmt.Println("Maxsus ishlab chiqariladimi: ",a)
	}

	return r
}




func main() {
	fmt.Println("Ferrari  modelining xarakteristikasi: ")
	f1:=new(Ferrari)
	f1.Brand = "F450"
	f1.Model = "Ferrari 1"
	f1.Color= "Red"
	f1.Speed = 300
	f1.Price=1.2
	f1.Special = true
	// fmt.Println(f1.Information())
	
	fmt.Println("BMW modelining xarakteristikasi: ")

	b1:=new(BMW)
	b1.Brand = "Bumer"
	b1.Model = "M4"
	b1.Color= "Black"
	b1.Speed = 400
	b1.Price=0.5
	b1.Special = false
	// fmt.Println(b1.Information())

	CarExecute(f1)
	CarExecute(b1)


}
