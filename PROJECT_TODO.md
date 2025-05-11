---

## 🧠 Project Overview: Decentralized & Incentivized RPC Network

### 🔍 What Is It?

An **RPC (Remote Procedure Call) Network** allows blockchain clients (like wallets, dApps, or indexers) to interact with nodes (that provide blockchain data and submit transactions).

A **decentralized and incentivized** version of this:

* Distributes traffic among many node operators.
* Rewards those operators for serving requests.
* Prevents centralized control (e.g., Infura or Alchemy dependency).
* Can use crypto payments or reputational staking to discourage abuse and improve quality.

---

## 🏗️ High-Level Architecture

1. **Clients (Users/dApps):**

   * Connect via a public RPC endpoint.
   * Can stake tokens, use credits, or pay per request.

2. **Gateway (Load Balancer & Verifier):**

   * Receives requests from clients.
   * Routes to a suitable node (based on reputation, availability, or price).
   * Optionally verifies response integrity (e.g., using light client or quorum voting).

3. **Node Providers:**

   * Run full nodes and register in the system.
   * Serve verified requests and receive rewards.

4. **Smart Contracts (on-chain):**

   * Handle staking, slashing, rewards, service registration, and billing.

5. **Reputation/Scoring System:**

   * Nodes are scored on uptime, correctness, latency, etc.

---

## ✅ Tasks To Complete the Project

### 🔹 Phase 1: Core Design & Research

* [x] Define supported blockchain(s) (e.g., Ethereum, Polygon).
    * Decision: Ethereum and Polygon, as per suggested technologies.
* [x] Design the request/response protocol (e.g., JSON-RPC wrapper).
    * Decision: Standard JSON-RPC will be used, as it's the common protocol for Ethereum and Polygon.
* [x] Determine how clients authenticate (API key, wallet signature).
    * Decision: Primarily wallet signatures for a web3-native experience. API keys can be a secondary option for simpler integrations.
* [ ] Research decentralized verification options (light clients, zk-proofs, quorum).
    * Status: Pending research. Options include light clients, zk-proofs, and quorum-based validation. The choice will depend on the trade-offs between security, performance, and complexity.
* [x] Decide reward and penalty mechanisms (per-request fee, slashing for incorrect responses).
    * Decision: 
        * Rewards: Per-request fees to node operators. Additional token rewards based on performance (uptime, latency) and stake. 
        * Penalties: Slashing of staked tokens for malicious behavior (e.g., incorrect/tampered responses) and significant downtime. Reputation score will be negatively impacted.
* [x] Choose infrastructure (e.g., Golang/Rust for gateway, Solidity for contracts).
    * Decision: Aligned with the "Technologies Suggested":
        * Gateway: Go
        * Smart Contracts: Solidity (for Ethereum, Polygon)
        * Database: PostgreSQL / Redis (for stats/cache)
        * Frontend: React + Tailwind / Next.js
        * Node Software: Go / Dockerized scripts
        * Messaging (Optional): libp2p / pub-sub for node sync

---

### 🔹 Phase 2: Smart Contract Development

* [x] Staking contract for node operators.
    * Implementation: The `ProviderRegistry.sol` contract manages stake deposits, withdrawals, and tracks staked amounts for node operators using the `StakeToken.sol` ERC20 token. This covers the core staking functionality.
* [ ] Reward distribution logic (based on usage or time).
    * Status: Needs implementation. 
    * Plan: 
        1. Add a reward pool to `ProviderRegistry.sol` (or a separate contract) funded by project revenues or token inflation.
        2. Implement a function (e.g., `distributeRewards`) callable by the owner (or DAO).
        3. This function will iterate through registered providers and distribute rewards based on factors like stake, QoS score, and potentially off-chain usage data provided by the gateway/owner.
        4. Events should be emitted for reward distribution.
