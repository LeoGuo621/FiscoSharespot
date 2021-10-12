# 后端响应框架

[toc]

## 请求列表

|      请求名称       | 请求方法 |    请求返回key    |          功能          |    参数    |
| :-----------------: | :------: | :---------------: | :--------------------: | :--------: |
|     blocknumber     |   GET    |   block_number    |    查询当前区块高度    |            |
|      nodepeers      |   GET    |    node_peers     |    查询P2P节点信息     |            |
|     syncstatus      |   GET    |    sync_status    | 查询群组内同步状态信息 |            |
|    blockbynumber    |   POST   | block_information | 根据区块号查询区块信息 | blocknumer |
| tottransactioncount |   GET    |   tot_trans_cnt   |    查询当前总交易数    |            |



## blocknumber

![image-20211012150301633](https://luochengyu.oss-cn-beijing.aliyuncs.com/image-20211012150301633.png)

## nodepeers

![image-20211012150427261](https://luochengyu.oss-cn-beijing.aliyuncs.com/image-20211012150427261.png)

## syncstatus

![image-20211012154705083](https://luochengyu.oss-cn-beijing.aliyuncs.com/image-20211012154705083.png)

json中的信息：

![image-20211012154832010](https://luochengyu.oss-cn-beijing.aliyuncs.com/image-20211012154832010.png)

## blockbynumber

![image-20211012162745111](https://luochengyu.oss-cn-beijing.aliyuncs.com/image-20211012162745111.png)

返回样例

![image-20211012162810188](https://luochengyu.oss-cn-beijing.aliyuncs.com/image-20211012162810188.png)

## tottransactioncount

![image-20211012214917971](https://luochengyu.oss-cn-beijing.aliyuncs.com/img/image-20211012214917971.png)
