package main

import (
	"fmt"
    "strconv"
	"github.com/Splash07/course3008/type_system/student"
	"reflect"
)

func main() {
	//map

	// aMap := map[string]int{"a": 1, "b": 3, "c": 5}

	// for i, v := range aMap {
	// 	fmt.Printf("index at %s has value %d \n", i, v)
	// }

	// //map[string]interface{}

	// aStudent := student.Student{FirstName: "Dylan", LastName: "Nguyen", Age: "10", Email: "lol@gmail.com"}

	// bs, _ = json.Unmarshal(aStudent)

	// var m map[string]interface{}

	// json.Unmarshal()

	// channel
	go func ({
		ch <- 1
		ch <- 2
		ch <- 8
		ch <- 4
		close(ch)
	})()

	for v := range ch {
		fmt.Println(v)
	}

	// Type casting (string <-> byte slice)
	aString := "hello world"package main
	
	bs := []byte(aString)

	bString := string(bs)

	fmt.Println(bString)

	aInt := 100

	aString = strconv

	fmt.Println(aString + "mot tram")

	bInt, err := strconv.Atoi(aString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bInt + 100)

	// int <-> float

	aInt = 300

	var aFloat float32
	aFloat = float32(300)
	type1 := reflect.TypeOf(aFloat)
	fmt.Println(aFloat, "type=", type1.Name(), "kind=", type1.Kind())

	cInt := int(aFloat)
	type1 = reflect.TypeOf(cInt)
	fmt.Println(cInt, "type=", typ1.Name(), "kind=", type1.Kind())

	//struct <-> struct

	type Boy struct {
		FirstName string
		LastName  string
		Age       int
		Email     string
	}

	boy := Boy {"Dylan", "Nguyen", 300, "dylan@gmail.com"}
