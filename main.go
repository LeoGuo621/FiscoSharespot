package main

import (
	"fiscoSharespot/routes"
	//	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {

	routes.Init()
	//caonima := "0xC"
	//n , err := strconv.ParseUint(caonima[2:],16,32)
	//if err != nil{
	//	fmt.Println("cao")
	//}
	//fmt.Println("fuckn=",n)
	r := gin.Default()

	r.GET("/BlockNumber", routes.BlockNumber)
	r.GET("/NodePeers", routes.NodePeers)
	r.GET("/SyncStatus", routes.SyncStatus)
	r.POST("/BlockByNumber", routes.BlockbyNumber)
	r.GET("/TotTxCount", routes.TotalTransactionCount)

	//r.GET("/blockbyhash", routes.BlockbyHash)

	r.Run()

}