package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"bufio"
	"strings"

	"github.com/Shoetan/server"
	"github.com/Shoetan/utils"
)


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
			log.Fatalf("Could not start TCP server:%v", err.Error())
		}else{
			fmt.Printf("TCP server irunning on address 🌐 %v\n", address)
		}

		defer listener.Close()

			// keep server running in and endless loop 
		for {
			conn, err := listener.Accept()
			if err != nil {
			fmt.Printf("Error accepting new connections 🛑 %v", err.Error())
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
			fmt.Println("Connected to TCP server ✅ :",clientConnection.RemoteAddr().String() )
			utils.SaveConnectionDetails(clientConnection)
		} else {
			fmt.Println("Failed to establish connection 🛑")
		}

		exitChan := make(chan bool)

		go func ()  {
			
			reader := bufio.NewReader(os.Stdin) 
			for {
				fmt.Printf("What would you like to do now that you are connected? 😁\n")
				fmt.Printf("1. To send message to server: 1 <add message> 💬\n")
				fmt.Printf("2. Exit server 🗑 \n")
	
				fmt.Println("Enter choice ")
	
				choice, _ := reader.ReadString('\n')
				choice = strings.TrimSpace(choice)

				parts := strings.Fields(choice) // seperate the input from the command line into parts

				var message string 

				if len(parts) >= 2 {
					choice = parts[0] // assign various parts into  respective variables
					message = parts[1]

				}
	
				switch choice {
				case "1":
					fmt.Println("You want to send a message 📁")
					utils.SendMessage(clientConnection, []byte(message))
	
				case "2":
					fmt.Println("You want to exit the server 🗑")
					exitChan <- true
					return
				}
			}
		}()
		
		value:= <- exitChan

		if value {
			fmt.Println("Disconnected from the TCP server 😔")
			clientConnection.Close()
		}

	default:
		fmt.Println("Unknow command 😭", command)	
		os.Exit(1)

	}
}


