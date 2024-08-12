package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/Shoetan/server"
	"github.com/Shoetan/utils"
)


func main() {

	host := utils.GetEnv("HOST")
	port := utils.GetEnv("PORT")
	address := net.JoinHostPort(host, port)

	command := strings.Join(os.Args[1:3], " ") 

	fmt.Println(command)

	switch command {
	case "server start":
		var err error
		listener, err := server.StartTcpServer("tcp", address)
	
		if err != nil {
			log.Fatalf("Could not start a tcp server:%v", err.Error())
		}else{
			fmt.Printf("The tcp server is running on address %v\n", address)
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

	case "server connect":
		conn, err := server.ConnectToTcpServer("tcp", address)

		if err != nil {
			fmt.Println("Could not connect to server")
		}

		defer conn.Close()

	default:
		fmt.Println("Unknow command", command)	
		os.Exit(1)
	}
}