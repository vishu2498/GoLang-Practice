package main

import (
	"errors"
	"testing"
)

func Test1(t *testing.T) {
	if f1()==nil {
		t.Fatal("received error")
	} else {
		t.Error("did not receive an error") //works similar to 'Log()' but uses function 'FailNow()' that reports that test-case is failed.
		//this is particularly used when test-case is failed but program should not stop execution.
		//'Errorf()' is also available.
	}
	t.Log("resumed")
}

func f1() error {
	var1 := errors.New("this is an error")
	return var1
}
