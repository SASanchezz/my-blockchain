package clientServer

import (
	"encoding/json"
	"fmt"
	"my-blockchain/mynet"
	"my-blockchain/p2p"
	"net"
	"os"
)

func StartListening(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting listener:", err)
		os.Exit(1)
	}
	// defer listener.Close()
	myIP := mynet.GetMyLocalAddress()
	fmt.Printf("Listening on: %s:%s", myIP.String(), port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// defer conn.Close()

	fmt.Println("Connection established with:", conn.RemoteAddr())
	var request *p2p.Request
	decoder := json.NewDecoder(conn)
	fmt.Println("Received:", decoder)
	err := decoder.Decode(&request)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	processRequestType(conn, request)
}
