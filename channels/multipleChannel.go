/**
Program: Multiple channels
Output:
received int: 0
received int: 1
received string: string:0
received string: string:1
received int: 2
received string: string:2
received string: string:3
received int: 3
received string: string:4
received string: string:5
received string: string:6
received string: string:7
received string: string:8
received string: string:9
received string:
received string:
received string:
received string:
received int: 4
received string:
received string:
received string:
received int: 5
received string:
received string:
received string:
received string:
received string:
received string:
received string:
received int: 6
received string:
received string:
received string:
received string:
received string:
received string:
received int: 7
received string:
received string:
received int: 8
received string:
received int: 9
received string:
received empty struct: {}
*/

package channel

import (
	"fmt"
	"sync"
)

func ReadFromMultipleChan(ch1 chan int, ch2 chan string, ch3 chan struct{}) {
	for {
		select {
		case i := <-ch1:
			fmt.Println("received int:", i)
		case s := <-ch2:
			fmt.Println("received string:", s)
		case empty := <-ch3:
			fmt.Println("received empty struct:", empty)
			return
		}
	}
}

func writeInt(ch1 chan int) {
	defer close(ch1)
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
}

func writeString(ch1 chan string) {
	defer close(ch1)
	for i := 0; i < 10; i++ {
		ch1 <- fmt.Sprintf("string:%d", i)
	}
}

func writeEmpty(ch1 chan struct{}) {
	defer close(ch1)
	ch1 <- struct{}{}
}

func mainFunc() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		writeInt(ch1)
	}()
	go func() {
		defer wg.Done()
		writeString(ch2)
	}()
	wg2 := sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		ReadFromMultipleChan(ch1, ch2, ch3)
	}()
	wg.Wait()       // waiting for int go routine and string go routine to complete
	writeEmpty(ch3) // triggering close of read go routine
	wg2.Wait()      // will wait for read channel to end.
}
