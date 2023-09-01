package integrationTest

import (
	"fmt"
	clientServer "my-blockchain/p2p/client_server"
	"my-blockchain/utils"
	"os"
	"time"
)

const timeToWait = 5 * time.Second

func IntegrationTest() {
	utils.InitEnvVariables()
	go seedServer()
	time.Sleep(timeToWait)
	go miner()
	time.Sleep(timeToWait)
	client_1()
	client_2()
	time.Sleep(timeToWait * 3)

	conn, err := clientServer.Connect(fmt.Sprintf("%s:%s", os.Getenv("SEED_NODE_HOST"), os.Getenv("SEED_NODE_PORT")))
	if err != nil {
		println("Error connecting to peer:", err.Error())
		return
	}
	lastBlockchain := clientServer.GetBlockchain(conn)
	fmt.Println("\nLast blockchain:", lastBlockchain)

}
