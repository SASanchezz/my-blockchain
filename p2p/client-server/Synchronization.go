package clientServer

import (
	"fmt"
	"my-blockchain/blockchain"
	"my-blockchain/p2p"
)

func SyncBlockchain(host string, port string) {
	conn, err := Connect(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		println("Error connecting to peer:", err.Error())
		return
	}
	request := &p2p.Request{
		RequestType: p2p.GetSyncBlockchainType,
	}
	sendMessge(conn, request)

	response := p2p.GetResponse(&conn)
	var blockchain *blockchain.Blockchain = &blockchain.Blockchain{}
	p2p.ConvertMapToObject(response.Payload.(map[string]interface{}), blockchain)
	HandleBlockchain(blockchain)
}

func SyncNodeAddresses(host string, port string) {
	conn, err := Connect(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		println("Error connecting to peer:", err.Error())
		return
	}
	request := &p2p.Request{
		RequestType: p2p.GetSyncNodeAddressesType,
	}
	sendMessge(conn, request)

	response := p2p.GetResponse(&conn)
	nAddreses := map[string]struct{}{}
	p2p.ConvertMapToObject(response.Payload.(map[string]interface{}), nAddreses)
	HandleNodeAddresses(&nAddreses)
}
