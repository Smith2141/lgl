package main

import (
	"fmt"
)

// Person — структура, описывающая человека.
type Person struct {
	Name string
	Year int
}

// NewPerson возвращает новую структуру Person.
func NewPerson(name string, year int) Person {
	return Person{
		Name: name,
		Year: year,
	}
}

// String возвращает информацию о человеке.
func (p Person) String() string {
	return fmt.Sprintf("person string: Имя: %s, Год рождения: %d", p.Name, p.Year)
}

// Print выводит информацию о человеке.
func (p Person) Print() {
	// вызовется метод String() для Person
	fmt.Println("person print: ", p)
}

// Student описывает студента с использованием вложенной структуры Person. То есть структура Student описывает.
type Student struct {
	Person // вложенный объект Person
	Group  string
}

func NewStudent(name string, year int, group string) Student {
	return Student{
		Person: NewPerson(name, year), // Явно создаём структуру Person
		Group:  group,
	}
}

// String возвращает информацию о студенте.
func (s Student) String() string {
	return fmt.Sprintf("student string: %s, Группа: %s", s.Person, s.Group)
}

func main() {
	s := NewStudent("John Doe", 1980, "701")
	s.Print()
	// вызовется метод String() для Student
	fmt.Println(s)
	fmt.Println(s.Name, s.Year, s.Group)
}
