package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/Splash07/course3008/rest_server/controllers"
	"github.com/labstack/echo"
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
	e.GET("/all_students", controllers.getAllStudents)

	// Server
	e.Run(standard.New(":8080"))
}