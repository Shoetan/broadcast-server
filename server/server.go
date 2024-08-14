package server

import (
	// "fmt"
	"fmt"
	"net"
)

func StartTcpServer(network string, address string) (net.Listener, error) {
	listener, err := net.Listen(network, address)

	if err != nil {
		return nil, err
	}

	return listener, err
}

func ConnectToTcpServer(network string, address string) (net.Conn, error) {
	connection, err := net.Dial(network, address)

	if err != nil {
		return nil, err
	}

	return connection, err
}

func DisconnectFromTcpServer(conn net.Conn)  {
	conn.Close()
}

func SendMessageToServer(conn net.Conn, message []byte)  {
	_, err := conn.Write(message)
	
	if err != nil {
		fmt.Println("Could not send message to the server")
	}
}

// Write a function to send a message to the server from the connected client 