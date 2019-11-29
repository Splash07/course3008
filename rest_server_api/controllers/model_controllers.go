package controllers
 
import (
    "net/http"
 
    "github.com/Splash07/course3008/type_system/student"
    "github.com/labstack/echo"
)
 
var students = []student.Student{
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
 
//------------
// Handlers
//------------
 
func GetAllStudents(c echo.Context) error {
    return c.JSON(http.StatusOK, students)
}
 
func AddStudent(c echo.Context) error {
    var newStudent student.Student
 
    if err := c.Bind(&newStudent); err != nil {
        return err
    }
 
    students = append(students, newStudent)
    return c.JSON(http.StatusCreated, newStudent)
}
 
func GetStudentByFullName(c echo.Context) error {
    firstName := c.Param("first_name")
    lastName := c.Param("last_name")
    for i := range students {
        if students[i].FirstName == firstName && students[i].LastName == lastName {
            return c.JSON(http.StatusOK, students[i])
        }
    }
 
    var err error
    return err
}