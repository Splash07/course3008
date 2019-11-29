package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/Splash07/course3008/type_system/student"
)

func getAllStudents(c echo.Context) error {
	students := [] student.Student {
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

	return c.JSON(http.StatusOK, students)
}