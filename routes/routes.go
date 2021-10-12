package routes

import (
	"context"
	"encoding/json"
	"fiscoSharespot/utils"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var conn *client.Client

func Init() {
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatalf("ParseConfigFile failed, err: %v", err)
	}
	conn, err = client.Dial(&configs[0])
	if err != nil {
		log.Fatal(err)
	}
}
func ResponseData(c *gin.Context, resp *utils.Resp) {
	resp.Msg = utils.RecodeText(resp.Recode)
	c.JSON(http.StatusOK, resp)
}

// BlockNumber GET: /blockNumber
func BlockNumber(c *gin.Context) {
	// new Resp instance
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	blockNumber, err := conn.GetBlockNumber(context.Background())
	if err != nil {
		fmt.Println("GetBlockNumber failed")
		resp.Recode = utils.RecodeFiscoErr
		//return err
	}
	// arrange response data
	mapResp := make(map[string]interface{})
	mapResp["block_number"] = blockNumber
	resp.Data = mapResp
	//return nil
}

// NodePeers GET /nodepeers
func NodePeers(c *gin.Context) {
	// new Resp instance
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	nodepeers, err := conn.GetPeers(context.Background())

	if err != nil {
		fmt.Println("GetNodePeers failed")
		resp.Recode = utils.RecodeFiscoErr
	}

	var jpeers interface{}
	json.Unmarshal(nodepeers, &jpeers)

	// arrange response data
	mapResp := make(map[string]interface{})
	mapResp["node_peers"] = jpeers
	resp.Data = mapResp
}

// SyncStatus GET /syncstatus
func SyncStatus(c *gin.Context) {
	// new Resp instance
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	syncstatus, err := conn.GetSyncStatus(context.Background())

	if err != nil {
		fmt.Println("GetSyncStatus failed")
		resp.Recode = utils.RecodeFiscoErr
	}

	var jsyncstatus interface{}
	json.Unmarshal(syncstatus, &jsyncstatus)

	// arrange response data
	mapResp := make(map[string]interface{})
	mapResp["sync_status"] = jsyncstatus
	resp.Data = mapResp
}

// BlockbyNumber POST /blockbynumber
// curl http://localhost:8080/blockbynumber -X POST -d "blocknumber=12"
func BlockbyNumber(c *gin.Context) {
	// new Resp instance
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	// 查询的区块号，默认0
	postnumber := c.DefaultPostForm("BlockNumber","0")
	postnumberint64, err := strconv.ParseInt(postnumber, 10, 64)
	fmt.Println(postnumberint64)
	if err != nil {
		fmt.Println("Data Transfer failed")
		resp.Recode = utils.RecodeUnknownErr
	}
	blockinfo, err := conn.GetBlockByNumber(context.Background(),postnumberint64,false)

	if err != nil {
		fmt.Println("GetBlockbyNumber failed")
		resp.Recode = utils.RecodeFiscoErr
	}

	var jblockinfo interface{}
	json.Unmarshal(blockinfo, &jblockinfo)

	// arrange response data
	mapResp := make(map[string]interface{})
	mapResp["block_infomation"] = jblockinfo
	resp.Data = mapResp
}

// TotalTransactionCount GET /tottransactioncount
type Trancnt struct {
	BlockNumber    string  `json:"blocknumber"`
	FailedTxSum    string  `json:"failedTxSum"`
	TxSum    string `json:"txSum"`
}
func TotalTransactionCount(c *gin.Context) {
	// new Resp instance
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	tottscnt, err := conn.GetTotalTransactionCount(context.Background())

	if err != nil {
		fmt.Println("GetTotalTransactionCount failed")
		resp.Recode = utils.RecodeFiscoErr
	}

	var jtottscnt Trancnt
	json.Unmarshal(tottscnt, &jtottscnt)
	fmt.Println(jtottscnt)
	n1 , err := strconv.ParseUint(jtottscnt.BlockNumber[2:],16,32)
	if err != nil{
		fmt.Println("cao")
		resp.Recode = utils.RecodeUnknownErr
	}
	jtottscnt.BlockNumber = strconv.Itoa(int(n1))
	fmt.Println(strconv.Itoa(int(n1)))

	n2 , err := strconv.ParseUint(jtottscnt.TxSum[2:],16,32)
	if err != nil{
		fmt.Println("cao")
		resp.Recode = utils.RecodeUnknownErr
	}
	jtottscnt.TxSum = strconv.Itoa(int(n2))
	fmt.Println(strconv.Itoa(int(n2)))

	n3 , err := strconv.ParseUint(jtottscnt.FailedTxSum[2:],16,32)
	if err != nil{
		fmt.Println("cao")
		resp.Recode = utils.RecodeUnknownErr
	}
	jtottscnt.FailedTxSum = strconv.Itoa(int(n3))
	fmt.Println(strconv.Itoa(int(n3)))
	// arrange response data
	mapResp := make(map[string]interface{})
	mapResp["tot_trans_cnt"] = jtottscnt
	resp.Data = mapResp
}
//
//// BlockbyHash POST /blockbyhash
//func BlockbyHash(c *gin.Context) {
//	// new Resp instance
//	resp := utils.Resp{
//		Recode: utils.RecodeOk,
//	}
//	defer ResponseData(c, &resp)
//
//	blockbyhash, err := conn.GetBlockByHash(context.Background())
//
//	if err != nil {
//		fmt.Println("GetBlockNumber failed")
//		resp.Recode = utils.RecodeFiscoErr
//	}
//
//	var jblockbyhash interface{}
//	json.Unmarshal(blockbyhash, &jblockbyhash)
//
//	// arrange response data
//	mapResp := make(map[string]interface{})
//	mapResp["block_information"] = jblockbyhash
//	resp.Data = mapResp
//}

