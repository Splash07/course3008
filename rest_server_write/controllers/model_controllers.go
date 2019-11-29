package controllers

import (
	"github.com/Splash07/course3008/type_system/student"
	"github.com/labstack/echo"
	"net/http"
)

//------------
// Handlers
//------------

func GetAllStudents(students[] student.Student, c echo.Context) error {
	return c.JSON(http.StatusOK, students)
}

func AddStudent(students[] student.Student, c echo.Context) error {
	newStudent := new (student.Student)

	if err := c.Bind(newStudent); err != nil {
		return err
	}

	students = append(students, newStudent)
	return c.JSON(http.StatusCreated, newStudent)
}

func GetStudentByFullName(students[] student.Student, c echo.Context) error {
	firstName, _ := strconv.Atoi(c.Param("first_name"))
	lastName, _ := strconv.Atoi(c.Param("last_name"))
	for i := range students {
		if (students[i].FirstName == firstName && students[i].LastName == lastName) {
			return c.JSON(http.StatusOK, students[i])
		}
	}
	
}


