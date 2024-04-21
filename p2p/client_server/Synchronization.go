package clientServer

import (
	"my-blockchain/blockchain"
	"my-blockchain/p2p"
	"my-blockchain/utils"
	"net"
	"net/url"
)

func SendMyData(url *url.URL, nodeData *p2p.NodeDataPayload) {
	conn, err := Connect(url.Host)
	if err != nil {
		println("Error connecting to peer:", err.Error())
		return
	}
	SendNodeData(conn, nodeData)
}

func SyncBlockchain(url *url.URL) {
	conn, err := Connect(url.Host)
	if err != nil {
		println("Error connecting to peer:", err.Error())
		return
	}
	blockchain := GetBlockchain(conn)

	HandleBlockchain(blockchain)
}

func GetBlockchain(conn net.Conn) *blockchain.Blockchain {
	request := &p2p.Request{
		RequestType: p2p.GetSyncBlockchainType,
	}
	sendMessge(conn, request)

	response := p2p.GetResponse(&conn)
	var blockchain *blockchain.Blockchain = &blockchain.Blockchain{}
	utils.ConvertMapToObject(response.Payload.(map[string]interface{}), blockchain)

	return blockchain
}

func SyncNodeAddresses(url *url.URL) {
	conn, err := Connect(url.Host)
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
	utils.ConvertMapToObject(response.Payload.(map[string]interface{}), nAddreses)
	HandleNodeAddresses(&nAddreses)
}
