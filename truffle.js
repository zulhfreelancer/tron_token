module.exports = {
  networks: {
    development: {
      host: "localhost",
      port: 7545,
      network_id: "*"
    },
    geth_private: {
      host: "localhost",
      port: 8545,
      network_id: "*"
    }
  }
};
