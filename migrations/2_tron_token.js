var tronToken = artifacts.require("./TronToken.sol");

module.exports = function(deployer, network, accounts) {
  deployer.deploy(tronToken, accounts[0]);
};
