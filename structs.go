package main

import "fmt"

func main() {
	person := person{
		firstName: "Eleandro",
		lastName: "Duzentos",
	}
	
	fmt.Println(person.getFullName())
	
	person.printLifeStatus()
	
	person.kill()
	
	person.printLifeStatus()
}

type person struct {
	firstName string
	lastName string
	isDead bool
}

func (p person) getFullName() string {
	return p.firstName + " " + p.lastName;
}

func (p *person) kill() {
	p.isDead = true
}

func (p person) printLifeStatus() {
	fmt.Println(p.isDead)
}
