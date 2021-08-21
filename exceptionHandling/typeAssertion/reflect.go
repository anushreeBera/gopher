/**
Program: Get type from interface using reflect
Output:
received int 1
rceeived string hello
*/

package typeAssertion

import (
	"fmt"
	"reflect"
)

func Foo(data interface{}) {
	t := reflect.ValueOf(data).Kind()
	switch t {
	case reflect.String:
		fmt.Println("received string", data)
	case reflect.Int:
		fmt.Println("received int", data)
	}

	//to print type
	//fmt.sprintf("%T",data)

	//to print type via switch type
	// switch t := data.(type) {
	// 	case string:
	// 	case int:
	// 	case bool:
	// }
}

func main() {
	Foo(1)
	Foo("hello")
}
