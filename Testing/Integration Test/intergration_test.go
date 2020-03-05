// +build integration
//This tag is used to tell GoLang compiler that we are performing integration testing here.

/*Integration tests are those tests which verify that different modules or services used by our application work together.
Also, these tests can be combination of multiple unit tests. These tests also check some code for how it works without actually changing the code.*/

//Integration tests are slower than unit tests.
//So, we provide "-short" flag with "go test" to reduce time consumption. It uses 'testing.Short()' function. It returns true if "-short" flag is used.
//Tests can check the value of 'testing.Short()' to determine whether they should execute code and/or tests that will take a long time.

//To perform this integration test, use command : "go test -tags=integration"
//This will specifically perform integration testing and not unit testing.
package main

import (
	"testing"
)

func f1(x, y int) int {
	return x + y
}

func TestIntegration(t *testing.T) {
	var1:=f1(2,2)
	if var1!=4 {
		t.Fatal("value not correct")
	}
}
