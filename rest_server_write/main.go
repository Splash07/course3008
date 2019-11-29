package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Splash07/course3008/type_system/student"

	// "log"
	// "net/http"
	// "github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Name of the database.
	DBName = "GolangTest"
	//URI    = "mongodb://dylannguyen:root123@clustergolangclass-ykxin.mongodb.net:27017"
	URI = "mongodb+srv://dylannguyen:root123@clustergolangclass-ykxin.mongodb.net/test?retryWrites=true&w=majority"
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
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderContentType},
	}))

	// Route => handler
	e.GET("/students_db", WriteStudentsToDatabase)

	// Start Server
	e.Logger.Fatal(e.Start(":9090"))
}

//-------------
// Handlers----
//-------------
func WriteStudentsToDatabase(c echo.Context) error {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	// consume GetAllStudents API from rest_server_api service
	response, err := client.Get("http://localhost:1010/api/all_students")

	if err != nil {
		fmt.Println("Error Occured!")
		return err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	// Unmarshal the response into a slice of Students struct
	var studentPayload []student.Student
	if err = json.Unmarshal(body, &studentPayload); err != nil {
		return err
	}

	//-------------open connection to MongoDB-------------

	// Set client options
	clientOptions := options.Client().ApplyURI(URI)

	// Connect to MongoDB
	dbclient, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println(err)
	}
	//-----------------------------------------------------

	// insert list of Students to Students collection of GolangTest DB
	collection := dbclient.Database(DBName).Collection("Students")

	// convert studentPayload to a slice of interfaces
	listOfStudents := make([]interface{}, len(studentPayload))
	for i, v := range studentPayload {
		listOfStudents[i] = v
	}

	insertManyResult, err := collection.InsertMany(context.TODO(), listOfStudents)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted class list: ", insertManyResult.InsertedIDs)

	var catchError error
	return catchError
}
