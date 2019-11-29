package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/Splash07/course3008/rest_server/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	students := []student.Student{
		{
			FirstName: "Dylan",
			LastName:  "Nguyen",
			Age:       24,
			Email:     "ntd@netcompany.com",
		},
		{
			FirstName: "Ha",
			LastName:  "Hoang",
			Age:       18,
			Email:     "hhh@netcompany.com",
		},
		{
			FirstName: "Dat",
			LastName:  "Nguyen",
			Age:       24,
			Email:     "ntda@netcompany.com",
		},
	}
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig {
		AllowOrigins: [] string{"*"},
		AllowMethods: [] string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))


	// Route => handler

	// No Param
	e.GET("/all_students", controllers.GetAllStudents(students))

	// Params: first_name, last_name
	e.GET("/student", controllers.GetStudentByFullName(students))

	// No Param
	e.POST("/new_student", controllers.AddStudent(students))

	// Start Server
	e.Logger.Fatal(e.Start(":1010"))
}