var Web3 = require("web3"); var web3 = new Web3("http://localhost:8545"); web3.eth.getCode("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512").then(console.log)
