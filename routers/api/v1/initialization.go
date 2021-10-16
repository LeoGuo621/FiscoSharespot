package v1

import (
	"fiscoSharespot/resource_management"
	"fiscoSharespot/user_management"
	"fiscoSharespot/utils"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var conn *client.Client
var Store cookie.Store
// Contract Info
var userManagementAddress common.Address
var userManagementContractTX *types.Transaction
var userManagementInstance *user_management.UserManagement

var resourceManagementAddress common.Address
var resourceManagementContractTX *types.Transaction
var resourceManagementInstance *resource_management.ResourceManagement

func init() {
	Store = cookie.NewStore([]byte("Fintechathon2021"))
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatalf("ParseConfigFile failed, err: %v", err)
	}
	conn, err = client.Dial(&configs[0])
	if err != nil {
		log.Fatalf("Dial failed, err: %v", err)
	}
	userManagementAddress, userManagementContractTX, userManagementInstance, err = user_management.DeployUserManagement(conn.GetTransactOpts(), conn)
	if err != nil {
		log.Fatalf("DeployUserManagement failed, err: %v", err)
	}

	resourceManagementAddress, resourceManagementContractTX, resourceManagementInstance, err = resource_management.DeployResourceManagement(conn.GetTransactOpts(), conn)
	if err != nil {
		log.Fatalf("DeployResourcePool failed, err: %v", err)
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

// GetResourceManagementContractAddress GET: /resourceManagementAddress
func GetResourceManagementContractAddress(c *gin.Context) {
	var resp utils.Resp
	resp.Recode = utils.RecodeOk
	defer ResponseData(c, &resp)

	mapResp := make(map[string]interface{})
	mapResp["resource_pool_contract_address"] = resourceManagementAddress
	resp.Data = mapResp
}