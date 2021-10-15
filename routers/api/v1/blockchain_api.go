package v1

import (
	"context"
	"encoding/json"
	"fiscoSharespot/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetBlockNumber GET /blockNumber
func GetBlockNumber(c *gin.Context) {
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

// GetNodePeers GET /nodePeers
func GetNodePeers(c *gin.Context) {
	// new Resp instance
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	nodePeers, err := conn.GetPeers(context.Background())

	if err != nil {
		fmt.Println("GetNodePeers failed")
		resp.Recode = utils.RecodeFiscoErr
	}

	var jsonPeers interface{}
	_ = json.Unmarshal(nodePeers, &jsonPeers)

	// arrange response data
	mapResp := make(map[string]interface{})
	mapResp["node_peers"] = jsonPeers
	resp.Data = mapResp
}

// GetSyncStatus GET /syncStatus
func GetSyncStatus(c *gin.Context) {
	// new Resp instance
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	syncStatus, err := conn.GetSyncStatus(context.Background())

	if err != nil {
		fmt.Println("GetSyncStatus failed")
		resp.Recode = utils.RecodeFiscoErr
	}

	var jsonSyncStatus interface{}
	_ = json.Unmarshal(syncStatus, &jsonSyncStatus)

	// arrange response data
	mapResp := make(map[string]interface{})
	mapResp["sync_status"] = jsonSyncStatus
	resp.Data = mapResp
}

// BlockByNumber POST /blockByNumber
// curl http://localhost:8080/blockByNumber -X POST -d "blockNumber=12"
func BlockByNumber(c *gin.Context) {
	// new Resp instance
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	postNumber := c.PostForm("blockNumber")
	postNumberInt64, err := strconv.ParseInt(postNumber, 10, 64)
	//fmt.Println(postNumberInt64)
	if err != nil {
		fmt.Println("ParseInt failed")
		resp.Recode = utils.RecodeUnknownErr
	}
	blockInfo, err := conn.GetBlockByNumber(context.Background(), postNumberInt64,false)
	if err != nil {
		fmt.Println("GetBlockByNumber failed")
		resp.Recode = utils.RecodeFiscoErr
	}

	var jsonBlockInfo interface{}
	_ = json.Unmarshal(blockInfo, &jsonBlockInfo)

	// arrange response data
	mapResp := make(map[string]interface{})
	mapResp["block_information"] = jsonBlockInfo
	resp.Data = mapResp
}

type TXCnt struct {
	BlockNumber string `json:"block_number"`
	FailedTXSum string `json:"failed_tx_sum"`
	TXSum       string `json:"tx_sum"`
}

// GetTotalTransactionCount GET /totalTransactionCount
func GetTotalTransactionCount(c *gin.Context) {
	// new Resp instance
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	totalTXCnt, err := conn.GetTotalTransactionCount(context.Background())

	if err != nil {
		fmt.Println("GetTotalTransactionCount failed")
		resp.Recode = utils.RecodeFiscoErr
	}

	var jsonTotalTXCnt TXCnt
	_ = json.Unmarshal(totalTXCnt, &jsonTotalTXCnt)
	//fmt.Println(jsonTotalTXCnt)
	blockNumberUint64, err := strconv.ParseUint(jsonTotalTXCnt.BlockNumber[2:],16,32)
	if err != nil{
		fmt.Println("ParseUint failed")
		resp.Recode = utils.RecodeUnknownErr
	}
	jsonTotalTXCnt.BlockNumber = strconv.Itoa(int(blockNumberUint64))
	//fmt.Println(strconv.Itoa(int(blockNumberUint64)))

	TXSumUint64, err := strconv.ParseUint(jsonTotalTXCnt.TXSum[2:],16,32)
	if err != nil{
		fmt.Println("ParseUint failed")
		resp.Recode = utils.RecodeUnknownErr
	}
	jsonTotalTXCnt.TXSum = strconv.Itoa(int(TXSumUint64))
	//fmt.Println(strconv.Itoa(int(TXSumUint64)))

	failedTXSumUint64, err := strconv.ParseUint(jsonTotalTXCnt.FailedTXSum[2:],16,32)
	if err != nil{
		fmt.Println("ParseUint failed")
		resp.Recode = utils.RecodeUnknownErr
	}
	jsonTotalTXCnt.FailedTXSum = strconv.Itoa(int(failedTXSumUint64))
	//fmt.Println(strconv.Itoa(int(failedTXSumUint64)))

	// arrange response data
	mapResp := make(map[string]interface{})
	mapResp["total_tx_cnt"] = jsonTotalTXCnt
	resp.Data = mapResp
}
