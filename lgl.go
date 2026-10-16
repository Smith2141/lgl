package main

import (
	"fmt"
	"reflect"
)

func ExtendedPrint(v any) {
	val := reflect.ValueOf(v)
	//  проверяем, а не передали ли нам указатель на структуру
	switch val.Kind() {
	case reflect.Pointer:
		fmt.Println("Hello 1!!!")
		if val.Elem().Kind() != reflect.Struct {
			fmt.Println("Hello 2!!!")
			fmt.Printf("Pointer to %v : %v", val.Elem().Type(), val.Elem())
			return
		}
		// если всё-таки это указатель на структуру, дальше будем работать с самой структурой
		val = val.Elem()

	case reflect.Struct: // работаем со структурой
	default:
		fmt.Println("Hello 3!!!")
		fmt.Printf("%v : %v", val.Type(), val)
		return
	}

	fmt.Println("Hello 4!!!")
	fmt.Printf("Struct of type %v and number of fields %d:\n", val.Type(), val.NumField())
	for fieldIndex := 0; fieldIndex < val.NumField(); fieldIndex++ {
	// for _,fieldIndex range val.NumField() {
		field := val.Field(fieldIndex) // field — тоже Value
		fmt.Printf("\tField %v: %v - val :%v\n", val.Type().Field(fieldIndex).Name, field.Type(), field)
		// имя поля мы получаем не из значения поля, а из его типа.
	}
}

type MyStruct struct {
	A int
	B string
	C bool
}

func main() {
	// s := MyStruct{
	// 	A: 3,
	// 	B: "some",
	// 	C: false,
	// }
	s1 := &MyStruct{
		A: 7,
		B: "text",
		C: true,
	}

	// ExtendedPrint(s)
	ExtendedPrint(s1)
	// ExtendedPrint(struct {
	// 	E int
	// 	C string
	// }{2, "other text"})
	// ExtendedPrint("some string")
}
