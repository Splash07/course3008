package main

import (
	"fmt"

	"github.com/dattiennguyen/course3008/type_system/student"
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

}
