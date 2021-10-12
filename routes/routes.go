package routes

import (
	"context"
	"fiscoSharespot/utils"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
		log.Fatal(err)
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

