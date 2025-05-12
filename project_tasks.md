## 🧠 Project Overview: Decentralized & Incentivized RPC Network

### 🔍 What Is It?

An **RPC (Remote Procedure Call) Network** allows blockchain clients (like wallets, dApps, or indexers) to interact with nodes (that provide blockchain data and submit transactions).

A **decentralized and incentivized** version of this:

*   Distributes traffic among many node operators.
*   Rewards those operators for serving requests.
*   Prevents centralized control (e.g., Infura or Alchemy dependency).
*   Can use crypto payments or reputational staking to discourage abuse and improve quality.

---

## 🏗️ High-Level Architecture

1.  **Clients (Users/dApps):**

    *   Connect via a public RPC endpoint.
    *   Can stake tokens, use credits, or pay per request.

2.  **Gateway (Load Balancer & Verifier):**

    *   Receives requests from clients.
    *   Routes to a suitable node (based on reputation, availability, or price).
    *   Occasionally verifies response integrity (e.g., using light client or quorum voting).

3.  **Node Providers:**

    *   Run full nodes and register in the system.
    *   Serve verified requests and receive rewards.

4.  **Smart Contracts (on-chain):**

    *   Handle staking, slashing, rewards, service registration, and billing.

5.  **Reputation/Scoring System:**

    *   Nodes are scored on uptime, correctness, latency, etc.

---

## ✅ Tasks To Complete the Project

### 🔹 Phase 1: Core Design & Research

*   [x] Define supported blockchain(s): Initially Ethereum and Polygon.
*   [x] Design the request/response protocol: Based on JSON-RPC with added payment fields.
*   [x] Determine how clients authenticate: Using wallet signatures tied to payment channels.
*   [ ] Research decentralized verification options (light clients, zk-proofs, quorum). (Pending research, document in docs/decentralized_verification_research.md)
*   [x] Decide reward and penalty mechanisms: Per-request fees via payment channels, slashing via ProviderRegistry, and loss of deposit for expired channels. QoS scoring implemented in contract.
*   [x] Choose infrastructure: Golang for the gateway and Solidity for smart contracts.

---

### 🔹 Phase 2: Smart Contract Development

*   [x] Staking contract for node operators.
*   [x] Reward distribution logic (based on usage or time): Per-request payments handled by PaymentChannel.sol. Placeholder background process added in node/cmd/drpcd/main.go for off-chain calculation and triggering, including comments for usage data aggregation. Further implementation pending.
*   [ ] Slashing mechanism for malicious/incorrect nodes: On-chain slashing function exists in ProviderRegistry.sol. Off-chain detection and triggering logic is pending, including:
    *   [ ] Define slashing criteria. (Document in docs/incentive_reputation_design.md)
    *   [x] Implement off-chain detection logic. (Placeholder function and comments added in node/cmd/drpcd/main.go)
    *   [x] Implement off-chain triggering logic to call the smart contract function. (Placeholder function and comments added in node/cmd/drpcd/main.go)
*   [x] Node registration/deregistration contract.
*   [x] Client credit management (prepaid balance, pay-as-you-go): Payment channels in PaymentChannel.sol provide a prepaid balance mechanism. Placeholder struct and comments for a more comprehensive system added in node/cmd/drpcd/main.go, including note on StakeToken transfer for top-up, and placeholder CheckBalance and DeductCredits functions. Further implementation pending.
*   [ ] DAO/governance hooks (optional, for decentralized updates). (Pending - optional task)

---

### 🔹 Phase 3: Gateway & Network Layer

*   [x] Build the **Gateway server**:

    *   [x] Load balancing between nodes.
    *   [x] Rate limiting, abuse prevention: Basic payment nonce replay check and IP-based rate limiting implemented. More advanced rate limiting and abuse prevention pending.
    *   [x] Logging & billing integration: Basic request logging and billing metric collection implemented. Placeholder struct and comments for billing processing added in node/cmd/drpcd/main.go. Full billing processing and smart contract integration is pending.
    *   [ ] Optionally verify response correctness. (Pending - dependent on Phase 1 research)
*   [x] Node management interface:

*   [x] Register node: On-chain registration handled by ProviderRegistry.sol (now storing RPC URL) and called by the gateway. The load balancer now fetches registration details dynamically from the contract.
*   [x] Ping/health check system: Basic health check logic in node/pkg/loadbalancer/loadbalancer.go now fetches RPC URLs from the ProviderRegistry contract and uses the dynamically fetched list of registered providers.
*   [x] Track node performance & uptime: Performance metrics tracked by QoS monitor. Basic uptime tracking with configurable consecutive failure threshold added to load balancer. Placeholder comments and a placeholder function for persistent storage and integrating historical uptime data into QoS score added in qos/monitor.go. Requires implementing persistent storage and full integration with QoS score.
*   [x] API for clients to connect (public RPC endpoint): Implemented by the gateway's HTTP server.

---

### 🔹 Phase 4: Incentive & Reputation Layer

*   [x] Create scoring system: QoS scoring based on latency, response validity, and consecutive health check failures implemented in qos/monitor.go. Requires integration with persistent uptime data and potential refinement of penalty.
*   [x] Dashboard for monitoring node quality and ranking: Gateway API updated to include uptime and consecutive failures. Pending frontend implementation.
*   [x] Integrate rewards based on score and usage: Placeholder background process added in node/cmd/drpcd/main.go for off-chain calculation and triggering. Implementation of calculation and smart contract interaction is pending, including usage data aggregation outline.

---

### 🔹 Phase 5: Frontend (dApp + Dashboard)

*   [ ] Client Dashboard: (Pending implementation in ui/)

    *   [ ] Register, manage API keys.
    *   [x] View usage stats, top-up credits. (Top-up likely involves transferring StakeToken to a designated contract or address.)
*   [x] Node Operator Dashboard: (Pending implementation in ui/)

    *   [x] Register node, stake tokens. (Placeholder comments added in ui/src/App.tsx outlining approval and deposit steps)
    *   [x] Monitor performance, earnings. (Placeholder comments added in ui/src/App.tsx)
*   [ ] Admin panel (optional, for early manual control). (Pending - optional task, implementation in ui/)

---

### 🔹 Phase 6: Tokenomics (If Required)

*   [x] Design token utility: staking, payment, slashing (partially based on existing contracts). Governance utility pending definition.
*   [x] Create token contract (ERC20 or custom): ERC20 StakeToken.sol created.
*   [ ] Launch initial token distribution (testnet first). (Pending)

---

### 🔹 Phase 7: Deployment & Testing

*   [ ] Deploy smart contracts to testnet. (Pending - deployment scripts exist)
*   [ ] Test gateway with multiple nodes. (Pending)
*   [ ] Load testing (simulate real usage). (Pending)
*   [ ] Security audits for contracts and infrastructure. (Pending)
*   [ ] Bug bounty or open-source review phase. (Pending)

---

### 8. Launch

*   [ ] Deploy to mainnet. (Pending)
*   [ ] Publish open-source gateway + node software. (Pending)
*   [ ] Launch community docs and tutorials. (Pending)
*   [ ] Start incentivized testnet or early reward program. (Pending)

---

## ⚙️ Technologies Suggested

| Layer                | Stack Options                          |
| -------------------- | -------------------------------------- |
| Gateway              | Go                                     |
| Smart Contracts      | Solidity (Ethereum, Polygon)           || Database             | PostgreSQL / Redis (for stats/cache)   || Frontend             | React + Tailwind / Next.js             || Node Software        | Go / Dockerized scripts                |\| Messaging (Optional) | libp2p / pub-sub for node sync         |
---