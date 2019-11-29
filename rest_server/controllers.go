package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/Splash07/course3008/type_system/student"
)

func getAllStudents(c echo.Context) error {
	students := &student.Student{
		FirstName: "Dylan",
		LastName:  "Nguyen",
		Age:       "24",
		Email:     "ntd@netcompany.com",
	}
	return c.JSON(http.StatusOK, students)
}