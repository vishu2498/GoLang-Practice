// Running all HTTP queries for this program requires Postman to be installed.

package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// making a sample struct on which operations would be performed
type values struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// this a sample data using the "values" struct
var finalValues = []values{
	{Name: "Vishu", Age: 24},
	{Name: "Saras", Age: 21},
	{Name: "Sakshi", Age: 23},
}

func main() {
	// gin is just like another HTTP handler package
	// Calling the "Default()" function gives us an engine to work with that will serve as an HTTP Handler. It contains both a Logger and a Recovery function that recovers from panics.
	router := gin.Default()

	// All the below functions take an endpoint from URL (localhost:9090) and a handler function

	// This works as HTTP GET method and is used for getting all values
	router.GET("/values", getValues)

	// This is used to get a specific single value from array of values
	router.GET("/values/:name", getSingleValue)

	// This works as HTTP POST request and is used to add a value (given from user) to array of values
	router.POST("/values", addValues)

	// This works as HTTP PATCH request and is used to change a single value
	// TODO: This is currently not working. Need to fix it afterwards.
	router.PATCH("/values/:name", changeSingleValue)

	// "Run" will start the server on the given domain and port
	// This must be place at the bottom after defining all the HTTP methods and any processing related to them
	router.Run("localhost:9090")
}

// The "context" contains all the requests and responses from and for the user and can be used for all HTTP processing purposes.
func getValues(context *gin.Context) {
	// "IndentedJSON" provides the response by converting all the golang object to JSON format.
	// It takes the HTTP method as 1st argument and the data to be converted in 2nd argument.
	context.IndentedJSON(http.StatusOK, finalValues)
}

func addValues(context *gin.Context) {
	var newValue values
	// "BindJSON" binds the received data from user in JSON format
	if err := context.BindJSON(&newValue); err != nil {
		return
	}

	// appending the newly received data to final data
	finalValues = append(finalValues, newValue)

	context.IndentedJSON(http.StatusCreated, newValue)
}

// checking that the name given by user exists in the array of values or not
// this is not a handler function.. just a processing function
func getValueByName(name string) (*values, error) {
	for _, value := range finalValues {
		if value.Name == name {
			return &value, nil
		}
	}

	return nil, errors.New("the requested name could not be found")
}

func getSingleValue(context *gin.Context) {
	// Taking parameter given by user. The argument given here must be equal what the HTTP method call contains in its endpoint URL.
	// It returns the value in string format.
	nameGotFromParameter := context.Param("name")

	value, err := getValueByName(nameGotFromParameter)

	if err != nil {
		// gin.H{} is map[string]interface{} usually used for custom messages
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "the requested name could not be found"})
		return
	}

	context.IndentedJSON(http.StatusOK, value)
}

func changeSingleValue(context *gin.Context) {
	nameGotFromParameter := context.Param("name")

	value, err := getValueByName(nameGotFromParameter)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "the requested name could not be found"})
		return
	}

	value.Name = "random-name"

	context.IndentedJSON(http.StatusOK, value)
}

/*
How to run this program: (and its flow)
1. go get github.com/gin-gonic/gin
2. Run this program normally using command like: "go run file.go"
3. For 1st GET method, run the URL: "localhost:9090/values". This will print all the values present in "finalValues" variable.
4. For 2nd GET method, run the URL: "localhost:9090/values/Vishu". This will print the array element having field name "Vishu"
5. For 3rd POST method, go to "Body" section. Select "raw" option and write like this:
   {
    "name": "Sarthak",
    "age": 22
   }
   Now, run the URL: "localhost:9090/values" and check that 201 status should be present.
   Now, run the URL: "localhost:9090/values" to check final values.
*/
