package main

import (
	"fmt"

	"github.com/Smith2141/lgl/company"
	"github.com/Smith2141/lgl/person"
)

func main() {
	p := person.Person{}
	comp := company.Company{}

	comp.Hire(p) // мы передаём переменную типа Person в функцию, аргументом которой является переменная Worker!

	fmt.Println("end")
}
