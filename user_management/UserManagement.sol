pragma solidity 0.4.25;
pragma experimental ABIEncoderV2;
import "./FactoryUtils.sol";

contract UserManagement {
    struct User{
        string id;
        address addr;
        string name;
        bytes32 passwordEncoded;
    }
    User[] public users;

    function addUser(
        string id,
        address addr,
        string  name,
        bytes32 passwordEncoded)
    external
    returns (bool) {
        if (getIndexByID(id) < 2**256 - 1)
            return false;
        User memory u = User(id, addr, name, passwordEncoded);
        users.push(u);
        return true;
    }

    function getIndexByID(string memory id) internal view returns (uint) {
        for (uint i = 0; i < users.length; i++)
            if (FactoryUtils.hashCompareInternal(users[i].id, id))
                return i;
        return 2**256 - 1;
    }

    function getAllUsers() external view returns (User[] memory) {
        return users;
    }
}
