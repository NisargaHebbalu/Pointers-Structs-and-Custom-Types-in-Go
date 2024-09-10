package main

import (
	"fmt"
)

func updateName(x *string) { // here * means we r accepting the pointer value
	*x = "Kate" // here * means we are using it for the dereference of the pointer that we have taken
}

func main() {

	name := "Kely"

	fmt.Println("memory addess: ", &name)

	m := &name // & is used to refer the address

	fmt.Println("memory address: ", m)
	fmt.Println("value stored in memory address: ", *m) // * is used to refer the value in that particular address
	fmt.Println(name)

	//pass by reference
	updateName(m) // updateName(&name)
	fmt.Println(name)
}
