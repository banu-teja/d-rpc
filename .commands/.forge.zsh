#!/bin/zsh

# Deployment
deploy_contracts() {
  cd contracts
  export FORGE_PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
  forge script script/DeployContracts.s.sol --broadcast --rpc-url http://127.0.0.1:8545 --private-key $FORGE_PRIVATE_KEY
}

# Funding
fund_user() {
  export STK_CONTRACT=0x5FbDB2315678afecb367f032d93F642f64180aa3
  export USER_ADDRESS=0x70997970C51812dc3A010C7d01b50e0d17dc79C8
  export CHANNEL_DEPOSIT=10000000000000000000
  forge script script/FundUser.s.sol --broadcast --rpc-url http://127.0.0.1:8545 --private-key $FORGE_PRIVATE_KEY
}

# Channel Management
open_channel() {
  export CHANNEL_CONTRACT=0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
  export PROVIDER_ADDRESS=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
  forge script script/OpenChannel.s.sol --broadcast --rpc-url http://127.0.0.1:8545 --private-key 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
}

query_channel() {
  export CHANNEL_ID=$(get_latest_channel_id)
  forge script script/QueryChannel.s.sol --rpc-url http://127.0.0.1:8545
}

close_channel() {
  export CHANNEL_ID=$(get_latest_channel_id)
  export USER_KEY=0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
  forge script script/CloseChannel.s.sol --broadcast --rpc-url http://127.0.0.1:8545 --private-key $FORGE_PRIVATE_KEY
}

# Helper function
get_latest_channel_id() {
  # Check if log file exists
  local log_file="/Users/xai/bhanu/web/web3/d-rpc/contracts/broadcast/OpenChannel.s.sol/31337/run-latest.json"
  if [[ ! -f "$log_file" ]]; then
    echo "Error: OpenChannel logs not found at $log_file" >&2
    echo "Please run 'open_channel' first to create a payment channel" >&2
    return 1
  fi

  # Extract channel ID from logs
  local channel_id=$(jq -r '.receipts[-1].logs[] | select(.topics[0] == "0x506f81b7a67b45bfbc6167fd087b3dd9b65b4531a2380ec406aab5b57ac62152") | .topics[1]' "$log_file" 2>/dev/null)

  # Verify we got a valid channel ID
  if [[ -z "$channel_id" || "$channel_id" == "null" ]]; then
    echo "Error: Could not extract channel ID from logs" >&2
    echo "The OpenChannel transaction may have failed" >&2
    return 1
  fi

  # Ensure proper 0x prefix and length
  if [[ "$channel_id" != "0x"* ]]; then
    channel_id="0x$channel_id"
  fi
  
  if [[ ${#channel_id} -ne 66 ]]; then
    echo "Error: Invalid channel ID length: $channel_id" >&2
    return 1
  fi

  echo "$channel_id"
}

# Main workflow
full_workflow() {
  deploy_contracts
  fund_user
  open_channel
  query_channel
  close_channel
}
