package main

import (
	"fmt"
	"sync"
)
var wg=sync.WaitGroup{}
func main() {
	//Using 'time.Sleep()' is not the best practice since it slows down the performance of the main() function and binds the application's runtime.
	//Also, if there would be more GoRoutines, 'time.Sleep()' will be used more and slow down the execution time.
	//The solution is using WaitGroups.

	//Wait Groups are designed to synchronize multiple goroutines together
	/*A WaitGroup waits for a collection of goroutines to finish.
	  The main goroutine calls 'Add' function of WaitGroup to set the number of goroutines to wait for.
	  Then each of the goroutines runs and calls 'Done' function of WaitGroup when finished.
	  At the same time, Wait can be used to block until all goroutines have finished.
	 */
	//It eliminates the need of time.Sleep()
	/*Add method is used to identify how many goroutines need to be waited.
	  When a goroutine exits, it must call Done().
	  The main goroutine blocks on Wait, once the counter becomes 0, the Wait() will return, and main goroutine can continue to run.
	 */
	var5:="testwg"
	wg.Add(1) //this is goroutine telling the WaitGroup to include itself for synchronization
	//Number in 'wg.Add()' tells the WaitGroup to add that much GoRoutines to wait for.
	go func(var5 string) {
		fmt.Println(var5)
		wg.Done() //this is the goroutine telling the WaitGroup that the work is done and it can leave the WaitGroup
	}(var5)
	var5="testing"
	wg.Wait() //this will make sure that main function doesn't go past this point until internal goroutine returns counter as 0
	//If the WaitGroup counter returns a negative value, 'wg.Add()' will panic.
}

//'Wait()' is actually returning a channel struct that emits a tick when the WaitGroup counter is zero.

//The WaitGroup counter can have a negative value in the case that if multiple GoRoutines return 'wg.Done()' more than the number defined in the 'wg.Add()'.
//In that case, 'wg.Done()' will decrease the counter to negative value and 'wg.Add()' will panic.
//So, it is best practice to only define 'wg.Add(1)' and immediately use the GoRoutine with 'wg.Done()'.

//'wg.Done()' can be written with 'defer' keyword too at the start of the GoRoutine.
