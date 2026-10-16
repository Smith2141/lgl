package main

import (
	"fmt"
	"reflect"
)

type MyType struct{}

// func IsNil(obj any) bool {
// 	if obj == nil {
// 		return true
// 	}

// 	objValue := reflect.ValueOf(obj)
// 	// проверяем, что тип значения ссылочный, то есть в принципе может быть равен nil
// 	if objValue.Kind() != reflect.Pointer {
// 		return false
// 	}
// 	// проверяем, что значение равно nil
// 	//  важно, что IsNil() вызывает панику, если value не является ссылочным типом. Поэтому всегда проверяйте на Kind()
// 	if objValue.IsNil() {
// 		return true
// 	}

// 	return false
// }

func main() {
	varInt := 100
	// var varInt int
	varIntValue := reflect.ValueOf(varInt)
	fmt.Println(varIntValue.IsZero()) // false
	fmt.Println(varIntValue.Int())    // 100

	var varPtr *int
	varPtrValue := reflect.ValueOf(varPtr)
	fmt.Println(varPtrValue.IsNil())  // true
	fmt.Println(varPtrValue.IsZero()) // true
}
