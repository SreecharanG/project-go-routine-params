package main

import (
	"fmt"
	"time"
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

}