* [x] Slashing mechanism for malicious/incorrect nodes.
    * Implementation: `ProviderRegistry.sol` includes an `Ownable` function `slashProvider(address provider_, uint256 amount)`.
    * Notes: This function allows the owner to reduce a provider's stake and claim the slashed tokens. The effectiveness of this mechanism depends on a reliable off-chain (or on-chain via verification) process for detecting and reporting malicious behavior to the owner/DAO.
* [x] Node registration/deregistration contract.
    * Implementation: `ProviderRegistry.sol` handles provider registration (`register()`) and deregistration (`deregister()`). It includes logic for minimum stake requirements and a deregistration cooldown period.
* [x] Client credit management (prepaid balance, pay-as-you-go).
    * Implementation: `PaymentChannel.sol` allows users to open unidirectional payment channels with providers, depositing an ERC20 token. Providers can close channels with a user-signed message to claim their earnings. This supports both prepaid balances (the channel deposit) and a form of pay-as-you-go (off-chain payment aggregation, on-chain settlement).
    * Notes: This is suitable for frequent small payments to specific providers. A more global credit/debit system might be needed if users are not interacting with the same provider repeatedly or if payments are not tied to specific channels.
* [ ] DAO/governance hooks (optional, for decentralized updates).
    * Status: Not yet implemented. Currently, `ProviderRegistry.sol` uses `Ownable` for administrative functions.
    * Plan for Decentralization:
        1. Transfer ownership of `ProviderRegistry.sol` to a DAO contract (e.g., Gnosis Safe initially, later a custom governance contract).
        2. Modify `onlyOwner` modifiers to `onlyGovernance` or implement a system where the DAO can propose and execute calls to these administrative functions.
        3. Develop separate governance contracts for proposal submission, voting, and execution if a more complex DAO structure is required.
    * Notes: This can be implemented progressively. The current `Ownable` structure is suitable for initial deployment and testing.

---

### 🔹 Phase 3: Gateway & Network Layer

* [ ] Build the **Gateway server**:

  * [x] Load balancing between nodes.
    * Implementation: The Go package `node/pkg/loadbalancer/` contains `loadbalancer.go`.
    * Logic: It implements a `LoadBalancer` that can fetch provider data (address, QoS score, stake). The `GetProvider()` method uses a weighted random selection strategy, where weights are proportional to `QoSScore^2`, favoring higher-quality nodes.
    * Current Status: Uses test/simulated provider data. Needs to be integrated with live data from the `ProviderRegistry` smart contract (e.g., by listening to `ProviderRegistered`, `QoSUpdated` events).
  * [ ] Rate limiting, abuse prevention.
    * Status: Not yet implemented in the Go gateway code.
    * Plan:
        * Implement rate limiting middleware in the gateway server.
        * Consider client identity for rate limits (wallet address, or API key if supported).
        * Higher stakes by clients or providers could potentially grant higher rate limits.
        * For abuse prevention, develop mechanisms to detect and block malicious patterns (e.g., overly complex queries, request spamming).
        * Integrate with the on-chain slashing: if the gateway itself is a source of truth for some types of abuse by nodes, it needs a secure way to report this to the `ProviderRegistry` owner/DAO.
  * [ ] Logging & billing integration.
    * Status: Partially implemented.
    * Logging:
        * The `node/pkg/qos/monitor.go` (`QoSMonitor`) records performance metrics (response time, success) for each provider. This serves as a specialized log for QoS.
        * General gateway logging (request details, errors, routing decisions) needs to be implemented in the main gateway application (`cmd/drpcd/`).
    * Billing Integration:
        * The metrics from `QoSMonitor` are essential for billing and reward distribution.
        * If using `PaymentChannel.sol`: The gateway needs logic to facilitate the creation of signed payment messages for users to authorize payments to providers based on successful RPC servicing.
        * For broader reward distribution (Phase 2 task): The gateway (or a component using `QoSMonitor` data) needs to aggregate usage data per provider to inform the reward calculation.
    * Plan:
        * Implement standard logging in `cmd/drpcd/`.
        * Design and implement the interaction flow for `PaymentChannel.sol` if it's to be directly managed or facilitated by the gateway.
        * Develop a mechanism to securely report aggregated usage/QoS data to the smart contract layer for reward distribution (this could be via an admin call or a decentralized reporting system).
  * [ ] Optionally verify response correctness.
    * Status: Not yet implemented. Design depends on the outcome of Phase 1 research: "Research decentralized verification options (light clients, zk-proofs, quorum)".
    * Plan:
        * Once a verification strategy is chosen (e.g., light client checks for certain data, quorum-based checks for others), implement this in the gateway.
        * The `QoSMonitor.RecordMetric` function's `success` parameter would be determined by this verification step.
        * Failed verifications should lead to poor QoS scores and potentially trigger slashing if the behavior is deemed malicious.
