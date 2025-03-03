package main

import (
	"fmt"
	"time"
)

func main() {
	s := &Server{
		msgch: make(chan Message),
		quitch: make(chan struct{}),
	}

	go s.StartAndListen()


	go func() {
		time.Sleep(2 * time.Second)
		sendMessageToServer(s.msgch, "Hello World")
	}()
	
	go func() {
		time.Sleep(4 * time.Second)
		quitServer(s.quitch)
	}()

	time.Sleep(5 * time.Second)
}

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch chan Message
	quitch chan struct{}
}

func (s *Server) StartAndListen() {
	for {
		select {
		case msg := <- s.msgch:
			fmt.Printf("Receive message from: %s payload: %s\n", msg.From, msg.Payload)
		case <- s.quitch:
			fmt.Println("Quit. Shutting down...")
			return
		}
	}
}

func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From: "Client 1",
		Payload: payload,
	} 
	msgch <- msg
}

func quitServer(quitch chan struct{}) {
	close(quitch)
}
