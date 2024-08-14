package utils

import (
	"fmt"
	"log"
	"net"
	"os"
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

	fmt.Println("clientPool:",ClientPool)
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
		fmt.Println("Error saving connection details")
	}
	
}

func LoadConnectionDetails() string {
	data, err := os.ReadFile("connection.txt")

	if err != nil {
		fmt.Println("Could not load connection details")
	}
	return string(data)
}

