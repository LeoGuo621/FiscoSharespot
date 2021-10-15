pragma solidity 0.4.25;
pragma experimental ABIEncoderV2;

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

    function uint2str(uint i) internal pure returns (string c) {
        if (i == 0) return "0";
        uint j = i;
        uint length;
        while (j != 0){
            length++;
            j /= 10;
        }
        bytes memory bStr = new bytes(length);
        uint k = length - 1;
        while (i != 0){
            bStr[k--] = byte(48 + i % 10);
            i /= 10;
        }
        c = string(bStr);
    }

    function strConcat(string _a, string _b) internal pure returns (string){
        bytes memory _ba = bytes(_a);
        bytes memory _bb = bytes(_b);
        string memory ret = new string(_ba.length + _bb.length);
        bytes memory bret = bytes(ret);
        uint k = 0;
        for (uint i = 0; i < _ba.length; i++)bret[k++] = _ba[i];
        for (i = 0; i < _bb.length; i++) bret[k++] = _bb[i];
        return string(ret);
    }
}