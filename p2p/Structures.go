package p2p

var NodeAddresses = map[string]struct{}{}

type NodeDataPayload struct {
	Port         string `json:"Port"`
	NodeDataType string `json:"NodeDataType"`
}

type Request struct {
	RequestType string      `json:"RequestType"`
	Payload     interface{} `json:"Payload"`
}
