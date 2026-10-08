package main

import (
	"fmt"

	. "github.com/Smith2141/lgl/company"
	. "github.com/Smith2141/lgl/person"
	. "github.com/Smith2141/lgl/robot"
)

func main() {
	var person Person = Person{Name: "Tom"}
	var company Company = Company{}
	var robot Robot = Robot{Model: "T1000", SerialId: 111, WorkCounter: 0}
	var robo *Robot = &robot

	var tasksW1 = []string{"Задача 1", "Задача 2", "Задача 3"}
	var tasksW2 = []string{"Задача 4", "Задача 5", "Задача 6"}

	company.Hire(person) // мы передаём переменную типа Person в функцию, аргументом которой является переменная Worker!
	company.Hire(robo)

	var roboWorking = robo.Work(tasksW1)
	fmt.Println("robot working: ", roboWorking)

	var personWorking = person.Work(tasksW2)
	fmt.Println("human working: ", personWorking)

	fmt.Println("end")
}
