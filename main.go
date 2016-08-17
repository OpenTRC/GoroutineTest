package main

import (
	"fmt"
)

func Go(channel chan bool) {
	fmt.Println("I am Going!!!")
	channel <- true // send a message
}

func main() {
	channel := make(chan bool)
	go Go(channel)
	<- channel // recieve the message

	go func() {
		sum := 0
		for i:=0; i<10; i++ {
			sum += i
		}
		fmt.Println(sum)
		channel <- true // send a message
	}()
	<- channel // recieve the message

	go func() {
		<- channel // waiting for recieving the message
		sum := 0
		for i:=0; i<10; i++ {
			sum += i
		}
		fmt.Println(sum)
		
	}()
	channel <- true // send a message, start to run the function

	close(channel) // close the channel. it's a great behavior
}