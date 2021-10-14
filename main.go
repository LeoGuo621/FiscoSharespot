package main

import (
	"fiscoSharespot/routes"
	//	"net/http"
	"github.com/gin-gonic/gin"
)
/*
	./solc-0.4.25 --bin --abi -o ./ ./ResPool.sol
	./abigen --bin ./ResPool.bin --abi ./ResPool.abi --pkg res_pool --type ResPool --out ./ResPool.go
*/
func main() {

<<<<<<< HEAD
func main() {

	routes.Init()
	//caonima := "0xC"
	//n , err := strconv.ParseUint(caonima[2:],16,32)
	//if err != nil{
	//	fmt.Println("cao")
	//}
	//fmt.Println("fuckn=",n)
	r := gin.Default()

=======
	routes.Init()
	//caonima := "0xC"
	//n , err := strconv.ParseUint(caonima[2:],16,32)
	//if err != nil{
	//	fmt.Println("cao")
	//}
	//fmt.Println("fuckn=",n)
	r := gin.Default()

>>>>>>> 49ffd75 (ResPool Contract Finished)
	r.GET("/BlockNumber", routes.BlockNumber)
	r.GET("/NodePeers", routes.NodePeers)
	r.GET("/SyncStatus", routes.SyncStatus)
	r.POST("/BlockByNumber", routes.BlockbyNumber)
	r.GET("/TotTxCount", routes.TotalTransactionCount)
<<<<<<< HEAD
=======
	r.POST("/AddRes", routes.AddRes)
	r.GET("/DeployRes", routes.DeployResContract)
	r.GET("/GetResAll",routes.GetResAll)
	r.POST("/GetRes",routes.GetRes)
	r.POST("/PreBuyRes", routes.PreBuyRes)
	r.POST("/ApRes", routes.ApRes)
	r.POST("/FinalRes", routes.FinalRes)

>>>>>>> 49ffd75 (ResPool Contract Finished)

	//r.GET("/blockbyhash", routes.BlockbyHash)

	r.Run()

}