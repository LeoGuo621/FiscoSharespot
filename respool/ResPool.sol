// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.4.25;
pragma experimental ABIEncoderV2;
import "./FactoryUtils.sol";
/**
 * @title Owner
 * @dev Set & change owner
 */
contract res_test {

    address private owner;
    struct Resource{
        string ResID;
        string ResOwnerID;
        string ResBuyerID;
        string ServiceType;
        uint State;
        string AccessCode;
        uint ServiceTime;
        uint Price;
    }
    
    Resource[] public ResPool;
    
    event OwnerSet(address indexed oldOwner, address indexed newOwner);

    modifier isOwner() {
        require(msg.sender == owner, "Caller is not owner");
        _;
    }    

    constructor() public {
        owner = msg.sender; // 'msg.sender' is sender of current call, contract deployer for a constructor
        emit OwnerSet(address(0), owner);
    }
    
    // 返回一个资源的ID，提供给用户以便查询
    // 新增资源，不判断资源池中是否存在，资源ID为时间戳和用户ID组合而成，为唯一ID
    function addRes(
        string _ResOwnerID,
        string _ServiceType,
        uint _ServiceTime,
        uint _Price) isOwner external returns(string){
            
        string memory strtime;
        strtime = FactoryUtils.uint2str(now);
        
        string memory nonce;
        nonce = FactoryUtils.uint2str(ResPool.length);
        
        string memory _ResID;
        _ResID = FactoryUtils.strConcat(strtime, _ResOwnerID);
        _ResID = FactoryUtils.strConcat(_ResID,nonce);

        // uint memory fuck = uint(keccak256(_ResID));
        // _ResID = FactoryUtils.uint2str(fuck);

        Resource memory NewRes = Resource(_ResID,_ResOwnerID,"",_ServiceType,1,"",_ServiceTime,_Price);
        
        ResPool.push(NewRes);
        
        return _ResID;
    }
    
    function preBuyRes(string _ResID,string _ResBuyerID) isOwner external returns(bool,string){

        for (uint i = 0; i < ResPool.length; i++){
            if (FactoryUtils.hashCompareInternal(ResPool[i].ResID,_ResID)){
                if (ResPool[i].State != 1){
                    return (false, "Resource Not For Sale!");
                }
                
                ResPool[i].State = 2;
                ResPool[i].ResBuyerID = _ResBuyerID;
                
                return (true,"Resource preBuy Successfully!");
            }
        }
        
        return (false,"Resource not Exist!");
    }

    function apRes(string _ResID,string _AccessCode) isOwner external returns(bool,string){

        for (uint i = 0; i < ResPool.length; i++){
            if (FactoryUtils.hashCompareInternal(ResPool[i].ResID,_ResID)){
                if (ResPool[i].State != 2){
                    return (false, "Resource Not For Access!");
                }
                
                ResPool[i].State = 3;
                ResPool[i].AccessCode = _AccessCode;
                
                return (true,"AccessCode Provided Successfully!");
            }
        }
        
        return (false,"Resource not Exist!");
    }
    
    function finalRes(string _ResID) isOwner external returns(bool,string){

        for (uint i = 0; i < ResPool.length; i++){
            if (FactoryUtils.hashCompareInternal(ResPool[i].ResID,_ResID)){
                if (ResPool[i].State != 3){
                    return (false, "Resource Not For Finish!");
                }
                
                ResPool[i].State = 4;

                return (true,"Service Completed Successfully!");
            }
        }
        
        return (false,"Resource not Exist!");
    }
    
    function getAllRes() external view returns (Resource[] memory) {
        return ResPool;
    }
    
    function getRes(string _ResID) external view returns(Resource memory){
        bool mark = false;
        uint k = 0;
        for (uint i = 0; i < ResPool.length; i++){
            if (FactoryUtils.hashCompareInternal(ResPool[i].ResID,_ResID)){
                mark = true;
                k = i;
                break;
            }
        }
        require (mark == true, "Resource Not Exist!");
        
        return ResPool[k]; 
    }
    
    function getOwner() external view returns (address) {
        return owner;
    }
}