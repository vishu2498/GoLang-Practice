//To perform benchmarking, use command: "go test -v -bench=." ('.' tells GoLang compiler to test all benchmark functions)
//"-benchmem" flag can be used with "go test" to provide information about the number of allocations and bytes per allocation for a given test.

package Performance_Benchmark_Test

import (
	"fmt"
	"strconv"
	"testing"
)

//Checking which function mentioned inside benchmark functions perform faster.

// BenchmarkSprintf() provides performance numbers for the 'fmt.Sprintf()' function.
func BenchmarkSprintf(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}

// BenchmarkFormat() provides performance numbers for the 'strconv.FormatInt()' function.
func BenchmarkFormat(b *testing.B) {
	number := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

// BenchmarkItoa() provides performance numbers for the 'strconv.Itoa()' function.
func BenchmarkItoa(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}

//While using "-benchmem" flag with "go test" we get 2 new values with each test name.

//"allocs/op" represents the number of heap allocations per operation.
//"B/op" represents the number of bytes per operation.