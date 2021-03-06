package routers

import (
	v1 "fiscoSharespot/routers/api/v1"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// InitRouter 初始化路由信息
func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(sessions.Sessions("session", v1.Store))
	r.Use(Cors())
	apiV1 := r.Group("/api/v1")
	{
		// blockchain api
		apiV1.GET("/ping", v1.PingHandler)
		apiV1.GET("/blockNumber", v1.GetBlockNumber)
		apiV1.GET("/nodePeers", v1.GetNodePeers)
		apiV1.GET("/syncStatus", v1.GetSyncStatus)
		apiV1.POST("/blockByNumber", v1.BlockByNumber)
		apiV1.GET("/totalTransactionCount", v1.GetTotalTransactionCount)

		// user api
		apiV1.GET("/userContractAddress",v1.GetUserManagementContractAddress)
		apiV1.POST("/user", v1.AddUser)
		apiV1.POST("login", v1.Login)
		apiV1.GET("/allUsers", v1.GetAllUsers)
		apiV1.POST("/transfer", v1.Transfer)
		apiV1.POST("/balanceByID", v1.BalanceByID)
		apiV1.POST("/balanceByPubKey", v1.BalanceByPubKey)
		apiV1.POST("/userByID",v1.UserByID)

		// resource api
		apiV1.GET("/resourceContractAddress",v1.GetResourceManagementContractAddress)
		apiV1.POST("/resource",v1.AddResource)
		apiV1.POST("/resourcePreBuy",v1.PreBuyResource)
		apiV1.POST("/resourceGrant",v1.GrantResource)
		apiV1.POST("/resourceFinish",v1.FinishResource)
		apiV1.POST("/resourceByResourceID",v1.ResourceByResourceID)
		apiV1.GET("/allResources",v1.GetAllResources)
		apiV1.POST("/resourceByStatus",v1.ResourceByStatus)
	}
	// 静态文件路由
	//r.StaticFS("/web", http.Dir("./dist/"))
	return r
}

// Cors 允许跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			// header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()
	}
}