package p2p

import (
	"encoding/json"
	"fmt"
	"my-blockchain/blockchain"
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
		nodeAddresses[conn.RemoteAddr().String()] = struct{}{}
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// defer conn.Close()

	fmt.Println("Connection established with:", conn.RemoteAddr())
	var request *Request
	decoder := json.NewDecoder(conn)
	fmt.Println("Received1:", decoder)
	err := decoder.Decode(&request)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	processRequestType(conn, request)
}

func syncNodeAddresses(myNodeAddresses *map[string]struct{}, incomingNodeAddresses *map[string]struct{}) {
	for na := range *incomingNodeAddresses {
		(*myNodeAddresses)[na] = struct{}{}
	}
}

func processRequestType(conn net.Conn, request *Request) {
	switch request.RequestType {
	case NewBlockType:
		var block *blockchain.Block = &blockchain.Block{}
		ConvertMapToObject(request.Payload.(map[string]interface{}), block)
		handleNewBlock(block)
		sendProcessedBlock(conn, block)

	case ProcessedBlockType:
		var block *blockchain.Block = &blockchain.Block{}
		ConvertMapToObject(request.Payload.(map[string]interface{}), block)
		fmt.Println("Received processed block2", block)
		handleProcessedBlock(block)

	case NodeAddressesType:
		fmt.Println("Received node addresses")
		nAddreses := request.Payload.(*map[string]struct{})
		handleNodeAddresses(nAddreses)
	default:
		fmt.Println("Received unknown request type")
	}
}
