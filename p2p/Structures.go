package p2p

import "my-blockchain/blockchain"

var nodeAddresses = map[string]struct{}{}

type Request struct {
	RequestType string      `json:"RequestType"`
	Payload     interface{} `json:"Payload"`
}

type NewBlockRequest struct {
	RequestType string            `json:"RequestType"`
	Payload     *blockchain.Block `json:"Payload"`
}

type ProcessedBlockRequest struct {
	RequestType string            `json:"RequestType"`
	Payload     *blockchain.Block `json:"Payload"`
}
