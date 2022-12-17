package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	joao := person{
		firstName: "joao",
		lastName:  "alvarez",
		contactInfo: contactInfo{
			email:   "joao@joao.com",
			zipCode: 48307,
		},
	}

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// alex.contactInfo.email = "alex@alex.com"
	// alex.contactInfo.zipCode = 48307

	joao.updateName("john")

	fmt.Printf("%T\n", joao)
	fmt.Println(joao)
	joao.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
