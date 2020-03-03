package main

import "testing"

func Test(t *testing.T) {
	if f1() == 30 {
		t.Log("value is correct")
	} else {
		t.Error("value is incorrect")
	}
}

func f1() int {
	var1, var2 := 10, 20
	var3 := var1 + var2
	return var3
}