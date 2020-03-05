package main

import (
	"io/ioutil"
	"testing"
)

func TestReadFile(t *testing.T) {
	data,err:=ioutil.ReadFile("file.txt")

	if err!=nil {
		t.Fatal("could not open file")
	}

	if string(data) != "hello from vishu" {
		t.Fatal("data did not match")
	}
}
