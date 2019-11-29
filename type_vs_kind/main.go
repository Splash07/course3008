package main

import (
	"fmt"
	"github.com/Splash07/course3008/type_system/student"
	"reflect"
)

func main() {
	aStruct := struct {
		Name string
		Age  int
	}{Name: "Dylan", Age: 35}

	aType := reflect.TypeOf(aStruct)

	fmt.Printf("aStruct has type %s and kind %s\n\n", aType.Name(), aType.Kind())

	bStruct := student.Student{FirstName: "Dylan", LastName: "Nguyen"}
	bType := reflect.TypeOf(bStruct)
	fmt.Printf("bStruct has type %s and kind %s\n\n", bType.Name(), bType.Kind())

	fmt.Println(bType.Field(0))
}
