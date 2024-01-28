package golanggoroutine

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
		time.Sleep(1 * time.Second)
		channel <- "Eko Kurniawan Khennedy"
		fmt.Println("Selesai mengirim data")
	}()

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Eko Kurniawan Khennedy"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go GiveMeResponse(channel)
	data := <- channel
	fmt.Println(data)
}

func OnlyIn(channel chan<- string) {
	channel <- "Eko Kurniawan Khanedy"
	time.Sleep(2 * time.Second)
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second);
	defer close(channel)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Eko"
		channel <- "Kurniawan"
		channel <- "Khannedy"
	}()

	go func () {
		fmt.Println(<- channel)
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func () {
		for i := 0; i < 10; i++ {
			channel <- "Perulanang ke:" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Mnerima data: ", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)
	
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel2", data)
			counter++
		}
		if counter == 3 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)
	
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}
		if counter == 3 {
			break
		}
	}
}