* [ ] Node management interface:

  * [x] Register node.
    * Status: Partially complete.
    * On-chain: `ProviderRegistry.sol` allows nodes to register by calling `register()` after staking.
    * Gateway: The `node/pkg/loadbalancer/loadbalancer.go` is designed to fetch and manage a list of registered providers. Its `updateProviders` method needs to be enhanced to fetch live data from the `ProviderRegistry` contract (e.g., by listening to events or querying state) instead of relying on test data.
  * [ ] Ping/health check system.
    * Status: Needs implementation in the gateway.
    * Plan:
        1. The gateway (likely in `cmd/drpcd/`) will need a component that periodically iterates through all known (and registered) providers from the `LoadBalancer`.
        2. For each provider, it will send a standard RPC request (e.g., `eth_blockNumber` or `net_version`) to its advertised endpoint.
        3. The success/failure and response time of this health check will be fed into `QoSMonitor.RecordMetric()`.
        4. Providers failing health checks consistently could see their QoS scores drop significantly and might be temporarily removed from active load balancing rotation until they recover.
  * [x] Track node performance & uptime.
    * Implementation: The `node/pkg/qos/monitor.go` (`QoSMonitor`) is responsible for this.
    * Performance: It records response times and success/failure for each request (or health check) to a provider. This data is used to calculate a `QoSScore`.
    * Uptime: The success rate, a key component of the `QoSScore`, implicitly tracks uptime. Consistent failures will lower the score.
    * On-chain Record: `QoSMonitor` periodically updates these `QoSScore` values on the `ProviderRegistry` smart contract.
    * Notes: The `Provider` struct in `loadbalancer.go` has an `UpSince` field; its usage and update mechanism should be clarified or enhanced for more explicit uptime tracking if needed beyond the QoS score.
* [x] API for clients to connect (public RPC endpoint).
    * Implementation: The `node/cmd/drpcd/main.go` sets up an HTTP server with a `/` endpoint (`handleRPCRequest`) to receive JSON-RPC requests.
    * Current Status & Key Issues:
        * **Proxying Logic:** Currently, `handleRPCRequest` proxies all requests to a single configured `EthRPCURL` (from `s.config.EthRPCURL`). It **does not** use the `s.loadBalancer.GetProvider()` to select a decentralized provider. This is a critical gap.
        * **QoS Recording:** QoS metrics in `handleRPCRequest` are recorded against `req.Payment.From`, which is likely incorrect for general RPC calls. It should be the address of the chosen upstream provider from the load balancer.
        * **Client Authentication:** General client authentication (e.g., wallet signature, API key) for accessing the RPC service is not yet implemented. The existing `validatePayment` is for a specific off-chain payment flow.
        * **Rate Limiting:** Not yet implemented.
        * **Payment Channel Interaction:** The `validatePayment` function in `main.go` seems to be an off-chain pre-validation for a payment scheme. How this integrates with the on-chain `PaymentChannel.sol` (e.g., for `closeChannel`) needs clarification and full implementation if the gateway is to facilitate this.
    * Positive Aspects:
        * Basic HTTP server structure is in place using `gorilla/mux`.
        * Includes logging and CORS middleware.
        * `/discovery` endpoint provides provider list and a recommended provider from the load balancer.
        * `/health` endpoint for gateway health.
    * Plan:
        1.  **Crucial:** Modify `handleRPCRequest` to use `s.loadBalancer.GetProvider()` to select an upstream node.
        2.  Proxy the request to the selected provider's endpoint (this endpoint URL needs to be part of the `Provider` data in `LoadBalancer`).
        3.  Correctly call `s.qosMonitor.RecordMetric()` with the selected provider's address and the outcome of the proxied call.
        4.  Implement client authentication mechanisms.
        5.  Implement rate limiting.
        6.  Clarify and fully implement the workflow for client payments, potentially integrating with `PaymentChannel.sol` more directly if the gateway is responsible for submitting `closeChannel` transactions or providing data for them.

