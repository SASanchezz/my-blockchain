package p2p

var NodeAddresses = map[string]struct{}{}

type Request struct {
	RequestType string      `json:"RequestType"`
	Payload     interface{} `json:"Payload"`
}
