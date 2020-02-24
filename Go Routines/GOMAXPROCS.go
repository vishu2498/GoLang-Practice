package main

import (
	"fmt"
	"runtime"
)

func main() {
	//By default, GoLang will provide the number of OS threads equal to number of cores on the PC which is equal to 'runtime.NumCPU()'.

	/*'GOMAXPROCS' keyword is used to control the number of threads available for Goroutine execution to a particular go program.
	  The hard limit is still on the number of CPU cores presented to the OS.
	  The GOMAXPROCS option allows us to tune it down.
	*/
	//GOMAXPROCS(1) is used when we want to execute single-threaded application (no parallelism)
	//Setting GOMAXPROCS() value to more than 1 will enable parallelism.
	//Setting GOMAXPROCS() to any negative value will have no effect

	fmt.Println("Threads available:",runtime.GOMAXPROCS(3))
	//Even if we define the number of GOMAXPROCS greater than the available cores, we will still have those max cores.
	//This means that the number we defined, the scheduler will spawn those number of GoLang threads on a single OS threads.
	//So, when the number provided is too large, it can cause trouble for scheduler to launch small threads everytime in a short period of time.
}