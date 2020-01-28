//A GO application finally shuts down and closes when main() ends

//main() in Golang can also cause goroutines that are present inside it to forcefully shut down.
//It may happen when the goroutines are functioning but the runtime for main() has ended
//We can use 'defer' to avoid this and close the functioning of main() only after completion of channel & goroutine

//Signal only channels in Golang are used when we require zero memory allocation for the channel and the receiver should know about the value sent
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
func main() {
	go logger()
	defer func() {
		close(logCh)
	}()
	logCh<-logEntry{time.Now(),logInfo,"App is starting"}
	logCh<-logEntry{time.Now(),logInfo,"App is shutting down"}
	time.Sleep(100*time.Millisecond)
}
func logger() {
	for entry:=range logCh {
		fmt.Printf("%v - [%v]%v\n",entry.time.Format("2006-01-02T15:05:05"),entry.severity,entry.message)
	}
}

