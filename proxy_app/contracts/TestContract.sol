// SPDX-License-Identifier: GPL-3.0
// deployed at 0x2964196Ea65E0Ad37B39C9D8c0dA92F578c00964
pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Storage
 * @dev Store & retrieve value in a variable
 */
contract Storage {

    uint256 number;
    
    mapping(string => string) txHashes;

    function add(string memory key, string memory value) public {
        txHashes[key] = value;
    }

    
    function get(string memory key) public view returns (string memory) {
        return txHashes[key];
    }
}