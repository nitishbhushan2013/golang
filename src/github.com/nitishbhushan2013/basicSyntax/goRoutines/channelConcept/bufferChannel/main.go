/*
Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel
Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
*/

package main

import "fmt"

// define buffer channel for two value. Now channel would store these two vakue without corrsponding receiver.
// In normal channel behavior, channel would provide value only to the receiver

func main() {
	ch := make(chan int, 2) // make buffer channel of three elements
	ch <- 1
	ch <- 2

	//main.goch <- 3
	fmt.Println(<-ch) // This is not possible with normal channel
	fmt.Println(<-ch)
}
