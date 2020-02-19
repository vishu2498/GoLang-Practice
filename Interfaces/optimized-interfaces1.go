//Mine
//Everywhere the optimization is done, '//changed' is written.
/*In this optimized program of 'interfaces1.go', the '(int,error)' data-type in 'Write()' method is replaced with 'string'
data-type. It was present previously since the variables of the method were holding the data from
the fmt.Println() function which has the return type of (n int, err error). So, removal of fmt.Println() function gave the
variables the option to just hold the output of the from the struct and argument. Since only string was needed, only 'string'
data-type is used.*/
package main
import "fmt"

func main() {
	var w Writer = ConsoleWriter{}
	fmt.Println(w.Write([]byte("Hello vishu"))) //changed
}

type Writer interface {
	Write([]byte) string //changed
}

type ConsoleWriter struct {

}

func (cw ConsoleWriter) Write(data []byte) string { //changed
	n:=string(data) //changed
	return n //changed
}
