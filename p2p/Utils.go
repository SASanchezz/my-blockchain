package p2p

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"reflect"
)

func ConvertMapToObject(m map[string]interface{}, s interface{}) {
	jsonData, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonData, &s)
	if err != nil {
		log.Fatal(err)
	}
}

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

func GetRandomNodeAddress() string {
	keys := reflect.ValueOf(NodeAddresses).MapKeys()
	return keys[rand.Intn(len(keys))].Interface().(string)
}
