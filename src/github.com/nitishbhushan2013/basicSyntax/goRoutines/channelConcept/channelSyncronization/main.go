/*
The idea of channel can be extended to bring syncronixation between different goroutine module
start a goroutine, give it a channel, block the program execution until channel notify back with value, indicating the completion of goroutine.
*/

package main

import (
	"fmt"
	"time"
)

/*
case1 : if you remove ALL the receiver from the main program THEN program will exit immediately
case 2: if you comment ANY of  the suppplier (the palce where channel would receive the value in the routine
function) THEN program would result an error, as channel would be waiting for the supplier value and thus receive would be waiting for the channel value
*/
func main() {
	ch1 := make(chan string)
	go Routine1(ch1)
	fmt.Println(<-ch1)

	go Routine2(ch1)
	fmt.Println(<-ch1)

	go Routine3(ch1)
	fmt.Println(<-ch1)

	fmt.Println("All completed ")
}

func Routine1(ch chan string) {
	fmt.Println("In the Routine1")
	time.Sleep(time.Millisecond * 10) // each worker would sleep for 10 ms
	ch <- "Hello"
}

func Routine2(ch chan string) {
	fmt.Println("In the Routine2")
	time.Sleep(time.Millisecond * 10) // each worker would sleep for 10 ms
	ch <- "Nitish"
}

func Routine3(ch chan string) {
	fmt.Println("In the Routine3")
	time.Sleep(time.Millisecond * 10) // each worker would sleep for 10 ms
	ch <- "Bhushan"
}
