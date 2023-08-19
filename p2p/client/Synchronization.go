package client

import (
	"my-blockchain/blockchain"
	"my-blockchain/p2p"
	"net"
)

func SendGetSyncBlockchain(conn net.Conn) {
	request := &p2p.Request{
		RequestType: p2p.GetSyncBlockchainType,
	}
	sendMessge(conn, request)

	response := p2p.GetResponse(&conn)
	var blockchain *blockchain.Blockchain = &blockchain.Blockchain{}
	p2p.ConvertMapToObject(response.Payload.(map[string]interface{}), blockchain)
	p2p.HandleBlockchain(blockchain)
}

func SendGetSyncNodeAddresses(conn net.Conn) {
	request := &p2p.Request{
		RequestType: p2p.GetSyncNodeAddressesType,
	}
	sendMessge(conn, request)

	response := p2p.GetResponse(&conn)
	nAddreses := response.Payload.(*map[string]struct{})
	p2p.HandleNodeAddresses(nAddreses)
}
