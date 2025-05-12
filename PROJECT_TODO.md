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
* [x] Reward distribution logic (based on usage or time).
    * Status: Implemented (Pull-based system).
    * Implementation: `ProviderRegistry.sol` now includes:
        * `rewardPoolBalance`: Tracks available STK for rewards.
        * `providerRewards`: Mapping `address => uint256` storing claimable rewards.
        * `fundRewardPool(amount)`: Owner function to add STK to the pool.
        * `calculateAndAllocateRewards(providers[], totalAmount)`: Owner function to calculate rewards (based on `stake * qosScore`) for a given list of providers and allocate the total amount proportionally to their `providerRewards` balance.
        * `claimRewards()`: Provider function to withdraw their `providerRewards` balance.
    * Notes: This design requires an off-chain process (run by the owner/DAO) to:
        1. Determine which providers are eligible for rewards in a period.
        2. Calculate the total reward amount for that period.
        3. Call `calculateAndAllocateRewards` with the list and total amount.
        4. Providers then call `claimRewards` themselves.
* [x] Slashing mechanism for malicious/incorrect nodes.
    * Implementation: `ProviderRegistry.sol` includes an `Ownable` function `slashProvider(address provider_, uint256 amount)`. Slashed tokens are transferred to the contract owner.
    * Notes: The effectiveness of this mechanism depends on a reliable off-chain (or on-chain via verification) process for detecting and reporting malicious behavior to the owner/DAO.
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
    * Implementation: The Go package `node/pkg/loadbalancer/` contains `loadbalancer.go`. The main gateway handler `handleRPCRequest` in `node/cmd/drpcd/main.go` uses this load balancer's `GetProvider()` method to select an upstream node.
    * Logic: `LoadBalancer` fetches provider data (address, QoS score, stake). `GetProvider()` uses weighted random selection based on `QoSScore^2`.
            * Current Status: The load balancer (`node/pkg/loadbalancer/`) fetches provider data (including endpoint URLs) via direct calls to the `ProviderRegistry` contract during its refresh cycle. The Go bindings are up-to-date. Event-based updates could be a future optimization.
  * [x] Rate limiting, abuse prevention.
    * Status: Basic implementation complete, refinements pending.
    * Implementation:
        * **Rate Limiting:** `checkRateLimit` function in `main.go` uses Redis (if `REDIS_ADDR` is set) with a fixed window counter (`ratelimit:<client_address>`) per authenticated client address (`X-DRPC-Signature`). Configurable via `rateLimitRequests` and `rateLimitDuration` constants.
        * **Abuse Prevention:** Basic protection via client authentication (`verifyClientSignature`), request body size limits (`http.MaxBytesReader`), and timestamp validation in signatures. More sophisticated pattern detection is not implemented.
    * Plan for Refinement:
        * Consider implementing sliding window rate limiting in Redis (more complex) for better accuracy.
        * Define and implement specific malicious pattern detection if needed (e.g., blocking overly complex `eth_call` requests).
        * Implement higher rate limits based on client stake (requires adding stake checking logic, potentially complex).
        * Clarify how gateway-detected abuse (if any specific types are defined) would trigger on-chain slashing.
  * [x] Logging & billing integration.
    * Status: Substantially implemented.
    * Logging:
        * The `node/pkg/qos/monitor.go` (`QoSMonitor`) records performance metrics.
        * Enhanced structured logging implemented in `loggingMiddleware` in `cmd/drpcd/main.go` (includes method, URI, status, duration, etc.).
    * Billing Integration (Usage Data for Rewards):
        * The gateway (`cmd/drpcd/main.go`) now increments a usage counter in Redis (`usage:<provider_address>`) for each successful request handled by a provider (requires `REDIS_ADDR` env var).
        * This provides the necessary data source for the off-chain process that calculates rewards and calls `calculateAndAllocateRewards`.
        * Integration for `PaymentChannel.sol` (if used alongside or instead of the reward pool) remains less defined.
    * Plan:
        * Implement the external off-chain service/script to read Redis usage data, calculate rewards, and call `calculateAndAllocateRewards`.
        * Clarify/implement gateway interaction for `PaymentChannel.sol` if needed.
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
  * [x] Ping/health check system.
    * Status: Implemented.
    * Implementation: The `startHealthChecker` function in `node/cmd/drpcd/main.go` runs periodically:
        1. Fetches all providers from the `LoadBalancer`.
        2. For each provider with an `EndpointURL`, calls `checkProviderHealth`.
        3. `checkProviderHealth` sends an `eth_blockNumber` request to the provider's endpoint.
        4. The success/failure and latency are recorded using `QoSMonitor.RecordMetric()`.
        5. The `QoSMonitor` uses these metrics to calculate scores, which are periodically updated on-chain (using a configurable `CHAIN_ID`).
    * Notes: Consistent failures will naturally lower the QoS score used for load balancing and potentially rewards.
  * [x] Track node performance & uptime.
    * Implementation: The `node/pkg/qos/monitor.go` (`QoSMonitor`) is responsible for this.
    * Performance: It records response times and success/failure for each request (or health check) to a provider. This data is used to calculate a `QoSScore`.
    * Uptime: The success rate, a key component of the `QoSScore`, implicitly tracks uptime. Consistent failures will lower the score.
    * On-chain Record: `QoSMonitor` periodically updates these `QoSScore` values on the `ProviderRegistry` smart contract.
    * Notes: The `Provider` struct in `loadbalancer.go` has an `UpSince` field; its usage and update mechanism should be clarified or enhanced for more explicit uptime tracking if needed beyond the QoS score.