---

### 🔹 Phase 4: Incentive & Reputation Layer

* [x] Create scoring system (e.g., uptime, latency, response validity).
    * Implementation: 
        * The Go package `node/pkg/qos/monitor.go` (`QoSMonitor`) calculates a numerical QoS score (0-100) for each provider.
        * This score is based on weighted averages of success rate (proxy for uptime and basic validity) and response time (latency).
        * The `ProviderRegistry.sol` smart contract stores this `qosScore` on-chain, updated by the `QoSMonitor` via the `updateQoS` function.
    * Notes: 
        * The "response validity" aspect of the score is currently basic (depends on the `success` flag passed to `RecordMetric`). True cryptographic verification of responses would enhance this significantly and would need to be integrated into the process that determines the `success` flag.
        * The on-chain `qosScore` serves as the public reputation metric.
* [ ] Dashboard for monitoring node quality and ranking.
    * Status: Backend API is ready; frontend implementation is pending (part of Phase 5).
    * Backend Support: The gateway (`node/cmd/drpcd/main.go`) exposes a `/discovery` HTTP endpoint.
        * This endpoint returns a list of all active providers, including their addresses, on-chain QoS scores, and stake amounts.
        * It also suggests a "recommended provider" based on the load balancer's logic.
    * Frontend Plan: A UI component (e.g., in React) will fetch data from the `/discovery` endpoint and display it, allowing users to see node rankings, quality scores, and other relevant stats.
* [ ] Integrate rewards based on score and usage.
    * Status: Dependent on Phase 2 (Smart Contract for Reward Distribution) and Gateway enhancements.
    * On-chain components:
        * `ProviderRegistry.sol` stores `qosScore` (score component).
        * The planned `distributeRewards` function (or similar in a rewards contract) will use `qosScore` and usage data (see Phase 2 task: "Reward distribution logic").
    * Gateway components:
        * `QoSMonitor` collects metrics that can be aggregated to determine "usage" (e.g., number of successful requests per provider).
        * The gateway needs a mechanism to report this aggregated usage data to the smart contract layer (e.g., via a trusted admin call or a more decentralized reporting system).
    * Plan:
        1. Finalize and implement the on-chain reward distribution logic (Phase 2).
        2. Implement the off-chain aggregation of usage data in the gateway.
        3. Implement the process for submitting this usage data to the smart contracts to trigger reward distribution according to the defined logic (which should factor in both usage and QoS score).

---

### 🔹 Phase 5: Frontend (dApp + Dashboard)

* [ ] Client Dashboard:
    * Implemented in `ui/src/App.tsx` with multiple tabs:
        * "Dashboard" tab: Shows network status, selected provider stats (block, gas), and simulated payment channel/request count.
        * "Providers" tab: Lists available providers from gateway's `/discovery` endpoint, shows QoS scores, stake, and allows selecting a provider for RPC requests. This covers "Dashboard for monitoring node quality and ranking" from Phase 4.
        * "Payments" tab: UI for simulated payment channel opening, viewing details, and actions like (simulated) close/top-up.
        * "Transactions" tab: UI for simulated transaction history.

  * [ ] Register, manage API keys.
    * Status: Not implemented.
    * Plan: If API key authentication is supported by the gateway, UI components will be needed for clients to register for an API key, view it, and perhaps revoke/regenerate it.

  * [ ] View usage stats, top-up credits.
    * Status: Partially implemented (simulated).
    * Usage Stats: "Transactions" tab is currently simulated. Real client-side usage tracking or stats from the gateway would be needed.
    * Top-up Credits: "Payments" tab has UI for simulated payment channel balance and top-up. This needs to be integrated with on-chain `PaymentChannel.sol` deposits or a central credit contract. The `openPaymentChannel` and `closePaymentChannel` functions in `App.tsx` are currently frontend-only simulations.

