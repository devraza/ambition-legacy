package main

import (
	"log"
	"net"
	"os"
)

const (
	serverAddress = "localhost:7764"
)

func main() {
	log.Println("Ambition going strong at", serverAddress)

	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalln("Failed to initialise TCP listener", err)
	}
	defer listener.close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}

		// Concurrency FTW
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// TODO implement actual server. Waiting on frontend for this
}
