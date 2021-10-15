# 后端响应框架

[toc]

## 请求列表

|   请求名称    | 请求方法 |               功能               |                      参数                      |
| :-----------: | :------: | :------------------------------: | :--------------------------------------------: |
|  blockNumber  |   GET       |         查询当前区块高度         |                                                |
|   nodePeers   |   GET       |         查询P2P节点信息          |                                                |
|  syncStatus   |   GET        |      查询群组内同步状态信息      |                                                |
| blockByNumber |   POST    |      根据区块号查询区块信息      |                  blockNumber                   |
|  totalTransactionCount   |   GET     |         查询当前总交易数         |                                                |
|   resourceAddress   |   GET                     |           获取部署的资源池合约地址信息            |                                                |
|   resourceAdd   |   POST                  | 新增一个资源 |               ownerID<br>serviceType<br>serviceTime<br>price                                 |
|    resourcePreBuy     |   POST                   |             预购一个资源             | resourceID<br>buyerID |
|    resourceGrant     |   POST                    |      Ap提供接入码      |                     resourceID<br>accessCode                      |
|   resourceFinish   |   POST                    |               资源服务结束               |                resourceID           |
|     resourceIDGet     |   POST          |            根据资源号获取资源信息            |             resourceID               |
|   resourceAllGet    |   GET     |           获取资源池所有资源信息           |                                           |
|    resourceStatusGet| POST |  根据资源状态获取资源信息ForSale,PreBuy,Serving,ServiceFinished|resourceStatus|



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
