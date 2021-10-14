package main

import (
	"fiscoSharespot/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)
/*
	./solc-0.4.25 --bin --abi -o ./store ./store/Store.sol
	./abigen --bin ./store/Store.bin --abi ./store/Store.abi --pkg store --type Store --out ./store/Store.go
*/
func main(){
	engine := gin.Default()
	engine.GET("/ping", routes.PingHandler)
	engine.GET("/blockNumber", routes.BlockNumber)
	engine.POST("/blockByNumber", routes.BlockByNumber)
	engine.POST("/users", routes.AddUser)
	if err := engine.Run(":8080"); err != nil {
		fmt.Printf("Server failed to start, err: %v", err)
	}

	//fmt.Println(number)
	//
	//if err != nil {
	//	return
	//}
	//input := "Store deployment 1.0"
	//address, tx, instance, err := store.DeployStore(conn.GetTransactOpts(), conn, input)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("contract address: ", address.Hex()) // the address should be saved, will use in next example
	//fmt.Println("transaction hash: ", tx.Hash().Hex())
	//
	////load the contract
	////contractAddress := common.HexToAddress("contract address in hex String")
	////instance, err := store.NewStore(contractAddress, conn)
	////if err != nil {
	////	log.Fatal(err)
	////}
	//
	//fmt.Println("================================")
	//storeSession := &store.StoreSession{Contract: instance, CallOpts: *conn.GetCallOpts(), TransactOpts: *conn.GetTransactOpts()}
	//
	//version, err := storeSession.Version()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("version :", version) // "Store deployment 1.0"
	//
	//// contract write interface demo
	//fmt.Println("================================")
	//key := [32]byte{}
	//value := [32]byte{}
	//copy(key[:], []byte("foo"))
	//copy(value[:], []byte("bar"))
	//
	//tx, receipt, err := storeSession.SetItem(key, value)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	//fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())
	//
	//// read the result
	//result, err := storeSession.Items(key)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("get item: " + string(result[:])) // "bar"
}
