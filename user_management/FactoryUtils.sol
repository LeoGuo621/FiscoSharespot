pragma solidity 0.4.25;

library FactoryUtils {

    function hashCompareInternal(string memory a, string memory b) internal pure returns (bool) {
        return keccak256(abi.encodePacked(a)) == keccak256(abi.encodePacked(b));
    }

    function isStringExist(string memory s, string[] memory v) internal pure returns (bool) {
        for (uint i; i < v.length; i++)
            if (hashCompareInternal(v[i], s))
                return true;
        return false;
    }
}