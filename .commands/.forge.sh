anvil --port 8545
export FORGE_PRIVATE_KEY=<PASTE_FIRST_ANVIL_PRIVATE_KEY>
cd /Users/xai/bhanu/web/web3/d-rpc/contracts
forge build
forge script script/DeployContracts.s.sol \
  --broadcast \
  --rpc-url http://127.0.0.1:8545 \
  --private-key $FORGE_PRIVATE_KEY

forge script script/RegisterProvider.s.sol \
  --broadcast \
  --rpc-url http://127.0.0.1:8545 \
  --private-key $FORGE_PRIVATE_KEY



anvil --port 8545
export FORGE_PRIVATE_KEY=0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
cd contracts

forge build

forge script script/DeployContracts.s.sol \
  --broadcast \
  --rpc-url http://127.0.0.1:8545 \
  --private-key $FORGE_PRIVATE_KEY

forge script script/RegisterProvider.s.sol \
  --broadcast \
  --rpc-url http://127.0.0.1:8545 \
  --private-key $FORGE_PRIVATE_KEY

export CHANNEL_CONTRACT=0x9fE46736679d2D9a65F0992F2272dE9c7fa6e0
export STK_CONTRACT=0x5FbDB2315678afecb367f032d93F642f64180aa3
export PROVIDER_ADDRESS=0x5b73C5498c1E3b4dbA84de0F1833c4a029d90519
export CHANNEL_DEPOSIT=100000000000000000000
export CHANNEL_DURATION=86400


export STK_CONTRACT=0x5FbDB2315678afecb367f032d93F642f64180aa3
export CHANNEL_CONTRACT=0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
export PROVIDER_ADDRESS=0x5b73C5498c1E3b4dbA84de0F1833c4a029d90519
export CHANNEL_DEPOSIT=10000000000000000000  # 10 ETH
export CHANNEL_DURATION=86400

forge script script/OpenChannel.s.sol \
  --broadcast \
  --rpc-url http://127.0.0.1:8545 \
  --private-key $FORGE_PRIVATE_KEY

  
forge script script/OpenChannel.s.sol \
  --broadcast \
  --rpc-url http://127.0.0.1:8545 \
  --private-key $FORGE_PRIVATE_KEY



# Export the user’s key (the one that opened the channel)
export USER_KEY=0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d

# Run the close‐channel script
forge script script/CloseChannel.s.sol \
  --broadcast \
  --rpc-url http://127.0.0.1:8545 \
  --private-key $FORGE_PRIVATE_KEY


export CHANNEL_ADDRESS=$(jq -r '.transactions[2].contractAddress' /Users/xai/bhanu/web/web3/d-rpc/contracts/broadcast/DeployContracts.s.sol/31337/run-latest.json)
export TOKEN_ADDRESS=$(jq -r '.transactions[0].contractAddress' /Users/xai/bhanu/web/web3/d-rpc/contracts/broadcast/DeployContracts.s.sol/31337/run-latest.json)
export CHANNEL_ID=0xeda7ce97f41348da96edbadff6d4f59aba9c70653f45a69e36b8be6de613151f




export RPC_URL="http://127.0.0.1:8545"
# acct #1 = deployer & provider
export FORGE_PRIVATE_KEY="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

# acct #2 = user
export USER_KEY="0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"

# these will match the DeployContracts output:
export STK_CONTRACT="0x5FbDB2315678afecb367f032d93F642f64180aa3"
export REGISTRY_CONTRACT="0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
export CHANNEL_CONTRACT="0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
export PROVIDER_ADDRESS="0x5b73C5498c1E3b4dbA84de0F1833c4a029d90519"

# channel parameters (10 ETH, 1 day)
export CHANNEL_DEPOSIT="10000000000000000000"
export CHANNEL_DURATION="86400"