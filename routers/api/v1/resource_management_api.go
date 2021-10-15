package v1

import (
	"fiscoSharespot/resource_management"
	"fiscoSharespot/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
)

//var Resaddress common.Address
//var RescontractTX *types.Transaction
//var Resinstance *resource_pool.ResourcePool
//var respoolSession *resource_pool.ResourcePoolSession

// DeployResourcePool
//curl http://localhost:8080/DeployRes
//func DeployResourcePool(c *gin.Context){
//	resp := utils.Resp{
//		Recode: utils.RecodeOk,
//	}
//	defer ResponseData(c, &resp)
//	//部署合约
//	resourceAddress, resourceContractTX, resourceInstance, err := resource_management.DeployResourceManagement(conn.GetTransactOpts(), conn)
//	if err != nil {
//		fmt.Println("DeployResPool failed")
//		resp.Recode = utils.RecodeFiscoErr
//	}
//
//	resourcePoolSession := &resource_management.ResourceManagementSession{
//		Contract:     resourceInstance,
//		CallOpts:     *conn.GetCallOpts(),
//		TransactOpts: *conn.GetTransactOpts(),
//	}
//
//	resources , err := resourcePoolSession.GetAllResource()
//	fmt.Println("Resources len:" , len(resources))
//	if err != nil{
//		fmt.Println("Get Resources Failed!")
//		resp.Recode = utils.RecodeFiscoErr
//	}
//
//	mapResp := make(map[string]interface{})
//	mapResp["contract_address"] = resourceAddress.Hex()
//	mapResp["tx_hash"] = resourceContractTX.Hash().Hex()
//	mapResp["now_length"] = len(resources)
//	resp.Data = mapResp
//}

// AddResource POST /resource
func AddResource(c *gin.Context)  {

	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)
	//部署合约

	resourcePoolSession := &resource_management.ResourceManagementSession{
		Contract:     resourceManagementInstance,
		CallOpts:     *conn.GetCallOpts(),
		TransactOpts: *conn.GetTransactOpts(),
	}

	id := c.PostForm("ownerID")
	serviceType := c.PostForm("serviceType")
	serviceTime := c.PostForm("serviceTime")
	price := c.PostForm("price")

	_serviceTime , _ := new(big.Int).SetString(serviceTime,10)

	_price , _ := new(big.Int).SetString(price, 10)

	tx, receipt, err := resourcePoolSession.AddResource(id, serviceType, _serviceTime, _price)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())

	resources , err := resourcePoolSession.GetAllResource()
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

// PreBuyResource POST /resourcePreBuy
func PreBuyResource(c *gin.Context)  {
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	resourceID := c.PostForm("resourceID")
	id := c.PostForm("buyerID")

	resourceManagementSession := &resource_management.ResourceManagementSession{
		Contract:     resourceManagementInstance,
		CallOpts:     *conn.GetCallOpts(),
		TransactOpts: *conn.GetTransactOpts(),
	}

	tx, receipt, err := resourceManagementSession.PreBuyResource(resourceID, id)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())

	resources , err := resourceManagementSession.GetResource(resourceID)

	if err != nil{
		fmt.Println("Get Resources Failed!")
		resp.Recode = utils.RecodeFiscoErr
	}

	fmt.Printf("%+v", resources)

	resp.Data = resources
}

// GrantResource POST /resourceGrant
func GrantResource(c *gin.Context)  {
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	resourceID := c.PostForm("resourceID")
	accessCode := c.PostForm("accessCode")

	resourceManagementSession := &resource_management.ResourceManagementSession{
		Contract:     resourceManagementInstance,
		CallOpts:     *conn.GetCallOpts(),
		TransactOpts: *conn.GetTransactOpts(),
	}

	tx, receipt, err := resourceManagementSession.GrantResource(resourceID, accessCode)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())

	resources , err := resourceManagementSession.GetResource(resourceID)

	if err != nil{
		fmt.Println("Get Resources Failed!")
		resp.Recode = utils.RecodeFiscoErr
	}

	fmt.Printf("%+v", resources)

	resp.Data = resources
}

// FinishResource POST /resourceFinish
func FinishResource(c *gin.Context)  {
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	resourceID := c.PostForm("resourceID")

	resourceManagementSession := &resource_management.ResourceManagementSession{
		Contract:     resourceManagementInstance,
		CallOpts:     *conn.GetCallOpts(),
		TransactOpts: *conn.GetTransactOpts(),
	}

	tx, receipt, err := resourceManagementSession.FinishResource(resourceID)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())

	resources , err := resourceManagementSession.GetResource(resourceID)

	if err != nil{
		fmt.Println("Get Resources Failed!")
		resp.Recode = utils.RecodeFiscoErr
	}

	fmt.Printf("%+v", resources)

	resp.Data = resources
}

// ResourceByResourceID POST /
func ResourceByResourceID(c *gin.Context)  {
	//初次调试，链上未部署合约，所以此方法先部署了合约，并获取了instance，后期需要把部署合约的功能写在init函数中
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	resourceID := c.PostForm("resourceID")

	resourceManagementSession := &resource_management.ResourceManagementSession{
		Contract:     resourceManagementInstance,
		CallOpts:     *conn.GetCallOpts(),
		TransactOpts: *conn.GetTransactOpts(),
	}

	resources , err := resourceManagementSession.GetResource(resourceID)

	if err != nil{
		fmt.Println("Get Resources Failed!")
		resp.Recode = utils.RecodeFiscoErr
	}

	fmt.Printf("%+v", resources)

	resp.Data = resources
}

// GetAllResources GET /allResources
func GetAllResources(c *gin.Context)  {
	//初次调试，链上未部署合约，所以此方法先部署了合约，并获取了instance，后期需要把部署合约的功能写在init函数中
	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	resourceManagementSession := &resource_management.ResourceManagementSession{
		Contract:     resourceManagementInstance,
		CallOpts:     *conn.GetCallOpts(),
		TransactOpts: *conn.GetTransactOpts(),
	}

	resources , err := resourceManagementSession.GetAllResource()
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

func ResourceByStatus(c *gin.Context)  {

	resp := utils.Resp{
		Recode: utils.RecodeOk,
	}
	defer ResponseData(c, &resp)

	resourceStatus := c.PostForm("resourceStatus")

	resourceManagementSession := &resource_management.ResourceManagementSession{
		Contract:     resourceManagementInstance,
		CallOpts:     *conn.GetCallOpts(),
		TransactOpts: *conn.GetTransactOpts(),
	}

	mark := true
	if resourceStatus != "ForSale" && resourceStatus != "PreBuy" && resourceStatus != "Serving" && resourceStatus !="ServiceFinished" {
		mark = false
	}
	if (!mark){
		mapResp := make(map[string]interface{})
		mapResp["error"] = "Status Not Exist"
		resp.Data = mapResp
	} else{
		resources , err := resourceManagementSession.GetAllResource()

		if err != nil{
			fmt.Println("Get Resources Failed!")
			resp.Recode = utils.RecodeFiscoErr
		}

		var resourcesStatusList[] resource_management.Struct0
		for i := 0; i< len(resources); i++ {
			if resources[i].Status == resourceStatus{
				fmt.Println(resources[i].Status)
				resourcesStatusList = append(resourcesStatusList, resources[i])
			}
		}

		fmt.Printf("Status Get%+v", resourcesStatusList)
		resp.Data = resourcesStatusList
	}
}
