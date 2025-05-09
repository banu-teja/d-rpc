# Decentralized & Incentivized RPC Network

This project implements a decentralized protocol where individuals or entities can run blockchain nodes and register as RPC providers. Users pay micro-payments via state channels for RPC requests, creating an alternative to centralized providers like Infura/Alchemy.

## Features

- **Provider Staking & Slashing**: Providers stake tokens. Malicious behavior or poor uptime could lead to slashing, verified through QoS monitoring.
- **Load Balancing & Discovery**: Implements an on-chain registry and discovery mechanism for users to find available, performant nodes.
- **Quality of Service (QoS) Metrics**: Nodes have performance metrics tracked, influencing their probability of selection.
- **Payment Integration**: Payment channels for users to pay per request or subscribe to services.

## Architecture

The project consists of three main components:

1. **Smart Contracts**: Written in Solidity, these handle provider registration, staking, payment channels, and QoS tracking.
2. **Node Service**: Written in Go, this runs the RPC provider service, integrates with the blockchain, and handles payment verification.
3. **Frontend**: A React application for users to discover and interact with the decentralized RPC network.

## Setup Instructions

### Prerequisites

- Node.js (v16+)
- Go (v1.20+)
- Foundry (for smart contract development)

### Smart Contracts

```bash
# Navigate to contracts directory
cd contracts

# Install dependencies
forge install

# Compile contracts
forge build

# Deploy contracts (local development)
forge script script/Deploy.s.sol --rpc-url http://localhost:8545 --broadcast

# Run tests
forge test
```

### Node Service

```bash
# Navigate to node directory
cd node

# Build the node service
go build -o drpcd ./cmd/drpcd

# Run the node service
export FORGE_PRIVATE_KEY="your_private_key"
export CHANNEL_CONTRACT="deployed_channel_contract_address"
export REGISTRY_CONTRACT="deployed_registry_contract_address"
export STK_CONTRACT="deployed_token_contract_address"
./drpcd
```

### Frontend

```bash
# Navigate to ui directory
cd ui

# Install dependencies
npm install

# Start development server
npm run dev
```

## Environment Variables

### Node Service

- `RPC_URL`: Ethereum RPC URL to proxy requests to (default: http://localhost:8545)
- `FORGE_PRIVATE_KEY`: Private key for the node operator
- `CHANNEL_CONTRACT`: Address of the PaymentChannel contract
- `REGISTRY_CONTRACT`: Address of the ProviderRegistry contract
- `STK_CONTRACT`: Address of the StakeToken contract
- `PORT`: HTTP port for the RPC service (default: 8080)
- `STAKE_AMOUNT`: Amount to stake when registering (default: 1 ETH)

## License

MIT
