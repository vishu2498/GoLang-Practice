//The test-file name must end with '_test.go'.
//To run the test, use command: "go test -v"
//"-v" is for verbose output
package _st_Test

import (
	"errors"
	"testing" //this package provides all the methods for testing purposes
)

//Testing function must be an exported function and must start with "Test" keyword.
//It must not return something.
//It must take one argument of data-type (pointer) '*testing.T' where 'T' is a struct.
func Test1(t *testing.T) {
	t.Log("passing the 1st test") //works similar to 'Println()' and records the text in the error log
	//It will be only displayed when '-v' is used with 'go test' or if the test fails.
	t.Logf("%v, %T","1st test passed","1st test passed") //works similar to 'Printf()' with string formatting
}

func Test2(t *testing.T) {
	//This test-case will fail since we received an error and it is not nil.
	if f1()!=nil {
		t.Fatal("received error") //'Fatal()' is similar to 'Log()' but uses function 'FailNow()' that reports that test-case is failed.
		//test-case will compulsorily fail when 't.Fatal()' is used and program will stop execution.
	}
}

func f1() error { //function returning an error
	var1:=errors.New("this is an error")
	return var1
}

//'FailNow()' calls 'Fail()' and 'Fail()' marks the test-case as failed.