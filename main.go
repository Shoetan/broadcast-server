package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Shoetan/server"
	"github.com/Shoetan/utils"
)




func main() {

	args := os.Args


// Todo
// 1. Get command from command line or terminal
// 2. Figure out if the client has to start the server or autiomactically connects to the TCP server on remote
// 3. Check command line from the terminal if the command is to connect to tcp server online or local

	host := utils.GetEnv("HOST")
	port := utils.GetEnv("PORT")
	address := net.JoinHostPort(host, port)

	listener, err := server.StartTcpServer("tcp", address)

	if err != nil {
		log.Fatalf("Could not start a tcp server:%v", err.Error())
	}else{
		fmt.Printf("The tcp server is running on address %v\n", address)
	}

	// if len(args) != 0 {
	// 	server.ConnectToTcpServer("tcp", address)
	// }

	fmt.Println(args)

	// keep server running in and endless loop 
	for {
		 conn, err := listener.Accept()

		 if err != nil {
			fmt.Printf("Error accepting new connections %v", err.Error())
			continue
		 }

		 go utils.HandleConnection(conn)
	}

}