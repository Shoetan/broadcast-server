package server

import (
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