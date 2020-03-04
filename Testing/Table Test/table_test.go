package main

import "testing"

/*A table test is like a basic unit test except it maintains a table of different values and results.
The different values are iterated over and run through the test code. With each iteration, the results are checked.
This helps to leverage a single test function to test a set of different values and conditions.*/
//The concept of "test tables" is a set (slice array) of test input and output values.
//Using for-range loop is common in table tests for iterations.
func TestCase1(t *testing.T) {
	values:=[]struct{
		val1 int
		val2 int
		val3 int
	}{
		{10,20,30,},
		{20,30,50},
		{50,20,70},
	}
	for _, table := range values {
		total := f1(table.val1, table.val2)
		if total != table.val3 {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.val1, table.val2, total, table.val3)
		}
	}
}

func TestCase2(t *testing.T) {
	values:=[]struct{
		val1 int
		val2 int
		val3 int
	}{
		{10,20,30,},
		{20,30,50},
		{50,20,70},
	}
	for _, table := range values {
		total := f2(table.val1, table.val2)
		if total != table.val3 {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.val1, table.val2, total, table.val3)
		}
	}
}

func f1(value1, value2 int) int {
	result := value1 + value2
	return result
}

//Intentionally changed the addition operator to multiplication operator to make the test-case fail.
func f2(value1, value2 int) int {
	result := value1 * value2
	return result
}