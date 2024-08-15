## TCP SERVER
A simple TCP-based server for local communication. Allows multiple clients to connect and exchange messages within a local network or over the internet if hosted 😂

## Demo

## Prerequisites
* Go (version 1.22 or later recommended)

## Getting Started
* Clone repository `git@github.com:Shoetan/broadcast-server.git`
* cd into the cloned repository

## Commands

* Start the TCP server on a local address `go run main.go start server`
* Connect a client to the TCP server `go run main.go connect server`
* Stop the TCP server `ctrl + c`

## Sub commands

* Send a message to the TCP server `1 <message>`
* Disconnect from the TCP server `2`


  
