/*
channel is used to represent a communication medium between sender and receiver.
each channel would corrosponds between a pair of sender-receiver.
 step1: define channel
 step2 : send channel to the calling function as a parameter to suppplu the value
 step3: use same channel in receiver to receive value that has been supplied by supplier
 step4: invoke the supplier with go command to make it routine.

*/

package main

import (
	"fmt"
	//"time"
)

func main() {
	ch := make(chan int) // drefine channel
	list := []int{1, 2, 3, 4, 5}
	go Printvalues(list, ch) // pass channel to sender
	value := <-ch            // value would be available in channel when supplier supplies the value to channel.  receive value from channel

	fmt.Println("received value from channel-->", value) //1
}

func Printvalues(numbers []int, c chan int) { // define channel with its type
	// perform business logic
	for _, num := range numbers {
		fmt.Println("num -->", num)
		c <- num
	}
}

// explanation  -Printvalues() prints 1, supplies value to channel, then iterate again and print 2
// By second iteration, channel has delivered value to receiver in the main function and since main ()
// has completed its opertaion, its not waiting for any more input, it existed the program. Thus Printvalues() also exited and for loop which were in the middle of iteration coould not print other numbers.
