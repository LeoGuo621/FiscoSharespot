package v1

import (
	"fiscoSharespot/utils"
	"github.com/gin-gonic/gin"
)

// PingHandler GET: /ping
func PingHandler(c *gin.Context) {
	var resp utils.Resp
	resp.Recode = utils.RecodeOk
	defer ResponseData(c, &resp)
}
