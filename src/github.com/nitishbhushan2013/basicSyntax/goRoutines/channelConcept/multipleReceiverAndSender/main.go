package main

import "fmt"

/* channel is the way to send or recieve value to channel type. Generally, it helps us to communicate
between routine and represented by <-

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
So channel send/receive operations block until the other side is ready:

1. A send operation on a channel blocks until a receiver is available for the same channel: if there’s no recipient for the value on ch, no other value can be put in the channel. And the other way around: no new value can be sent in ch when the channel is not empty! So the send operation will wait until ch becomes available again.

2. A receive operation for a channel blocks until a sender is available for the same channel: if there is no value in the channel, the receiver blocks.

*/

func PrintNumber(list []int, c chan int) {
	for _, i := range list {
		c <- i
	}
}

// RecieveChannelValue : This function would iterate for each of the number in the list to receive the channel value. The idea is to loop through exact number of time to receive channel value.

func RecieveChannelValue(list []int, c chan int) {
	for _, i := range list {
		fmt.Println(i)
		num := <-c
		fmt.Println("num -->", num)
	}

}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	chan1 := make(chan int)
	go PrintNumber(nums, chan1)
	RecieveChannelValue(nums, chan1)
}
