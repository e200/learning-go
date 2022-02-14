package main

import "fmt"

func main() {
	//enGreeter := enGreeter{}
	ptGreeter := ptGreeter{}
	
	greet(ptGreeter)
}

type greeter interface {
	greet(name string)
}

type enGreeter struct {}
type ptGreeter struct {}

func (enGreeter) greet(name string) {
	var innerName string = "World!"
	
	if (name != "") {
		innerName = name
	}
	
	fmt.Println("Hello " + innerName)
}

func (ptGreeter) greet(name string) {
	var innerName string = "World!"
	
	if (name != "") {
		innerName = name
	}
	
	fmt.Println("Olá " + innerName)
}

func greet(g greeter) {
	greeter.greet(g, "")
}
