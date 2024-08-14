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

	

	if len(os.Args) < 2 {
		fmt.Println("Expected 'start server', 'connect server', or 'send message' subcommands")
		os.Exit(1)
	}

	command := strings.Join(os.Args[1:], " ") // get command from terminal 


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

		if clientConnection != nil {
			fmt.Println("Connected to TCP server:",clientConnection.RemoteAddr().String() )
			utils.SaveConnectionDetails(clientConnection)
		} else {
			fmt.Println("Failed to establish connection")
		}

	case "send message":
		// get connection details from stored file 
		client := utils.LoadConnectionDetails()
		fmt.Println(client)
			
	default:
		fmt.Println("Unknow command", command)	
		os.Exit(1)

	}
}


