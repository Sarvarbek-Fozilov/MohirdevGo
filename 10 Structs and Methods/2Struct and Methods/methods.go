package main

import "fmt"

type User struct {
	ID, Age                 int
	Name, Surname, Username string
	Pay                     *Payment
}
// * Struct
type Payment struct {
	Salary, Bonus float64
}

//* METHOD
func(p Payment) GetSalaryWithBonus() float64{
	return p.Salary + p.Bonus

}

//*CONSTRUCTOR
func NewUser()*User{
	var n = new(User)
	return n
}

//*METHODLAR
func(u User) GetFullName() string{
	return u.Name + " " + u.Surname
}

func main() {
	u1:=User{
		ID: 1,
		Age:34,
		Name:"John",
		Surname: "Wick",
		Username:"JWick",
		Pay: &Payment{Salary: 1000,Bonus: 300},
	}
	fmt.Println(u1.Age)
	fmt.Println(u1.GetFullName())

	fmt.Println(u1.Pay.GetSalaryWithBonus())

}
