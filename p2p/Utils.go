package p2p

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/url"
	"os"
	"reflect"
)

func GetResponse(conn *net.Conn) *Request {
	for {
		var request *Request
		decoder := json.NewDecoder(*conn)
		err := decoder.Decode(&request)
		if err != nil {
			continue
		}

		fmt.Println("Received:", request.RequestType, "from", (*conn).RemoteAddr())
		return request
	}
}

func GetSeedNodeUrl() *url.URL {
	return GetUrl(os.Getenv("SEED_NODE_HOST"), os.Getenv("SEED_NODE_PORT"))
}

func GetUrl(address string, port string) *url.URL {
	url, _ := url.Parse(fmt.Sprintf("%s:%s", address, port))
	url.Host = fmt.Sprintf("%s:%s", address, port)
	return url
}

func GetRandomNodeAddress() string {
	print("NodeAddresses", NodeAddresses)
	keys := reflect.ValueOf(NodeAddresses).MapKeys()
	return keys[rand.Intn(len(keys))].Interface().(string)
}
