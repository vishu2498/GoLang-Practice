package main

import (
	"fmt"
	"runtime"
	"sync"
)
//When we use multiple goroutines with waitgroups, they create a mess while being executed
//They perform in jumbled order since each one of them will be racing against each other to execute
//This is called as racing condition
//So, while waitgroups solve the problem of running multiple goroutines at once, it can't execute each one of them in the order as they are called

//Mutex is basically a lock that the application is going to honor.
//If mutex is locked and something is trying to manipulate the values that are locked by mutex, it has to wait till the mutex unlocks
//Mutex is used to protect parts of the code so that only one entity can be manipulating the code at a time
//So this is used also ensure that only one operation is occuring at a time on some data

//RWMutex = Read Write Mutex
//RWMutex allows multiple (or infinite) entities to read the data at once but allows only one entity to write or manipulate the data
//And if any data is being read at a time, it can't we written or manipulated at the same time
/*So, if anything wants to write the data, it has to wait until all that data is bring read.
  And then the writer comes in and locks the mutex so that it can perform the writing and no other thing can read or write the data at that time.
*/

//RLock() is Read Lock
//RUnlock() is Release the Read Lock
//Lock() is Write Lock
//Unlock() is Release the Write Lock

//Lock the Mutex outside the context of goroutine so that we have a reliable execution model
/*'GOMAXPROCS' keyword is used to control the number of threads available for Goroutine execution to a particular go program.
  The hard limit is still on the number of CPU cores presented to the OS.
  The GOMAXPROCS option allows us to tune it down.
*/

//We are first locking the mutexes first and unlocking them inside the body of the goroutine when they are done executing
//If the goroutine printhello1() runs and goroutine increment1() doesn't run at the same time, that's where the printing from printhello1() will occur twice
var wg1 =sync.WaitGroup{}
var counter1 = 0
var m=sync.RWMutex{}
func main() {
	runtime.GOMAXPROCS(100) //we are defining this value to tell the application to use this many threads
	//GOMAXPROCS(1) is used when we want to execute single-threaded application
	//Setting GOMAXPROCS() to any negative value will have no effect
	for i:=0;i<10 ;i++  {
		wg1.Add(2)
		m.RLock() //here, we are trying for Read Lock so that we can protect 'counter' variable from concurrent reading & writing
		go printHello1()
		m.Lock()
		go increment1()
	}
	wg1.Wait()
}
func printHello1()  {
	fmt.Println("Hello ",counter1)
	m.RUnlock()
	wg1.Done()
}
func increment1() {
	counter1++
	m.Unlock()
	wg1.Done()
}

/*If the Read & Write Locks were present in both the GoRoutines, then even if the counter will be fixed
but the result will still be not appropriate because we don't have chance to lock the mutex before we are trying to read it second time.
Hence, the Read & Write Locks were placed in the main() function. So, the Locking is done in the main() function and
the Locks are individually unlocked at the individual function level.*/