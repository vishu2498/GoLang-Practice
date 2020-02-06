package main

/*Our program is just a series of machine instructions that need to be executed one after the other sequentially.
To make that happen, the operating system uses the concept of a Thread.
It’s the job of the Thread to account for and sequentially execute the set of instructions it’s assigned.
Execution continues until there are no more instructions for the Thread to execute. This is why Thread is called “a path of execution”.
Every program we run creates a Process and each Process is given an initial Thread. Threads have the ability to create more Threads.
All these different Threads run independently of each other and scheduling decisions are made at the Thread level, not at the Process level.
Threads can run concurrently (each taking a turn on an individual core), or in parallel (each running at the same time on different cores).
Threads also maintain their own state to allow for the safe, local, and independent execution of their instructions.
The OS scheduler is responsible for making sure cores are not idle if there are Threads that can be executing.
It must also create the illusion that all the Threads that can execute are executing at the same time.
In the process of creating this illusion, the scheduler needs to run Threads with a higher priority over lower priority Threads.
However, Threads with a lower priority can’t be starved of execution time. The scheduler also needs to minimize scheduling latencies as much as possible by making quick and smart decisions.*/

//Concurrency is the ability of the program to run multiple parts of code at the same time.

//A GoRoutine is a lightweight thread of execution.
//GoRoutines can be directly known as Multi-Threading in GoLang.
/*A GoRoutine is a function or method which executes independently and simultaneously in connection with any other GoRoutines present in a program.
Or in other words, every concurrently executing activity in GoLang is known as a GoRoutine.
We can consider a GoRoutine like a light weighted thread. The cost of creating GoRoutines is very small as compared to the thread.
Every program contains at least a single GoRoutine and that GoRoutine is known as the main GoRoutine.
All the GoRoutines are working under the main GoRoutine and if the main GoRoutine terminated, then all the GoRoutines present in the program also terminated.
GoRoutines always work in the background.
Its common for Go applications to have thousands of GoRoutines running concurrently.
 */
//We can think of GoRoutines as application-level threads and they are similar to OS Threads in many ways.
/*However, these GoRoutines do not use a complete OS thread to execute a single GoRoutine.
Multiple GoRoutine threads are deployed on a single OS thread and it is the responsibility of Go Runtime scheduler to schedule the timing of working of those threads on a single OS thread.*/

//To turn a function
func main() {

}
