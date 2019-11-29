package main

import (
	controllers "github.com/Splash07/course3008/rest_server_api/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Route => handler

	// No Param
	e.GET("/api/all_students", controllers.GetAllStudents)

	// Params: first_name, last_name
	e.GET("/api/student", controllers.GetStudentByFullName)

	// No Param
	e.POST("/api/new_student", controllers.AddStudent)

	// Start Server
	e.Logger.Fatal(e.Start(":1010"))
}