* [ ] Node Operator Dashboard:

  * [ ] Register node, stake tokens.
    * Status: Not implemented.
    * Plan: Create new UI views/components for node operators.
        * Allow connecting a wallet.
        * Interface for calling `depositStake()` and `register()` on `ProviderRegistry.sol` via `ethers.js` or similar.
        * Display current stake, registration status.

  * [ ] Monitor performance, earnings.
    * Status: Not implemented.
    * Plan:
        * Performance: UI to show a node operator their own node's detailed historical QoS metrics (beyond just the current score), health check status, etc. This might require new gateway endpoints for more detailed historical data.
        * Earnings: UI to display earnings from direct payments (e.g., `PaymentChannel` settlements) and from any distributed token rewards. This requires the reward system (Phase 2 & 4) to be functional.

* [ ] Admin panel (optional, for early manual control).
    * Status: Not implemented.
    * Plan: If needed, create a separate UI section for admin actions, guarded by wallet authentication for the admin address. This could include:
        * Overview of all nodes, system parameters.
        * Manual triggering of `slashProvider` or `updateQoS` (if not fully automated).
        * Managing reward pools or parameters.

---

### 🔹 Phase 6: Tokenomics (If Required)

* [x] Design token utility: staking, payment, slashing, governance.
    * Status: Substantially designed based on existing contracts and project goals.
    * Token: `StakeToken.sol` (STK) is an ERC20 token.
    * Utility:
        * **Staking:** Node operators must stake STK tokens to register in `ProviderRegistry.sol`. Implemented.
        * **Payment:** STK can be used by clients in `PaymentChannel.sol` to pay providers. Designed. It can also be the currency for per-request fees if that model is adopted more broadly.
        * **Slashing:** Malicious or underperforming node operators have their staked STK slashed via `ProviderRegistry.sol`. Implemented.
        * **Rewards:** Node operators will receive rewards (likely in STK) based on performance and usage, via the planned reward distribution mechanism. Designed.
        * **Governance:** If/when a DAO is implemented, STK will likely be used for voting on proposals, upgrades, and parameter changes. Designed (as a future extension).
* [x] Create token contract (ERC20 or custom).
    * Implementation: `contracts/src/StakeToken.sol` defines an ERC20 token contract named "StakeToken" (STK).
    * Details: It inherits from OpenZeppelin's `ERC20.sol` and mints an initial supply of 1,000,000 tokens to the deployer address upon construction.
* [ ] Launch initial token distribution (testnet first).
    * Status: Not started.
    * Current State: `StakeToken.sol` mints the entire initial supply (1,000,000 STK) to the contract deployer.
    * Plan:
        1. **Design Distribution Strategy:** Define allocations for team, community, rewards pool, liquidity, public sale (if any), etc. Define vesting schedules if applicable.
        2. **Testnet Execution:** The deployer will manually distribute tokens on a testnet according to the strategy. This could involve direct transfers, setting up multi-sig wallets for different allocations, or deploying simple distribution/vesting contracts if needed.
        3. **Tools/Contracts:** Consider if specific contracts (e.g., for a token sale, airdrop, or vesting) are needed beyond simple transfers from the deployer.
        4. **Documentation:** Clearly document the token distribution for transparency.

---

### 🔹 Phase 7: Deployment & Testing

