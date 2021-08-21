/**
Program: Inheritance
Output:
hello
*/

package OOP

import (
	"fmt"
)

type Child interface {
	Method2()
	Parent
}

type Parent interface {
	Method1()
}

type child struct {
	I int
	S string
	parent
}

type parent struct {
	S string
}

func main() {
	ch := child{I: 10, parent: parent{S: "hello"}}
	fmt.Println(ch.S, ch.parent.S)
}
