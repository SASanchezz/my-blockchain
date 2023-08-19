package server

import (
	"encoding/json"
	"fmt"
	"my-blockchain/p2p"
	"net"
	"os"
)

func StartListening(host string) {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		fmt.Println("Error starting listener:", err)
		os.Exit(1)
	}
	// defer listener.Close()

	fmt.Println("Listening on:", host)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		p2p.NodeAddresses[conn.RemoteAddr().String()] = struct{}{}
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// defer conn.Close()

	fmt.Println("Connection established with:", conn.RemoteAddr())
	var request *p2p.Request
	decoder := json.NewDecoder(conn)
	fmt.Println("Received1:", decoder)
	err := decoder.Decode(&request)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	processRequestType(conn, request)
}
