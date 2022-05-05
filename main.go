package main

import "fmt"

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
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// alex.zipCode = 94000
	// alex.email = "jim@gmail.com"
	alex.updateName("Jimmi")
	alex.print()
	(&alex).print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
