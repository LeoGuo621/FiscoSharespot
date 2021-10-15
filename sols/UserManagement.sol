pragma solidity 0.4.25;
pragma experimental ABIEncoderV2;
import "./FactoryUtils.sol";

contract UserManagement {

    event Transfer(string from, string to, uint amount);

    struct User{
        string id;
        string pubKey;
        string name;
        string passwordEncoded;
        uint balance;
    }

    address public admin;

    User[] private users;

    constructor() {
        admin = msg.sender;
    }

    modifier onlyAdmin {
        require(
            msg.sender == admin,
            "Only admin can call this function."
        );
        _;
    }

    function addUser(
        string id,
        string pubKey,
        string  name,
        string passwordEncoded)
    external
    onlyAdmin
    returns (bool) {
        require(!isUserValidByID(id));
        // The initial balance of all users is 1000
        uint subsidy = 1000;
        User memory u = User(id, pubKey, name, passwordEncoded, subsidy);
        users.push(u);
        return true;
    }

    function changePasswordEncoded(
        string id,
        string targetPasswordEncoded)
    external
    onlyAdmin
    returns (bool) {
        uint index = getIndexByID(id);
        if (index == 2**256 - 1)
            return false;
        users[index].passwordEncoded = targetPasswordEncoded;
        return true;
    }

    function deleteUser(string id) external onlyAdmin {
        uint index = getIndexByID(id);
        // if (index == 2**256 - 1)
        //     return;
        delete users[index];
    }

    function getIndexByID(string memory id) internal view returns (uint) {
        require(isUserValidByID(id));
        for (uint i = 0; i < users.length; i++)
            if (FactoryUtils.hashCompareInternal(users[i].id, id))
                return i;
        return 2**256 - 1;
    }

    function getAllUsers() external view onlyAdmin returns (User[] memory) {
        return users;
    }

    function getUserByID(string memory id) public view returns (User memory) {
        require(isUserValidByID(id));
        for (uint i = 0; i < users.length; i++)
            if (FactoryUtils.hashCompareInternal(users[i].id, id))
                return users[i];
        return User("", "", "", "", 0);
    }

    function getUserByPubKey(string memory pubKey) public view returns (User memory) {
        require(isUserValidByPubKey(pubKey));
        for (uint i = 0; i < users.length; i++)
            if (FactoryUtils.hashCompareInternal(users[i].pubKey, pubKey))
                return users[i];
        return User("", "", "", "", 0);
    }

    function getBalanceByID(string memory id) public view returns (uint) {
        require(isUserValidByID(id));
        for (uint i = 0; i < users.length; i++)
            if (FactoryUtils.hashCompareInternal(users[i].id, id))
                return users[i].balance;
        return 0;
    }

    function getBalanceByPubKey(string memory pubKey) public view returns (uint) {
        require(isUserValidByPubKey(pubKey));
        for (uint i = 0; i < users.length; i++)
            if (FactoryUtils.hashCompareInternal(users[i].pubKey, pubKey))
                return users[i].balance;
        return 0;
    }

    function isUserValidByID(string memory id) public view returns (bool) {
        require(!FactoryUtils.hashCompareInternal("", id));
        for (uint i = 0; i < users.length; i++)
            if (FactoryUtils.hashCompareInternal(users[i].id, id))
                return true;
        return false;
    }

    function isUserValidByPubKey(string memory pubKey) public view returns (bool) {
        require(!FactoryUtils.hashCompareInternal("", pubKey));
        for (uint i = 0; i < users.length; i++)
            if (FactoryUtils.hashCompareInternal(users[i].pubKey, pubKey))
                return true;
        return false;
    }

    function getPubKeyByID(string memory id) public view returns (string) {
        require(isUserValidByID(id));
        for (uint i = 0; i < users.length; i++)
            if (FactoryUtils.hashCompareInternal(users[i].id, id))
                return users[i].pubKey;
        return "";
    }

    function getIDByPubKey(string memory pubKey) public view returns (string) {
        require(isUserValidByPubKey(pubKey));
        for (uint i = 0; i < users.length; i++)
            if (FactoryUtils.hashCompareInternal(users[i].pubKey, pubKey))
                return users[i].id;
        return "";
    }

    function transfer(string from, string to, uint amount) external onlyAdmin returns (bool) {
        require(isUserValidByID(from));
        require(isUserValidByID(to));
        uint sponsorIdx = getIndexByID(from);
        uint receiverIdx = getIndexByID(to);
        require(users[sponsorIdx].balance >= amount);
        users[sponsorIdx].balance = users[sponsorIdx].balance - amount;
        users[receiverIdx].balance = users[receiverIdx].balance + amount;
        emit Transfer(from, to, amount);
        return true;
    }
}
