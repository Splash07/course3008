package main

import 
	"fmt"

type Student struct {
	FirstName string `json: "first_name bson: "full_name" validate: "required"`
	LastName  string
	Age       int
	Email     string
}

func (s Student) GetFullName() string {
	return s.FirstName + " " + s.LastName
}

func (s *Student) SetEmail(e string) {
	s.Email = e
}

func main() {
	// fmt.Println("hello dylan")

	// // Int
	// aInt := 1
	// // fmt.Println(aInt)

	// var bInt int
	// bInt = 2
	// // fmt.Printf("bInt = %d\n", bInt)

	// // String
	// aString := "Hello abc"
	// // fmt.Printf("aString = %s\n", aString)

	// // Array/Slice
	// aSlice := []string{"a", "b", "c"}
	// // fmt.Println(aSlice)

	// bSlice := []int{1, 4, 6, 8, 8}
	// // fmt.Println(bSlice)

	// // for _, v := range bSlice {
	// // 	if v%2 == 0 {
	// // 		fmt.Println(v)
	// // 	}
	// // }

	// // Map
	// aMap := map[string]int{"age": 1000, "level": 5}
	// // fmt.Println(aMap)

	dylan := Student{
		FirstName: "Dylan",
		LastName:  "Nguyen",
		Age:       100,
		Email:     "nguyendat293@gmail.com",
	}

	// // dat := Student{"Dat", "Nguyen", 100, "nguyendat293@gmail.com"}

	// // fmt.Printf("%+v", dylan)
	// // fmt.Printf("%+v", dat)

	// // Interface
	// // var i interface{}
	// // i = dylan
	// // fmt.Println(i)

	// // Channel
	// // ch := make(chan int, 2)
	// // ch <- 10 // write to channel
	// // ch <- 10 // write to channel

	// // // Fmt.Println(<-ch)
	// // Fmt.Println(<-ch)

	// // Function

	// // Function receiver
	fmt.Println(dylan.GetFullName())
	dylan.SetEmail("lolololol")

	fmt.Println(dylan.Email)

	// Int <----> String
	// 1st method: WIthout using a library
	anInt := 100
	cString := fmt.Sprintf("%d hello", anInt)
	fmt.Printf(cString)

	// 2nd method


	// String ----> Int
	aString := "12345"
    newInt := fmt.
}
