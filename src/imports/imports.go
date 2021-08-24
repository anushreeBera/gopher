/**
Program: Imports sequence
Output:
foo global variable
foo init
global variable
main init
This function lives in an another file!
*/

package imports

import (
	"fmt"

	"play.ground/foo"
)

var _ = func() string {
	fmt.Println("global variable")
	return ""
}()

func init() {
	fmt.Println("main init")
}

func main() {
	foo.Bar()
}
