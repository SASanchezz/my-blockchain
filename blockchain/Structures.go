package blockchain

var LocalBlockchain *Blockchain = NewBlockchain()

type Block struct {
	From              string `json:"From"`              // the sender
	To                string `json:"To"`                // the receiver
	Amount            int64  `json:"Amount"`            // the amount to be sent
	Data              string `json:"Data"`              // the data to be sent
	Timestamp         int64  `json:"Timestamp"`         // the time when the block was created
	PreviousBlockHash []byte `json:"PreviousBlockHash"` // the hash of the previous block
	BlockHash         []byte `json:"BlockHash"`         // the hash of the current block
	Nonce             int64  `json:"Nonce"`             // nonce to change hash of the current block
}

type Blockchain struct {
	Chain []*Block
}
