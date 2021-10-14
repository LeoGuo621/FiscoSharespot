package routes

import (
	"context"
	"encoding/json"
<<<<<<< HEAD
=======
	"fiscoSharespot/respool"
>>>>>>> 49ffd75 (ResPool Contract Finished)
	"fiscoSharespot/utils"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
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
<<<<<<< HEAD
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
=======

var Resaddress common.Address
var RescontractTX *types.Transaction
var Resinstance *res_pool.ResPool
var respoolSession *res_pool.ResPoolSession

//curl http://localhost:8080/DeployRes
func DeployResContract(c *gin.Context){
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)
	//部署合约
	Resaddress, RescontractTX, Resinstance, err := res_pool.DeployResPool(conn.GetTransactOpts(), conn)
	if err != nil {
		fmt.Println("DeployResPool failed")
		resp.Recode = utils.RecodeFiscoErr
	}

	respoolSession = &res_pool.ResPoolSession{
		Contract:     Resinstance,
		CallOpts:     *conn.GetCallOpts(),
		TransactOpts: *conn.GetTransactOpts(),
	}

	resources , err := respoolSession.GetAllRes()
	fmt.Println("Resources len:" , len(resources))
	if err != nil{
		fmt.Println("Get Resources Failed!")
		resp.Recode = utils.RecodeFiscoErr
	}

	mapResp := make(map[string]interface{})
	mapResp["contract_address"] = Resaddress.Hex()
	mapResp["tx_hash"] = RescontractTX.Hash().Hex()
	mapResp["now_length"] = len(resources)
	resp.Data = mapResp
}
//curl http://localhost:8080/AddRes -X POST -d "OwnerID=jyh" -d "ServiceType=Wifi" -d "ServiceTime=100" -d "Price=10"
func AddRes(c *gin.Context)  {
	//初次调试，链上未部署合约，所以此方法先部署了合约，并获取了instance，后期需要把部署合约的功能写在init函数中
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)
	//部署合约
	id := c.PostForm("OwnerID")
	sertype := c.PostForm("ServiceType")
	sertime := c.PostForm("ServiceTime")
	serprice := c.PostForm("Price")

	_sertime , _ := new(big.Int).SetString(sertime,10)

	_price , _ := new(big.Int).SetString(serprice, 10)

	tx, receipt, err := respoolSession.AddRes(id, sertype, _sertime, _price)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())

	resources , err := respoolSession.GetAllRes()
	fmt.Println("Resources len:" , len(resources))
	if err != nil{
		fmt.Println("Get Resources Failed!")
		resp.Recode = utils.RecodeFiscoErr
	}
	for i := 0; i < len(resources); i++ {
		fmt.Printf("%+v", resources)
	}
	resp.Data = resources[len(resources)-1]
}

func PreBuyRes(c *gin.Context)  {
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	resid := c.PostForm("ResID")
	id := c.PostForm("BuyerID")

	tx, receipt, err := respoolSession.PreBuyRes(resid, id)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())

	resources , err := respoolSession.GetRes(resid)

	if err != nil{
		fmt.Println("Get Resources Failed!")
		resp.Recode = utils.RecodeFiscoErr
	}

	fmt.Printf("%+v", resources)

	resp.Data = resources
}

func ApRes(c *gin.Context)  {
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	resid := c.PostForm("ResID")
	acode := c.PostForm("AccessCode")

	tx, receipt, err := respoolSession.ApRes(resid, acode)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())

	resources , err := respoolSession.GetRes(resid)

	if err != nil{
		fmt.Println("Get Resources Failed!")
		resp.Recode = utils.RecodeFiscoErr
	}

	fmt.Printf("%+v", resources)

	resp.Data = resources
}

func FinalRes(c *gin.Context)  {
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	resid := c.PostForm("ResID")

	tx, receipt, err := respoolSession.FinalRes(resid)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())

	resources , err := respoolSession.GetRes(resid)

	if err != nil{
		fmt.Println("Get Resources Failed!")
		resp.Recode = utils.RecodeFiscoErr
	}

	fmt.Printf("%+v", resources)

	resp.Data = resources
}
//curl http://localhost:8080/GetRes -X POST -d "ResID=xxxxx"
func GetRes(c *gin.Context)  {
	//初次调试，链上未部署合约，所以此方法先部署了合约，并获取了instance，后期需要把部署合约的功能写在init函数中
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	resid := c.PostForm("ResID")

	resources , err := respoolSession.GetRes(resid)

	if err != nil{
		fmt.Println("Get Resources Failed!")
		resp.Recode = utils.RecodeFiscoErr
	}

	fmt.Printf("%+v", resources)

	resp.Data = resources
}


func GetResAll(c *gin.Context)  {
	//初次调试，链上未部署合约，所以此方法先部署了合约，并获取了instance，后期需要把部署合约的功能写在init函数中
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	resources , err := respoolSession.GetAllRes()
	fmt.Println("Resources len:" , len(resources))
	if err != nil{
		fmt.Println("Get Resources Failed!")
		resp.Recode = utils.RecodeFiscoErr
	}

	for i := 0; i < len(resources); i++ {
		fmt.Printf("%+v", resources)
	}

	resp.Data = resources
}

>>>>>>> 49ffd75 (ResPool Contract Finished)

