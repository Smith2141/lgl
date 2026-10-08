package main

import (
	"fmt"

	"github.com/Smith2141/lgl/company"
	"github.com/Smith2141/lgl/person"
	"github.com/Smith2141/lgl/robot"
)

func main() {
	p := person.Person{}
	comp := company.Company{}

	comp.Hire(p) // мы передаём переменную типа Person в функцию, аргументом которой является переменная Worker!

	robot := robot.Robot{Model: "T1000", SerialId: 111, WorkCounter: 0}
	// robot := robot.Robot{"T1000", 111, 0}
	robo := &robot
	comp.Hire(robo)

	tasks := []string{"Задача 1", "Задача 2", "Задача 3"}

	rw := robo.Work(tasks)
	fmt.Println("robot working: ", rw)

	fmt.Println("end")
}
