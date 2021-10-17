package main

import (
	"fiscoSharespot/routers"
	"log"
)
/*
	./solc-0.4.25 --bin --abi -o ../sols/ ../sols/UserManagement.sol
	./abigen.exe --bin ../sols/UserManagement.bin --abi ../sols/UserManagement.abi --pkg user_management --type UserManagement --out ../user_management/user_management.go
*/

func main(){
	r := routers.InitRouter()
	err := r.Run(":8080")
	if err != nil {
		log.Printf("start dapp failed %s", err)
	}
}
