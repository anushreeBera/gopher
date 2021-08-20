/**
Program: Read, write channel with waitgroup
Output:
data: 0 active: true
data: 1 active: true
data: 2 active: true
data: 3 active: true
data: 4 active: true
data: 5 active: true
data: 6 active: true
data: 7 active: true
data: 8 active: true
data: 9 active: true
data: 0 active: false
data: 0 active: false
data: 0 active: false
data: 0 active: false
data: 0 active: false
*/

package channel

import (
	"fmt"
	"sync"
)

func Read(ch <-chan int) {
	//data, ok := <-ch
	/*for data := range ch {
		fmt.Println("reader:",data)
	}*/
	for i := 0; i < 15; i++ {
		data, ok := <-ch
		fmt.Println("data:", data, "active:", ok)
	}
}

func Write(ch chan<- int) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func main() {
	ch := make(chan int, 10) // sized channel
	//ch2 := make(chan int) // un buffered channel
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		Read(ch)
	}()
	go func() {
		defer wg.Done()
		Write(ch)
	}()

	wg.Wait()
	//Read(ch2)
}
