package main

import (
	"my-blockchain/node"
	"my-blockchain/p2p"
	"my-blockchain/utils"
	"os"
)

func main() {
	utils.InitEnvVariables()
	node.Run(p2p.SeedNode, os.Getenv("SERVER_PORT"))
}
