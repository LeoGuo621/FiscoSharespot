pragma solidity ^0.4.25;
pragma experimental ABIEncoderV2;
import "./FactoryUtils.sol";

contract ResourceManagement {

    address public owner;
    struct Resource{
        string ResourceID;
        string ResourceOwnerID;
        string ResourceBuyerID;
        string ServiceType;
        string Status;
        string AccessCode;
        uint ServiceTime;
        uint Price;
    }

    Resource[] private ResourcePool;

    modifier isOwner() {
        require(msg.sender == owner, "Caller is not owner");
        _;
    }

    constructor() public {
        owner = msg.sender; // 'msg.sender' is sender of current call, contract deployer for a constructor
    }

    //Check the ResourceID ,if valid return true and position, if not return false and 0
    function isResourceValidByID(string memory ResourceID) public view returns (bool) {
        require(!FactoryUtils.hashCompareInternal("", ResourceID));
        for (uint i = 0; i < ResourcePool.length; i++)
            if (FactoryUtils.hashCompareInternal(ResourcePool[i].ResourceID, ResourceID))
                return true;
        return false;
    }

    // 返回一个资源的ID，提供给用户以便查询
    // 新增资源，不判断资源池中是否存在，资源ID为时间戳和用户ID组合而成，为唯一ID
    function addResource(
        string _ResourceOwnerID,
        string _ServiceType,
        uint _ServiceTime,
        uint _Price) isOwner external returns(string){

        require(!FactoryUtils.hashCompareInternal("", _ResourceOwnerID),"Resource Owner Empty!");
        require(!FactoryUtils.hashCompareInternal("", _ServiceType),"Service Type Not Set!");


        string memory strtime;
        strtime = FactoryUtils.uint2str(now);

        string memory nonce;
        nonce = FactoryUtils.uint2str(ResourcePool.length);

        string memory _ResourceID;
        _ResourceID = FactoryUtils.strConcat(strtime, nonce);
        _ResourceID = FactoryUtils.strConcat(_ResourceID,_ResourceOwnerID);

        // uint memory fuck = uint(keccak256(_ResourceID));
        // _ResourceID = FactoryUtils.uint2str(fuck);

        Resource memory NewRes = Resource(_ResourceID,_ResourceOwnerID,"",_ServiceType,"ForSale","",_ServiceTime,_Price);

        ResourcePool.push(NewRes);

        return _ResourceID;
    }

    function preBuyResource(string _ResourceID,string _ResourceBuyerID) isOwner external returns(bool){

        require(
            !FactoryUtils.hashCompareInternal("", _ResourceID),
            "Resource ID Empty!"
        );
        require(
            !FactoryUtils.hashCompareInternal("", _ResourceBuyerID),
            "Resource Buyer ID Empty!"
        );

        for (uint i = 0; i < ResourcePool.length; i++){
            if (FactoryUtils.hashCompareInternal(ResourcePool[i].ResourceID,_ResourceID)){
                // if (!FactoryUtils.hashCompareInternal(ResourcePool[i].Status,"ForSale")){
                //     return (false, "Resource Not For Sale!");
                // }
                require(
                    FactoryUtils.hashCompareInternal(ResourcePool[i].Status,"ForSale"),
                    "Resource Not For Sale!"
                );

                ResourcePool[i].Status = "PreBuy";
                ResourcePool[i].ResourceBuyerID = _ResourceBuyerID;

                return true;
            }
        }

        require(false,"Resource not Exist!");
        return false;
    }

    function grantResource(string _ResourceID,string _AccessCode) isOwner external returns(bool){

        require(
            !FactoryUtils.hashCompareInternal("", _ResourceID),
            "Resource ID Empty!"
        );
        require(
            !FactoryUtils.hashCompareInternal("", _AccessCode),
            "Access Code Empty!"
        );

        for (uint i = 0; i < ResourcePool.length; i++){
            if (FactoryUtils.hashCompareInternal(ResourcePool[i].ResourceID,_ResourceID)){
                require(
                    FactoryUtils.hashCompareInternal(ResourcePool[i].Status,"PreBuy"),
                    "Resource Not For Serving!"
                );

                ResourcePool[i].Status = "Serving";
                ResourcePool[i].AccessCode = _AccessCode;

                return true;
            }
        }

        require(false,"Resource not Exist!");
        return false;
    }

    function finishResource(string _ResourceID) isOwner external returns(bool){

        require(
            !FactoryUtils.hashCompareInternal("", _ResourceID),
            "Resource ID Empty!"
        );

        for (uint i = 0; i < ResourcePool.length; i++){
            if (FactoryUtils.hashCompareInternal(ResourcePool[i].ResourceID,_ResourceID)){
                require(
                    FactoryUtils.hashCompareInternal(ResourcePool[i].Status,"Serving"),
                    "Resource Can't Be Finished!"
                );

                ResourcePool[i].Status = "ServiceFinished";

                return true;
            }
        }
        require(false,"Resource not Exist!");
        return false;
    }

    function getAllResource() isOwner external view returns (Resource[] memory) {
        return ResourcePool;
    }

    //function isStatusValid(string memory ResourceStatus) public view returns (bool) {
    //    require(
    //        !(FactoryUtils.hashCompareInternal("", ResourceStatus)),
    //        "Resource Status Empty!"
    //    );
    //    if (FactoryUtils.hashCompareInternal("ForSale", ResourceStatus)){
    //        return true;
    //    }
    //    if (FactoryUtils.hashCompareInternal("PreBuy", ResourceStatus)){
    //        return true;
    //    }
    //    if (FactoryUtils.hashCompareInternal("Serving", ResourceStatus)){
    //        return true;
    //    }
    //    if (FactoryUtils.hashCompareInternal("ServiceFinished", ResourceStatus)){
    //        return true;
    //    }
    //    return false;
    //}

    // function getResourceByStatus(string _Status) isOwner external view returns(Resource[] memory){
    //     require(
    //         !(FactoryUtils.hashCompareInternal("", _Status)),
    //         "Resource Status Empty!"
    //     );
    //     require(
    //         isStatusValid(_Status),
    //         "Not Valid Status!"
    //     );
    //     Resource[] memory ResourceByStatus;
    //     for (uint i = 0; i < ResourcePool.length; i++){
    //         Resource memory cur = ResourcePool[i];
    //         if (FactoryUtils.hashCompareInternal(cur.Status,_Status)){
    //             ResourceByStatus.push(cur);
    //             // return ResourcePool[i];
    //         }
    //     }
    //     return ResourceByStatus;
    // }

    function getResource(string _ResourceID) isOwner external view returns(Resource memory){
        bool mark = false;
        uint k = 0;
        for (uint i = 0; i < ResourcePool.length; i++){
            if (FactoryUtils.hashCompareInternal(ResourcePool[i].ResourceID,_ResourceID)){
                mark = true;
                k = i;
                break;
            }
        }
        require (mark == true, "Resource Not Exist!");

        return ResourcePool[k];
    }
}