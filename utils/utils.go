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

// handle connections to the TCP server
// in this connection find a way to store the connect clients 
// Also find a way to remove clients from the list of connected clients
// Remember when you close the main server all the connected clients are closed