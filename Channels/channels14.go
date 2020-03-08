//Merging two channels so that two channels send values to the same channel and we receive it from final channel.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chan1 := asChan(1, 3, 5, 7)
	chan2 := asChan(2, 4, 6, 8)
	chan3 := merge(chan1, chan2)
	for v := range chan3 {
		fmt.Println(v)
	}
}

func asChan(values ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range values {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func merge(chan1, chan2 <-chan int) <-chan int {
	newchan := make(chan int)
	go func() {
		defer close(newchan)
		chan1done, chan2done := false, false
		for !chan1done || !chan2done {
			select {
			case receive, ok := <-chan1:
				if !ok {
					chan1done = true
					continue
				}
				newchan <- receive
			case receive, ok := <-chan2:
				if !ok {
					chan2done = true
					continue
				}
				newchan <- receive
			}
		}
	}()
	return newchan
}
