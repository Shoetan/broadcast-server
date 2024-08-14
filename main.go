package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	// "os/signal"
	// "syscall"

	"github.com/Shoetan/server"
	"github.com/Shoetan/utils"
)

var clientConnection net.Conn

func main() {

	host := utils.GetEnv("HOST")
	port := utils.GetEnv("PORT")
	address := net.JoinHostPort(host, port)

	command := strings.Join(os.Args[1:3], " ") // get command from terminal 

	switch command {
	case "start server":
		var err error
		listener, err := server.StartTcpServer("tcp", address)
	
		if err != nil {
			log.Fatalf("Could not start a tcp server:%v", err.Error())
		}else{
			fmt.Printf("The tcp server is running on address.. %v\n", address)
		}

		defer listener.Close()

			// keep server running in and endless loop 
		for {
			conn, err := listener.Accept()
			if err != nil {
			fmt.Printf("Error accepting new connections %v", err.Error())
			continue
			}
			
			go utils.HandleConnection(conn)
	}
	
	case "connect server":
		clientConnection, err := server.ConnectToTcpServer("tcp", address)
		if err != nil {
			fmt.Println("Could not connect to server")
		}
		fmt.Println("Connected to TCP server:",clientConnection.RemoteAddr().String() )
		fmt.Println("Client connected to the TCP server", clientConnection.LocalAddr().String())

	case "send message":
		fmt.Println("Client sending message:", clientConnection)
			
	default:
		fmt.Println("Unknow command", command)	
		os.Exit(1)

	}
}

	// Store list of connected address
	//Check if the address is in the list of connected address
	// Disconnect the address from the TCP server
	// Do a disconnect from TCP server function in the server pkg
	// Bring the function here and implement it 
	// Import the client pool here and check if the connection is in the client pool
	// if yes close the connection, if no tell the user is not connected
	// if clientConnection == {
	// 	server.DisconnectFromTcpServer(clientConnection)
	// }
	// Disconnect server is to remove the present connection from the server and close all 
	// user should disconnect from the server by sending a message like QUIT
