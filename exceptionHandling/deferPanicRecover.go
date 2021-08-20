/**
Program: Defer, panic and recover example
Output:
panic recovered: I just panic
*/

package exceptionhandling

import (
	"fmt"
)

func panicer() {
	panic("I just panic")
}

func foo() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("panic recovered:", err)
		}
	}()
	panicer()
}

func main() {
	foo()
}
