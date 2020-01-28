//'select' is mainly used when there are many channels inside the program and you require a goroutine to monitor those channels and terminate them

package main

import (
	"fmt"
	"time"
)
const (
	logInfo="INFO"
	logWarning="WARNING"
	logError="ERROR"
)

type logEntry struct {
	time time.Time
	severity string
	message string
}
var logCh = make(chan logEntry,50)
//Signal only channels in Golang are used when we require zero memory allocation for the channel and the receiver should know about the value sent
var doneCh = make(chan struct{})
func main() {
	go logger()
	logCh<-logEntry{time.Now(),logInfo,"App is starting"}
	logCh<-logEntry{time.Now(),logInfo,"App is shutting down"}
	time.Sleep(100*time.Millisecond)
	doneCh<- struct{}{} //double {} means that 1st bracket is showing that struct has no fields and 2nd bracket is intializaing the struct
}
func logger() {
	for {
		select {
		case entry:=<-logCh:
			fmt.Printf("%v - [%v]%v\n",entry.time.Format("2006-01-02T15:05:05"),entry.severity,entry.message)
		case <-doneCh:
			break
		}
	}
}
