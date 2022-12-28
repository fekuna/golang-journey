package main

import (
	"fmt"
	"strconv"
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

// --- Buffered channel
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

// --- RANGE CHANNEL

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

//  --- SELECT CHANNEL

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Sedang Menunggu..")
		}

		if counter == 2 {
			break
		}
	}

}
