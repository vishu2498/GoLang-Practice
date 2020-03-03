package main

import "testing"

func Test(t *testing.T) {
	value:=f1(10,20)
	if value == 30 {
		t.Log("value is correct")
	} else {
		t.Error("value is incorrect")
	}
}

func f1(value1, value2 int) int {
	result := value1 + value2
	return result
}