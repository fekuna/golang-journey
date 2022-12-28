package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		// time.Sleep(2 * time.Second)
		channel <- "Alfan Almunawar"
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <-channel
	fmt.Println(data)

	// time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	channel <- "Alfan Almunawar"
	fmt.Println("Selesai Mengirim Response")
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel

	fmt.Println("TestChannelAsParameter", data)
}

// -- IN OUT CHANNEL
// Cahnnel Only Receive data
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Alfan Almunawar"
}

// Channel Only Send data
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// Buffered channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 1)

	go func() {
		channel <- "Alfan"
		channel <- "Almunawar"
		channel <- "Pan"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}
