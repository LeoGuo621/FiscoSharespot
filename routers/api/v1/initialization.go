package v1

import (
	"fiscoSharespot/contracts"
	"fiscoSharespot/utils"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var conn *client.Client
// Contract Info
var userManagementAddress common.Address
var userManagementContractTX *types.Transaction
var userManagementInstance *contracts.UserManagement

func init() {
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatalf("ParseConfigFile failed, err: %v", err)
	}
	conn, err = client.Dial(&configs[0])
	if err != nil {
		log.Fatalf("Dial failed, err: %v", err)
	}
	userManagementAddress, userManagementContractTX, userManagementInstance, err = contracts.DeployUserManagement(conn.GetTransactOpts(), conn)
	if err != nil {
		log.Fatalf("DeployUserManagement failed, err: %v", err)
	}
}

// ResponseData response
func ResponseData(c *gin.Context, resp *utils.Resp) {
	resp.Msg = utils.RecodeText(resp.Recode)
	c.JSON(http.StatusOK, resp)
}

// GetUserManagementContractAddress GET: /userManagementContractAddress
func GetUserManagementContractAddress(c *gin.Context) {
	var resp utils.Resp
	resp.Recode = utils.RecodeOk
	defer ResponseData(c, &resp)

	mapResp := make(map[string]interface{})
	mapResp["user_management_contract_address"] = userManagementAddress
	resp.Data = mapResp
}