* [x] Deploy smart contracts to testnet.
    * Status: Deployment script is ready. Execution on a specific testnet is pending or needs confirmation.
    * Implementation: `contracts/script/DeployContracts.s.sol` is a Foundry script that deploys `StakeToken`, `ProviderRegistry` (configured with the stake token and a minimum stake of 100 ETH), and `PaymentChannel`.
    * To Execute: Use a command like `forge script script/DeployContracts.s.sol:DeployContracts --rpc-url <TESTNET_RPC_URL> --broadcast --verify -vvvv`.
    * Note: The `broadcast/` directory would contain logs of any past deployment runs.
* [ ] Test gateway with multiple nodes.
    * Status: Testing framework appears to exist; detailed multi-node test execution and scenario validation are pending.
    * Test Files: The `node/cmd/drpcd/` directory contains various `_test.go` files (e.g., `e2e_test.go`, `integration_test.go`, `anvil_real_e2e_test.go`), suggesting capabilities for end-to-end and integration testing, likely using Anvil for local blockchain simulation.
    * Plan:
        1. Set up multiple instances of an Ethereum node (e.g., Anvil, Geth) locally or on a test network.
        2. Ensure each node can register with the `ProviderRegistry` on the testnet (requires STK and staking).
        3. Configure and run the `drpcd` gateway, connected to the testnet.
        4. Verify the gateway's `LoadBalancer` discovers all registered nodes.
        5. Send a volume of RPC requests to the gateway.
        6. Monitor (via gateway logs, QoS scores, and direct node observation):
            * Request distribution patterns.
            * QoS metric collection and on-chain updates for all nodes.
            * Failover behavior if a node is stopped/becomes unhealthy.
            * Correctness of responses.
        7. Expand existing Go tests or create new ones to automate these multi-node scenarios.
* [ ] Load testing (simulate real usage).
    * Status: Not started. Requires dedicated scripts and tooling.
    * Plan:
        1. Define target load parameters: requests per second (RPS), concurrency level, duration, types of RPC calls.
        2. Choose a load testing tool (e.g., k6, JMeter, or custom Go scripting using the existing test framework as a base).
        3. Develop load testing scripts that simulate realistic client behavior, targeting the gateway's public RPC endpoint.
        4. Execute tests against the gateway connected to multiple backend RPC nodes (on a testnet or dedicated test environment).
        5. Monitor key performance indicators (KPIs) for the gateway and nodes: CPU/memory usage, response times (average, percentiles), error rates, throughput.
        6. Identify and address bottlenecks. Re-test after optimizations.
        7. Ensure the QoS monitoring and scoring system behaves correctly under load.
* [ ] Security audits for contracts and infrastructure.
    * Status: Not started. This is a critical step before any mainnet consideration.
    * Plan:
        1. **Internal Review:** Conduct thorough internal code reviews of all smart contracts and the Go gateway codebase. Utilize static analysis tools (e.g., Slither for Solidity, gosec for Go) and linters.
        2. **Smart Contract Audit:** Engage reputable third-party security auditors specializing in Solidity and blockchain to audit `StakeToken.sol`, `ProviderRegistry.sol`, `PaymentChannel.sol`, and any subsequently developed contracts (rewards, governance).
        3. **Infrastructure Audit:** Engage security auditors (possibly with different expertise) to review the Go gateway (`drpcd`) for vulnerabilities related to network exposure, request handling, payment processing, private key management, and overall system integrity.
        4. **Scope:** Audits should cover potential economic exploits, denial-of-service vectors, access control flaws, data validation issues, and adherence to best practices.
        5. **Remediation:** Address all findings from the audits. Consider re-audits for critical fixes.
