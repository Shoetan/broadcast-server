package utils

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
)


func GetEnv(key string) string{
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error Loading env file %v", err.Error())
	}

	env := os.Getenv(key)

	return env
}

// handle connections to the TCP server

func HandleConnection(conn net.Conn)  {
	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("Client %s connected \n", clientAddr)
}