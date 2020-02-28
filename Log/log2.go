package main
import "log"

func main() {
	log.Panic("log panic") //will create a panic with the given string as a log
	//This function will exit the program with exit code 2.
	//'log.Panicf()' and 'log.Panicln()' are also present.
	//If the panic is not recovered, then the program will terminate with a stack trace.
}
