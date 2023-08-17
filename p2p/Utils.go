package p2p

import (
	"encoding/json"
	"log"
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
