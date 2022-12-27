package main

import (
	"fmt"
	"testing"
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
