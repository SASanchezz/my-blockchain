package p2p

import (
	"encoding/json"
	"fmt"
	"my-blockchain/blockchain"
	"net"
)

func Connect(target string) net.Conn {
	conn, err := net.Dial("tcp", target)
	if err != nil {
		fmt.Println("Error connecting to peer:", err)
		panic(err) //TODO handle this error more gracefully
	}
	nodeAddresses[conn.RemoteAddr().String()] = struct{}{}

	return conn
}

func sendMessge(conn net.Conn, request *Request) {
	encoder := json.NewEncoder(conn)

	err := encoder.Encode(request)
	if err != nil {
		fmt.Println("encode.Encode error: ", err)
	}

	fmt.Println("Sent message", request)
}

func broadcastMessge(request *Request) {
	for address := range nodeAddresses {
		conn := Connect(address)
		sendMessge(conn, request)
	}
}

func SendGetSyncBlockchain(conn net.Conn) {
	request := &Request{
		RequestType: GetSyncBlockchainType,
	}
	sendMessge(conn, request)
}

func SendGetSyncNodeAddresses(conn net.Conn) {
	request := &Request{
		RequestType: GetSyncNodeAddressesType,
	}
	sendMessge(conn, request)
}

func SendNewBlock(conn net.Conn, block *blockchain.Block) {
	request := &Request{
		RequestType: NewBlockType,
		Payload:     block,
	}
	sendMessge(conn, request)
}

func sendProcessedBlock(conn net.Conn, block *blockchain.Block) {
	request := &Request{
		RequestType: ProcessedBlockType,
		Payload:     block,
	}
	sendMessge(conn, request)
}

func SendNodeAddresses(conn net.Conn, nodeAddresses *map[string]struct{}) {
	request := &Request{
		RequestType: NodeAddressesType,
		Payload:     nodeAddresses,
	}
	sendMessge(conn, request)
}

func SendBlockchain(conn net.Conn, blockchain blockchain.Blockchain) {
	request := &Request{
		RequestType: NodeAddressesType,
		Payload:     BlockchainType,
	}
	sendMessge(conn, request)
}

func BroadcastProcessedBlock(block *blockchain.Block) {
	request := &Request{
		RequestType: ProcessedBlockType,
		Payload:     block,
	}
	broadcastMessge(request)
}

func BroadcastNewBlock(block *blockchain.Block) {
	request := &Request{
		RequestType: NewBlockType,
		Payload:     block,
	}
	broadcastMessge(request)
}
