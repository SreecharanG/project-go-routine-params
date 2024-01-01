package main

import (
	"fmt"
	"time"
	"strings"
)


func currencyStructure(num string){
	fmt.Println(num)
}


func channels() {
	myChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	msg := <-myChannel
	fmt.Println(msg)
}


func channelsSelectPrimitive() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	go func() {
		anotherChannel <- "cow"
	}()

	select {
	case msgFromMyChannel := <- myChannel: 
	fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <- anotherChannel: 
	fmt.Println(msgFromAnotherChannel)
	}
}

func forSelectChannel() {
	charChannel := make(chan string, 3)

	chars := []string{"a", "b", "c"}

	for _, s := range chars { 
		select {
		case charChannel <- s:
		}
	}

	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}

}

func forSelectChannelAnotherMethod() {
	go func() {
		for {
			select {
			default:
				fmt.Println("DOING WORK")
			}
		}
	}()

	time.Sleep(time.Second * 5)

	// Done Channel will helps to stop the go routines. 
}

// If go routine will happen it called go routine leak 
// Done channel will help to stop the go routine leak 

func doneChannelForSelectChannel(done <-chan bool) {
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				fmt.Println("Do Work")
			}
		}
	}()
}

// spinners for displaying fib value

func spinner(delay time.Duration) {
    for i := 0; i < 20; i++ {
		bar := fmt.Sprintf("[%s%s]", strings.Repeat("=", i), strings.Repeat(" ", 20-i))
		fmt.Printf("\r%s", bar)
		time.Sleep(200 * time.Millisecond)
	}
}

func fib(x int) int {
    if x < 2 {
        return x
    }
    return fib(x-1) + fib(x-2)
}

func returnfibNum() {
	go spinner(100 * time.Millisecond)
    const n = 45
    fibN := fib(n) // slow
    fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

// Echo config


func main() {
	// Go Concurrency Structure.
	// go currencyStructure("1")
	// go currencyStructure("2")
	// go currencyStructure("3")

	time.Sleep(time.Second * 0)

	// channels for the go routine. 
	// channels() 

	// Select primitives for calling the functions. 
	// channelsSelectPrimitive()

	// For select buffered channel 
	// forSelectChannel()
	// forSelectChannelAnotherMethod()

	// Done Channel

	// done := make(chan bool)
	// go doneChannelForSelectChannel(done)
	// time.Sleep(time.Second * 3)
	// close(done)

	
}