package client

import (
	"encoding/json"
	"fmt"
	"my-blockchain/p2p"
	"net"
)

func Connect(target string) (net.Conn, error) {
	conn, err := net.Dial("tcp", target)
	if err != nil {
		return nil, err
	}
	p2p.NodeAddresses[conn.RemoteAddr().String()] = struct{}{}

	return conn, nil
}

func sendMessge(conn net.Conn, request *p2p.Request) {
	encoder := json.NewEncoder(conn)

	err := encoder.Encode(request)
	if err != nil {
		fmt.Println("encode.Encode error: ", err)
	}

	fmt.Println("Sent", request.RequestType, "to", conn.RemoteAddr())
}

func broadcastMessge(request *p2p.Request) {
	for address := range p2p.NodeAddresses {
		conn, err := Connect(address)
		if err != nil {
			fmt.Println("Error connecting to peer:", err)
			continue
		}
		sendMessge(conn, request)
	}
}
