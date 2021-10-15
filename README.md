# 后端响应框架

[toc]

## 接口列表

|               API               | 请求方式 |          功能          |                       参数                        |
| :-----------------------------: | :------: | :--------------------: | :-----------------------------------------------: |
|          /api/v1/ping           |   GET    |          PING          |                         /                         |
|       /api/v1/blockNumber       |   GET    |    获取最新区块高度    |                         /                         |
|        /api/v1/nodePeers        |   GET    |        查询节点        |                         /                         |
|       /api/v1/syncStatus        |   GET    |      查询同步状态      |                         /                         |
|      /api/v1/blockByNumber      |   POST   |  通过区块高度获取区块  |                   "blockNumber"                   |
|  /api/v1/totalTransactionCount  |   GET    |      获取总交易数      |                         /                         |
|   /api/v1/userContractAddress   |   GET    |  获取用户管理合约地址  |                         /                         |
|          /api/v1/user           |   POST   |        新增用户        | "userID", "userPubKey", "name", "passwordEncoded" |
|        /api/v1/allUsers         |   GET    |      获取全部用户      |                         /                         |
|        /api/v1/transfer         |   POST   |          转账          |              "from", "to", "amount"               |
|       /api/v1/balanceByID       |   POST   |   通过用户ID获取余额   |                     "userID"                      |
|     /api/v1/balanceByPubKey     |   POST   |  通过用户公钥获取余额  |                   "userPubKey"                    |
|        /api/v1/userByID         |   POST   | 通过用户ID获取用户对象 |                     "userID"                      |
| /api/v1/resourceContractAddress |   GET    |  获取资源管理合约地址  |                         /                         |
|        /api/v1/resource         |   POST   |        新增用户        | "ownerID", "serviceType", "serviceTime", "price"  |
|     /api/v1/resourcePreBuy      |   POST   |        资源预购        |              "resourceID", "buyerID"              |
|      /api/v1/resourceGrant      |   POST   |        资源授权        |            "resourceID", "accessCode"             |
|     /api/v1/resourceFinish      |   POST   |      资源服务终止      |                   "resourceID"                    |
|  /api/v1/resourceByResourceID   |   POST   |   通过资源ID获取资源   |                   "resourceID"                    |
|      /api/v1/allResources       |   GET    |      获取所有资源      |                         /                         |
|    /api/v1/resourceByStatus     |   POST   |  通过资源状态获取资源  |                 "resourceStatus"                  |



## BlockNumber

![image-20211012220400860](https://luochengyu.oss-cn-beijing.aliyuncs.com/img/image-20211012220400860.png)

## NodePeers

![image-20211012220432518](https://luochengyu.oss-cn-beijing.aliyuncs.com/img/image-20211012220432518.png)

## SyncStatus

![image-20211012220456977](https://luochengyu.oss-cn-beijing.aliyuncs.com/img/image-20211012220456977.png)

json中的信息：

![image-20211012154832010](https://luochengyu.oss-cn-beijing.aliyuncs.com/image-20211012154832010.png)

## BlockByNumber

![image-20211012162745111](https://luochengyu.oss-cn-beijing.aliyuncs.com/image-20211012162745111.png)

返回样例

![image-20211012221031017](https://luochengyu.oss-cn-beijing.aliyuncs.com/img/image-20211012221031017.png)

## TotTxCount

![image-20211012220537653](https://luochengyu.oss-cn-beijing.aliyuncs.com/img/image-20211012220537653.png)
