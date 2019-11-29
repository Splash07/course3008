package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/Splash07/course3008/rest_server/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
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
	e.GET("/all_students", controllers.GetAllStudents)

	// Server
	e.Run(standard.New(":8080"))
}