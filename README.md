# ShareSpot区块链开放式网络共享

一个基于FISCO BCOS区块链网络的无线网络资源共享平台

- [FISCO BCOS](https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/introduction.html)是由国内企业主导研发、对外开源、安全可控的企业级金融联盟链底层平台，由金链盟开源工作组协作打造，并于2017年正式对外开源。
- **ShareSpot**最初由东南大学移动通信国家重点实验室的团队开发。
- 版权所有为开发团队。未经团队授权，严禁任何修改和商业用途。

## 项目介绍

ShareSpot是一个基于FISCO BCOS区块链的无线资源共享与交易平台，用户既可以是无线资源的提供者，又可以是无线资源的消费者。拥有闲置无线资源的用户，可以通过 ShareSpot 申报无线资源的各类信息，并以此得到收益。需要接入无线资源的用户，可以通过 ShareSpot 查看所有可购买的无线资源，在支付相应服务费用后可以享受合约规定时间内的无线服务。

### 项目环境及开发工具

项目环境要求：

- go
- Docker

开发工具：

- Goland
- Remix
- FISCO BCOS

开发环境：

- GO：1.17.2
- Mysql：5.7
- soidity：^0.4.25

### 项目演示

ShareSpot项目Demo地址：http://exhard.cn/ss/ui-login.html

### 项目架构图

![image-20211017151558581](https://luochengyu.oss-cn-beijing.aliyuncs.com/img/image-20211017151558581.png)

## 后端接口列表

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

## 用途

ShareSpot是一个面向无线资源持有者和普通用户的Wi-Fi等资源的共享平台。对于每一个用户，他们可以在ShareSpot平台发布他们闲置的无线网络资源（Wi-Fi、蓝牙、计算资源等）供用户购买。也可以以消费者的身份在ShareSpot平台向资源持有者申请购买他们期望的资源。

ShareSpot打破了不同资源所有者之间的障碍，为用户提供了一个安全的公共平台，可以从任何地方访问可用的无线网络资源。借助FISCO-BCOS，资源共享与交易可以非常安全、顺畅。此外，由于区块链的去中心化特性，项目代码可以直接分发到任何移动设备和商业硬件。
