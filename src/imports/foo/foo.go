package foo

import "fmt"

var _ = func() string {
	fmt.Println("foo global variable")
	return ""
}()

func init() {
	fmt.Println("foo init")
}

func Bar() {
	fmt.Println("This function lives in an another file!")
}