* [x] API for clients to connect (public RPC endpoint).
    * Implementation: The `node/cmd/drpcd/main.go` sets up an HTTP server with a `/` endpoint (`handleRPCRequest`) to receive JSON-RPC requests.
    * Current Status & Key Issues:
        * **Proxying Logic:** The `handleRPCRequest` function *does* implement logic to use `s.loadBalancer.GetProvider()` to select a provider, proxy the request to the provider's `EndpointURL`, and record QoS metrics against the selected provider's address using `s.qosMonitor.RecordMetric()`. The previous assessment that it used a single hardcoded URL seems outdated based on the current `main.go` code. **However, this relies on the LoadBalancer providing correct endpoint URLs, which is likely the next integration point.**
        * **Client Authentication:** Header-based client signature authentication (`verifyClientSignature`) and rate limiting (`checkRateLimit`) *are* implemented in `handleRPCRequest`.
        * **Payment Channel Interaction:** The `validatePayment` function (using payment details embedded in the RPC request) still exists but its integration with the primary header-based auth and the on-chain `PaymentChannel.sol` needs clarification or removal if superseded.
        * **Load Balancer Data:** The core issue likely lies in the `LoadBalancer` needing to fetch live provider data, *including endpoint URLs*, from the `ProviderRegistry` contract.
    * Positive Aspects:
        * Basic HTTP server structure is in place using `gorilla/mux`.
        * Includes logging, CORS, client signature auth, and rate limiting middleware/checks.
        * `/discovery` endpoint provides provider list (needs live data).
        * `/health` endpoint for gateway health.
        * Separate `startHealthChecker` function and related code exists, addressing the "Ping/health check system" task partially.
    * Plan:
        1.  **(Done)** Enhance the `LoadBalancer` to fetch live provider data including `endpointURL` (Go bindings updated, fetch logic confirmed).
        2.  **(Done)** Implement the Smart Contract reward distribution logic (Phase 2 - pull-based system added).
        3.  **(Done)** Refine/complete the "Ping/health check system" integration.
        4.  Clarify/Refactor the payment validation logic (`validatePayment` vs. header auth).
        5.  Implement remaining Phase 3 tasks (Rate limiting / abuse prevention refinement, logging/billing integration refinement, optional verification).
        6.  **(In Progress)** Implement the off-chain components for reward calculation/allocation (`rewardallocator` script created, needs refinement - see Phase 4 task).
        7.  Implement remaining Phase 3 tasks (Abuse prevention details, optional verification).
        8.  Proceed with other incomplete phases (Frontend, Testing, etc.).

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
    * Status: On-chain component (allocation/claiming) implemented. Off-chain script for calculation/triggering created, needs refinement.
    * On-chain components:
        * `ProviderRegistry.sol` stores `qosScore` (score component) and `stake`.
        * The `calculateAndAllocateRewards` function uses `stake * qosScore` for proportional allocation. Usage data is *not* directly used on-chain in this implementation, but the *selection* of providers and the *total reward amount* passed to the function can be based on off-chain usage data.
        * Providers use `claimRewards()` to pull their allocated rewards.
    * Gateway components:
        * `QoSMonitor` collects metrics (`qosScore`) and could potentially aggregate usage counts.
        * The gateway (or a separate backend service) needs a mechanism to:
            1. Aggregate usage data per provider over a reward period (Data source: Redis counters incremented by the gateway, e.g., `usage:<provider_address>`).
            2. Determine the list of eligible providers (`providersToReward`).
            3. Calculate the `totalRewardAmount` for the period (based on usage, tokenomics, etc.).
            4. Call the `calculateAndAllocateRewards` function on the `ProviderRegistry` contract with this data.
    * Implementation (Off-chain):
        * A Go script `node/cmd/rewardallocator/main.go` has been created.
        * It connects to Redis and Ethereum, loads owner key, registry address etc. from env vars.
        * Fetches provider usage from Redis keys (`usage:<provider_address>`).
        * **Placeholder:** Uses a hardcoded provider list (needs replacement with event querying or other discovery method).
        * Filters providers based on registration status and minimum usage (`MIN_USAGE_COUNT` env var).
        * Calls `calculateAndAllocateRewards` with the eligible list and a total reward amount (`TOTAL_REWARD_AMOUNT` env var).
        * Optionally resets Redis counters (`RESET_COUNTERS` env var).
    * Plan:
        1. Implement robust provider discovery in `rewardallocator` (e.g., querying `ProviderRegistered` events).
        2. Refine error handling and add scheduling (e.g., cron) for `rewardallocator`.
        3. Enhance the UI to show available rewards and allow providers to trigger `claimRewards`.

