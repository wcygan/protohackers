package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"
)

const (
	port = ":8080"
)

// https://protohackers.com/problem/0
func main() {
	startServer()
}

func startServer() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	slog.Info(fmt.Sprintf("Server is running on: %s", listener.Addr()))
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// Spawn a goroutine to handle the connection
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		// Read bytes
		bytesIn, err := conn.Read(buffer)
		if err != nil {
			slog.Info(fmt.Sprintf("Error: %s", err))
			return
		}

		// Write the bytes back unmodified
		bytesOut, err := conn.Write(buffer)
		if err != nil {
			slog.Info(fmt.Sprintf("Error: %s", err))
			return
		}

		if bytesOut != bytesIn {
			slog.Info(fmt.Sprintf("%s bytes read, %s bytes written: %s", bytesIn, bytesOut))
		}
	}
}
