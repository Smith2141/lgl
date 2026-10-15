package main

import (
	"fmt"
	"time"
)

// Собственный тип для ошибок. Аналогичен стандартному error
type Error interface {
	Error() string
}

// MyError — структура, реализующая нашу ошибку
type MyError struct {
	// ...
}

// Error — метод для удовлетворения интерфейсу Error
func (e *MyError) Error() string {
	return "..."
}

// переменная типа ошибки — указатель на пустую структуру
var ErrFriday13 = &MyError{}

func CheckTodayIsOkay() Error {
	var err *MyError // указатель на переменную типа Error
	// получаем текущее время
	t := time.Now()
	// если день недели пятница и число месяца 13, то вернём ошибку
	if t.Weekday() == time.Friday && t.Day() == 13 {
		fmt.Println("in condition")
		err = ErrFriday13
	}
	// вернём ошибку ... указатель же nil
	return err
}

func main() {
	err := CheckTodayIsOkay()
	// проверяем err на nil — и внезапно всегда не nil
	if err != nil {
		fmt.Println("error is not nil")
		return
	}

	fmt.Println("error is nil")
}
