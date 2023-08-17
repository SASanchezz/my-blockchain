package main

import "my-blockchain/p2p"

func main() {
	p2p.StartListening("localhost:3001")
}
