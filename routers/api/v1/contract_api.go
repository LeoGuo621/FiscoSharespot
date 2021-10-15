package v1

import (
	"fiscoSharespot/contracts"
	"fiscoSharespot/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

// AddUser POST /user
func AddUser(c *gin.Context)  {
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)
	//部署合约
	id := c.PostForm("id")
	pubKey := c.PostForm("pubKey")
	name := c.PostForm("name")
	passwordEncoded := c.PostForm("passwordEncoded")

	userManagementSession := &contracts.UserManagementSession{
		Contract:     userManagementInstance,
		CallOpts:     *conn.GetCallOpts(),
		TransactOpts: *conn.GetTransactOpts(),
	}
	tx, receipt, err := userManagementSession.AddUser(id, pubKey, name, passwordEncoded)
	if err != nil {
		fmt.Println("AddUser failed")
		resp.Recode = utils.RecodeContractErr
	}

	mapResp := make(map[string]interface{})
	mapResp["tx_hash"] = tx.Hash().Hex()
	mapResp["block_number"] = receipt.GetBlockNumber()
	mapResp["block_hash"] = receipt.GetBlockHash()
	resp.Data = mapResp
}

// GetAllUsers GET /allUsers
func GetAllUsers(c *gin.Context)  {
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	userManagementSession := &contracts.UserManagementSession{
		Contract:     userManagementInstance,
		CallOpts:     *conn.GetCallOpts(),
		TransactOpts: *conn.GetTransactOpts(),
	}
	users, err := userManagementSession.GetAllUsers()
	if err != nil {
		fmt.Println("GetAllUsers failed")
		resp.Recode = utils.RecodeContractErr
	}
	resp.Data = users
}