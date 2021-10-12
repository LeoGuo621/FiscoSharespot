# 后端响应框架

[toc]

## 请求列表

|   请求名称    | 请求方法 |    请求返回key    |          功能          |    参数     |
| :-----------: | :------: | :---------------: | :--------------------: | :---------: |
|  BlockNumber  |   GET    |   block_number    |    查询当前区块高度    |             |
|   NodePeers   |   GET    |    node_peers     |    查询P2P节点信息     |             |
|  SyncStatus   |   GET    |    sync_status    | 查询群组内同步状态信息 |             |
| BlockByNumber |   POST   | block_information | 根据区块号查询区块信息 | BlockNumber |
|  TotTxCount   |   GET    |   tot_trans_cnt   |    查询当前总交易数    |             |



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
