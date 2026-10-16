package main

import (
	"fmt"
	"reflect"
)

type MyType struct{}

func IsNil(obj any) bool {
    if obj == nil {
        return true
    }

    objValue := reflect.ValueOf(obj)
    // проверяем, что тип значения ссылочный, то есть в принципе может быть равен nil
    if objValue.Kind() != reflect.Pointer {
        return false
    } 
    // проверяем, что значение равно nil 
    //  важно, что IsNil() вызывает панику, если value не является ссылочным типом. Поэтому всегда проверяйте на Kind() 
    if  objValue.IsNil() {
        return true
    }

    return false
}

func main() {
    var t *MyType
    fmt.Printf("Проверка типа (%v) на nil: %v\n", reflect.TypeOf(t), IsNil(t)) // TypeOf возвращает тип переданного объекта. 
}