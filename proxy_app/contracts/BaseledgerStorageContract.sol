// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;
import 'node_modules/@openzeppelin/contracts/access/Ownable.sol';
/**
 * @title BaseledgerStorageContract
 * @dev Store & retrieve baseledger transaction id and proof
 */
contract BaseledgerStorageContract is Ownable {
    mapping(string => string) exitProofs;

    function add(string memory key, string memory value) public onlyOwner {
        exitProofs[key] = value;
    }
    
    function get(string memory key) public view returns (string memory) {
        return exitProofs[key];
    }
}