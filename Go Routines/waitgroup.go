package main

import (
	"fmt"
	"sync"
)
var wg=sync.WaitGroup{}
func main() {
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
	wg.Add(1) //this is goroutine telling the waitgroup to include itself for synchronization
	go func(var5 string) {
		fmt.Println(var5)
		wg.Done() //this is the goroutine telling the waitgroup that the work is done and it can leave the waitgroup
	}(var5)
	var5="testing"
	wg.Wait() //this will make sure that main function doesn't go past this point until internal goroutine returns counter as 0
}
