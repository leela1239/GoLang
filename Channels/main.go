package main

import "fmt"

func main() {
	fmt.Println("Goroutines")
	/*
		   	chnl := make(chan int)
		   	chnl <- 1
		   	res := <-chnl
		   	fmt.Println(res) //Gives panic as deadlock because

		   ch := make(chan int) creates an unbuffered channel.
		   ch <- 1 tries to send a value into the channel.
		   For an unbuffered channel, send and receive operations must happen simultaneously; the sender blocks until a receiver is ready.
		   Since there is no goroutine receiving from the channel yet, the send (ch <- 1) blocks forever.
		   Your program deadlocks because the send blocks and never reaches the receive statement (res := <-ch).

		   How to fix it ?

		   1. Using goroutine to receive concurrently

		   ch := make(chan int)
			go func() {
				res := <-ch
				fmt.Println(res)
			}()

			ch <- 1
			close(ch)

		   2.Using Budffered Channel

			ch := make(chan int, 1) // buffered channel with capacity 1
			ch <- 1         // does NOT block, because buffer can hold one item
			res := <-ch     // receive value
			fmt.Println(res)
			close(ch)



	*/

	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch) // need to close the channel because
	for k := range ch {
		fmt.Println(k)
	}
	/* The sending goroutine sends 5 values.
	The receiving goroutine reads those 5 values.
	But then it expects more (because channel is still open).
	Since no more values are sent and the channel is open, it blocks forever.
	*/

	// =====================================================================================================================
	//Channel follows FIFO

	ch1 := make(chan int, 2)
	ch1 <- 1
	fmt.Println(ch1)
	ch1 <- 2
	res = <-ch1
	fmt.Println(res)

	/*

	   Why is res equal to 1 and not 2?
	   Channels are FIFO queues (First In, First Out).

	   When you send 1 into the buffered channel, it gets placed first in the buffer.

	   Then you send 2, which goes second in the buffer.

	   When you do <-ch (receive), it retrieves the oldest (first) value in the channel buffer, which is 1.

	   The next receive would give you 2.

	*/

}
