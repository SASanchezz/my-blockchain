package integrationTest

import (
	"fmt"
	"my-blockchain/p2p"
	clientServer "my-blockchain/p2p/client_server"
	"my-blockchain/utils"
	"time"
)

const timeToWait = 5 * time.Second

func IntegrationTest() {
	utils.InitEnvVariables()
	go seedServer()
	time.Sleep(timeToWait)
	go miner()
	time.Sleep(timeToWait)
	Client_1()
	Client_2()
	time.Sleep(timeToWait * 3)

	conn, err := clientServer.Connect(p2p.GetSeedNodeUrl().Host)
	if err != nil {
		println("Error connecting to peer:", err.Error())
		return
	}
	lastBlockchain := clientServer.GetBlockchain(conn)
	fmt.Println("\nLast blockchain:", lastBlockchain)

}
