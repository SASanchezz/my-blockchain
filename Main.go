package main

import (
	integrationTest "my-blockchain/tests/integration_test"
	"my-blockchain/utils"
)

func main() {
	utils.InitEnvVariables()
	integrationTest.Client_1()
	integrationTest.Client_2()
}
