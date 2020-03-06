//To perform benchmarking, use command: "go test -v -bench=BenchmarkCounterTest"
//We can use '-run="none"' flag with "go test" command to tell GoLang compiler to NOT run unit tests.
package main

import (
	"fmt"
	"testing"
)

func BenchmarkCounterTest(b *testing.B) {
	//'N' is a counter (int) in the struct 'B'.
	//Each time the framework calls the benchmark function, it will increase the value of b.N.
	//On the first call, the value of b.N will be 1.

	number := 98
	b.ResetTimer()

	//It’s important to place all the code to benchmark inside the loop and to use the b.N value.
	//If this isn’t done, the results can’t be trusted.
	for i := 0; i < b.N; i++ {
		//converting 'int' to 'string'
		fmt.Sprintf("%d", number) //'fmt.Sprintf()' returns string
	}
}

//In the output of Benchmark Test, we get a number near "PASS" keyword and the test name.
//This number represents the number of times the code inside the loop was executed.

//After this, we get a number with "ns/op" keyword.
//This number represents the performance of the code based on the number of nanoseconds per operation.

//After this, "ok" keyword represents that the benchmark finished properly.
//Then the name of the code file that was executed is displayed.

//After this, it displays the total time the benchmark ran. The default minimum run time for a benchmark is 1 second.
//We can use another flag called "-benchtime" with "go test" if we want to have the test run longer. (eg. -benchtime="3s")
//But this will increase the number of times the code inside the loop gets executed.
//Sometimes by increasing the bench time, we can get a more accurate reading of performance.