---

### 🔹 Phase 5: Frontend (dApp + Dashboard)

* [x] Client Dashboard:
    * Implemented in `ui/src/App.tsx` with multiple tabs:
        * "Dashboard" tab: Shows network status (active providers), selected provider stats (block, gas, network via gateway RPC), and simulated payment channel/request count.
        * "Providers" tab: Lists available providers from gateway's `/discovery` endpoint (address, QoS score, stake), allows selecting a provider.
        * "Payments" tab: UI for *simulated* payment channel operations.
        * "Transactions" tab: UI for *simulated* transaction history.
    * Wallet Connection: Implemented via `window.ethereum` (MetaMask), displays connected account and network.

  * [ ] Register, manage API keys.
    * Status: Not implemented.
    * Plan: If API key authentication is supported by the gateway, UI components will be needed for clients to register for an API key, view it, and perhaps revoke/regenerate it.

  * [ ] View usage stats, top-up credits.
    * Status: Partially implemented (simulated for payments/transactions).
    * Usage Stats: "Transactions" tab is currently simulated.
    * Top-up Credits: "Payments" tab is simulated. Real integration with `PaymentChannel.sol` or a credit system would be needed.

* [x] Node Operator Dashboard:
  * Implemented in `ui/src/App.tsx` under "Operator Zone" tab (requires wallet connection):
    * Displays operator's data from `ProviderRegistry` (stake, QoS, registration status, endpoint URL, registration/deregistration times).
    * Displays claimable rewards from `ProviderRegistry` and allows claiming.
    * Allows depositing STK tokens (approve + depositStake).
    * Allows setting/updating their endpoint URL.
    * Allows registering and deregistering the node.
    * Allows withdrawing stake (after deregistration and cooldown - cooldown logic not fully enforced by UI yet, relies on contract).
    * Uses `ethers.js` for all contract interactions.
    * Basic transaction status reporting (processing, success, error).

  * [ ] Register node, stake tokens.
    * Status: Implemented as part of Node Operator Dashboard.

  * [ ] Monitor performance, earnings.
    * Status: Partially implemented.
    * Performance: QoS score is displayed. More detailed historical metrics are not yet shown.
    * Earnings: Claimable rewards from the `ProviderRegistry` reward pool are displayed and can be claimed. Direct payment channel earnings are not tracked on this dashboard.
    * Plan:
        * Consider UI for more detailed historical QoS or gateway-reported stats (might need new gateway endpoints).

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
