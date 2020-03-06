//To perform benchmarking, use command: go test -v -bench="BenchmarkTest"
//We have to mention the benchmark function name with "-bench" flag in "go test" command.

package main

import (
	"fmt"
	"testing"
)

/*Benchmarking is a way to test the performance of code. It’s useful when we want to test
the performance of different solutions to the same problem and see which solution performs better.
It can also be useful to identify CPU or memory issues for a particular piece of code that might be critical to the performance of our application.
Benchmarking can be used to test different concurrency patterns*/

//The test file name must end with "_test.go".
//Benchmark functions begin with the word "Benchmark" and take as their only parameter a pointer of type 'testing.B'.
//In order for the benchmarking framework to calculate performance, it must run the code over and over again for a period of time.

//B is a type passed to Benchmark functions to manage benchmark timing and to specify the number of iterations to run.
//B is a struct.
func BenchmarkTest(b *testing.B) {
	//ResetTimer zeroes the elapsed benchmark time and memory allocation counters and deletes user-reported metrics.
	//It does not affect whether the timer is running.
	b.ResetTimer()

	for i:=0; i<10; i++ {
		fmt.Println(10)
	}
}