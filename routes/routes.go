package routes

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fiscoSharespot/user_management"
	"fiscoSharespot/utils"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)


type HandlerFunc func(c *gin.Context) error

var conn *client.Client

func init() {
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatalf("ParseConfigFile failed, err: %v", err)
	}
	conn, err = client.Dial(&configs[0])
	if err != nil {
		log.Fatalf("Dial failed, err: %v", err)
	}
}

// ResponseData response
func ResponseData(c *gin.Context, resp *utils.Resp) {
	resp.Msg = utils.RecodeText(resp.Recode)
	c.JSON(http.StatusOK, resp)
}

// PingHandler GET: /ping
func PingHandler(c *gin.Context) {
	var resp utils.Resp
	resp.Recode = utils.RecodeOk
	defer ResponseData(c, &resp)
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

// BlockByNumber POST /blockByNumber
// curl http://localhost:8080/blockbynumber -X POST -d "blockNumber=12"
func BlockByNumber(c *gin.Context) {
	// new Resp instance
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	// 查询的区块号，默认0
	//postNumber = c.DefaultPostForm("BlockNumber","0")
	postNumber := c.Query("blockNumber")
	//if err := c.Bind(&postNumber); err != nil {
	//	fmt.Println("Failed to bind postNumber: ", err)
	//	resp.Recode = utils.RecodeParamErr
	//}
	//fmt.Println("postNumber: ", postNumber, "#")
	postNumberInt64, err := strconv.ParseInt(postNumber, 10, 64)
	//fmt.Println(postNumberInt64)
	if err != nil {
		fmt.Println("Data Transfer failed")
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

func AddUser(c *gin.Context)  {
	//初次调试，链上未部署合约，所以此方法先部署了合约，并获取了instance，后期需要把部署合约的功能写在init函数中
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)
	//部署合约
	id := c.PostForm("id")
	addr := common.HexToAddress(c.PostForm("addr"))
	name := c.PostForm("name")
	// 对前端传来的字符串hash再次hash
	passwordEncoded := sha256.Sum256([]byte(c.PostForm("passwordEncoded")))
	address, contractTX, instance, err := user_management.DeployUserManagement(conn.GetTransactOpts(), conn)
	if err != nil {
		fmt.Println("DeployUserManagement failed")
		resp.Recode = utils.RecodeFiscoErr
	}

	userManagementSession := &user_management.UserManagementSession{
		Contract:     instance,
		CallOpts:     *conn.GetCallOpts(),
		TransactOpts: *conn.GetTransactOpts(),
	}

	tx, receipt, err := userManagementSession.AddUser(id, addr, name, passwordEncoded)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())

	users, err := userManagementSession.GetAllUsers()
	fmt.Println("users len: ", len(users))
	if err != nil {
		fmt.Println("GetAllUsers failed")
		resp.Recode = utils.RecodeFiscoErr
	}
	for i := 0; i < len(users); i++ {
		fmt.Printf("%s\n", users[i].Id)
		fmt.Printf("%s\n", users[i].Addr.Hex())
		fmt.Printf("%s\n", users[i].Name)
		fmt.Println(hex.EncodeToString(users[i].PasswordEncoded[:]))
	}

	mapResp := make(map[string]interface{})
	mapResp["contract_address"] = address.Hex()
	mapResp["tx_hash"] = contractTX.Hash().Hex()
	mapResp["user0_addr"] = users[0].Addr.Hex()
	resp.Data = mapResp
}