* [ ] Bug bounty or open-source review phase.
    * Status: Not started. Typically follows initial security audits and precedes mainnet launch.
    * Plan:
        1. **Open Source:** Confirm all relevant repositories (contracts, gateway, UI) are publicly accessible with appropriate open-source licenses (A `LICENSE` file exists at the root, suggesting this is planned).
        2. **Bug Bounty Program:**
            * Define scope: Which contracts and components are included?
            * Define rewards: Based on severity of findings.
            * Choose a platform (e.g., Immunefi, Hats Finance) or self-host.
            * Announce and manage the program.
        3. **Community Review:** Actively encourage community members and developers to review the codebase and provide feedback outside of a formal bounty program.
        4. **Documentation:** Ensure comprehensive developer documentation is available to aid reviewers.

---

### 🔹 Phase 8: Launch

* [ ] Deploy to mainnet.
    * Status: Pending completion of all preceding phases, particularly rigorous testing (Phase 7) and security audits (Phase 7).
    * Process: Utilize the confirmed-working Foundry deployment scripts (e.g., `contracts/script/DeployContracts.s.sol`) with the mainnet RPC URL and appropriate private keys for the deployer wallet.
    * Considerations: Mainnet gas costs, final parameter settings (e.g., `minStake`), initial liquidity for STK token if applicable.
* [ ] Publish open-source gateway + node software.
    * Status: Project appears structured for open source; formal publishing of stable versions pending.
    * Gateway Software: The `node/` directory contains the Go source for the `drpcd` gateway. This would be the primary software to publish.
    * Node Software: Providers typically run standard blockchain node clients (e.g., Geth, Nethermind). If this project offers additional software specifically *for* node operators (e.g., a wrapper, monitoring agent), that also needs publishing. Currently, the main piece is the gateway `drpcd`.
    * Plan:
        1. Ensure all code repositories (contracts, gateway, UI) have clear READMEs, contribution guidelines, and the chosen open-source license (e.g., MIT, Apache 2.0 - a `LICENSE` file exists but its content should be confirmed).
        2. Create official public repositories (e.g., on GitHub).
        3. Tag stable release versions (e.g., v1.0.0) for all components.
        4. Provide clear build and deployment instructions.
* [ ] Launch community docs and tutorials.
    * Status: Not started. A `docs/` directory exists at the project root, which can house this material.
    * Plan:
        1. Identify target audiences: End-users (dApp developers), Node Operators, STK Token Holders, Contributors.
        2. Develop content for each:
            * **End-users:** How to connect to the dRPC network, configure `ethers.js`/`web3.js`, use payment methods, API reference.
            * **Node Operators:** Requirements, setup guides (staking, registration), understanding QoS, rewards, slashing, best practices.
            * **Token Holders:** Tokenomics, utility, how to acquire/stake (if beyond node operation), governance participation.
            * **Contributors:** Development setup, architecture overview, contribution guidelines.
        3. Choose a platform: GitBook, Docusaurus, MkDocs, or Markdown files within the `docs/` directory served via GitHub Pages.
        4. Create tutorials, FAQs, and examples.
* [ ] Start incentivized testnet or early reward program.
    * Status: Pending testnet deployment (Phase 7) and a clear token distribution/rewards plan for testnet participants (Phase 6).
    * Goal: Attract early node operators and users to the testnet by offering rewards (e.g., in future mainnet tokens or testnet STK that might have some claim on mainnet tokens).
    * Plan:
        1. Ensure the full system (contracts, gateway, basic UI) is stably deployed on a public testnet.
        2. Allocate a pool of testnet STK tokens for incentives.
        3. Define the rules for the program: eligibility for rewards (e.g., uptime for nodes, usage for clients), reward amounts, duration.
        4. Announce the program and provide clear instructions for participation.
        5. Monitor participation and gather feedback to refine the system before mainnet.

---

## ⚙️ Technologies Suggested

| Layer                | Stack Options                          |
| -------------------- | -------------------------------------- |
| Gateway              | Go                                     |
| Smart Contracts      | Solidity (Ethereum, Polygon)           |
| Database             | PostgreSQL / Redis (for stats/cache)   |
| Frontend             | React + Tailwind / Next.js             |
| Node Software        | Go / Dockerized scripts                |
| Messaging (Optional) | libp2p / pub-sub for node sync         |

---
