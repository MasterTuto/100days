package main

import (
	"fmt"
	"log"
)

func logSelectedChannel(c1 chan string, c2 chan string, c3 chan string) {
	select {
	case message := <-c1:
		log.Printf("[Channel 1] %s", message)
	case message := <-c2:
		log.Printf("[Channel 2] %s", message)
	case message := <-c3:
		log.Printf("[Channel 3] %s", message)
	}
}

func logChannelMessage(channel chan string, channelName string) {
	message := <-channel
	log.Printf("[%s] %s", channelName, message)
}

func writeMessage(channel chan string, message string) {
	channel <- message
}

func main() {

	messageChannel := make(chan string)
	for i := 0; i < 10; i++ {

		go writeMessage(messageChannel, "Message "+fmt.Sprint(i))
	}

	for i := 0; i < 10; i++ {
		go logChannelMessage(messageChannel, "Message Channel")
	}

	bufferedChannel := make(chan string, 3)
	for i := 0; i < 10; i++ {
		go writeMessage(bufferedChannel, "Message "+fmt.Sprint(i))
	}

	go func() {
		for message := range bufferedChannel {
			log.Printf("[Buffered Channel] %s", message)
		}
	}()

	c1 := make(chan string, 10)
	c2 := make(chan string, 10)
	c3 := make(chan string, 10)

	go func() {
		for i := 0; i < 10; i++ {
			writeMessage(c1, "Message to c1: "+fmt.Sprint(i+1))
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			writeMessage(c2, "Message to c2: "+fmt.Sprint(i+1))
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			writeMessage(c3, "Message to c3: "+fmt.Sprint(i+1))
		}
	}()

	for i := 0; i < 10; i++ {
		logSelectedChannel(c1, c2, c3)
	}

	log.Println("Done")
}
