package v1

import (
	"fiscoSharespot/user_management"
	"fiscoSharespot/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"math/big"
)

// Login POST /login
func Login(c *gin.Context)  {
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}

	defer ResponseData(c, &resp)
	//部署合约
	id := c.PostForm("userID")
	passwordEncoded := c.PostForm("passwordEncoded")

	userManagementSession := &user_management.UserManagementSession{
		Contract:     userManagementInstance,
		CallOpts:     *conn.GetCallOpts(),
		TransactOpts: *conn.GetTransactOpts(),
	}
	user, err := userManagementSession.GetUserByID(id)
	if err != nil {
		fmt.Println("no such user!")
		resp.Recode = utils.RecodeLoginErr
	}
	if user.PasswordEncoded != passwordEncoded {
		fmt.Println("password error!")
		resp.Recode = utils.RecodeLoginErr
	}
	// session
	sess := sessions.Default(c)
	option := sessions.Options{Path: "/", MaxAge: 3600, HttpOnly: true}
	sess.Options(option)

	sess.Set(user.Id, user)
	err = sess.Save()
	if err != nil {
		fmt.Println("sess.Save failed")
		resp.Recode = utils.RecodeUnknownErr
	}
	resp.Data = sess.Get(user.Id)
}


// AddUser POST /user
func AddUser(c *gin.Context)  {
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)
	//部署合约
	id := c.PostForm("userID")
	pubKey := c.PostForm("userPubKey")
	name := c.PostForm("name")
	passwordEncoded := c.PostForm("passwordEncoded")

	userManagementSession := &user_management.UserManagementSession{
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

	userManagementSession := &user_management.UserManagementSession{
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

// BalanceByID POST /balanceByID
func BalanceByID(c *gin.Context)  {
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	userID := c.PostForm("userID")

	userManagementSession := &user_management.UserManagementSession{
		Contract:     userManagementInstance,
		CallOpts:     *conn.GetCallOpts(),
		TransactOpts: *conn.GetTransactOpts(),
	}

	balance, err := userManagementSession.GetBalanceByID(userID)

	if err != nil{
		fmt.Println("GetBalanceByID Failed!")
		resp.Recode = utils.RecodeContractErr
	}
	mapResp := make(map[string]interface{})
	mapResp["balance"] = balance
	resp.Data = mapResp
}

// BalanceByPubKey POST /balanceByPubKey
func BalanceByPubKey(c *gin.Context)  {
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	userPubKey := c.PostForm("userPubKey")

	userManagementSession := &user_management.UserManagementSession{
		Contract:     userManagementInstance,
		CallOpts:     *conn.GetCallOpts(),
		TransactOpts: *conn.GetTransactOpts(),
	}

	balance, err := userManagementSession.GetBalanceByPubKey(userPubKey)

	if err != nil{
		fmt.Println("GetBalanceByPubKey Failed!")
		resp.Recode = utils.RecodeContractErr
	}
	mapResp := make(map[string]interface{})
	mapResp["balance"] = balance
	resp.Data = mapResp
}

// UserByID POST /userByID
func UserByID(c *gin.Context)  {
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	userID := c.PostForm("userID")

	userManagementSession := &user_management.UserManagementSession{
		Contract:     userManagementInstance,
		CallOpts:     *conn.GetCallOpts(),
		TransactOpts: *conn.GetTransactOpts(),
	}

	user, err := userManagementSession.GetUserByID(userID)

	if err != nil{
		fmt.Println("GetUserByID Failed!")
		resp.Recode = utils.RecodeContractErr
	}
	resp.Data = user
}

// Transfer POST /transfer
func Transfer(c *gin.Context) {
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)
	from := c.PostForm("from")
	to := c.PostForm("to")
	amountStr := c.PostForm("amount")
	userManagementSession := &user_management.UserManagementSession{
		Contract:     userManagementInstance,
		CallOpts:     *conn.GetCallOpts(),
		TransactOpts: *conn.GetTransactOpts(),
	}
	amount , _ := new(big.Int).SetString(amountStr, 10)
	tx, receipt, err := userManagementSession.Transfer(from, to, amount)
	if err != nil {
		fmt.Println("Transfer failed")
		resp.Recode = utils.RecodeContractErr
	}

	mapResp := make(map[string]interface{})
	mapResp["tx_hash"] = tx.Hash().Hex()
	mapResp["block_number"] = receipt.GetBlockNumber()
	mapResp["block_hash"] = receipt.GetBlockHash()
	resp.Data = mapResp
}