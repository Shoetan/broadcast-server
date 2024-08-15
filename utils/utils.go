package utils

import (
	"fmt"
	"log"
	"net"
	"os"
	// "strings"
	"sync"

	"github.com/joho/godotenv"
)

var (
	ClientPool = make(map[string]net.Conn)
	poolMutex  sync.Mutex
) 

func HandleConnection(conn net.Conn)  {
	clientAddr := conn.RemoteAddr().String()

	poolMutex.Lock()
	ClientPool[clientAddr] = conn
	poolMutex.Unlock()
	
	buffer := make([]byte, 1024)  
	n, _ := conn.Read(buffer)
	message := string(buffer[:n])

	fmt.Printf("Received message from connection: %s\n", message)
}


func GetEnv(key string) string{
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error Loading env file %v", err.Error())
	}

	env := os.Getenv(key)

	return env
}

func SaveConnectionDetails(conn net.Conn)  {
	// save the connection to a file 
	err := os.WriteFile("connection.txt", []byte(conn.LocalAddr().String()), 0644)

	if err != nil {
		fmt.Println("Error saving connection details ❌")
	}
	
}

func LoadConnectionDetails() string {
	data, err := os.ReadFile("connection.txt")

	if err != nil {
		fmt.Println("Could not load connection details")
	}

	return string(data)

}

func SendMessage(conn net.Conn, message []byte)  {
	_, err := conn.Write(message)

	if err != nil {
		fmt.Println("Could not send message to the server")
	} else {
		fmt.Println("Message sent ...")
	}
}

