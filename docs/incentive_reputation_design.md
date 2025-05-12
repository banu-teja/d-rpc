# Incentive and Reputation Layer Design

This document details the design of the incentive and reputation mechanisms within the d-rpc network. This includes how providers are rewarded, how malicious behavior is penalized (slashing), and how provider quality is measured and used.

## 1. Reward Mechanisms

- Detail the per-request payment mechanism via payment channels.
- Outline any future plans for additional reward distribution based on aggregated usage or QoS.

## 2. Slashing Mechanism

- Describe the on-chain slashing function in `ProviderRegistry.sol`.
- **Define Slashing Criteria:** What specific actions or inactions by a provider will lead to slashing? This section should detail the types of misbehavior and the criteria for triggering a slashing event.\n\n    - **Serving Incorrect Data:**\n        - How is incorrect data defined and detected? (Dependent on decentralized verification research)\n        - What is the threshold for incorrect responses that triggers slashing? (e.g., percentage of incorrect responses over a period)\n        - Are there different severity levels for incorrect data?\n\n    - **Downtime/Unavailability:**\n        - What duration or frequency of downtime triggers slashing?\n        - How is downtime measured and confirmed? (Related to ping/health check system)\n\n    - **Malicious Attacks (e.g., Sybil Attacks, DDoS):**\n        - How are these attacks detected?\n        - What are the slashing conditions in such cases?\n\n    - **Failure to Participate:**\n        - Are there requirements for providers to participate in certain activities (e.g., signing messages, contributing to verification)?\n        - What are the penalties for non-participation?\n\n    - **Other Misbehavior:**\n        - Are there any other actions or inactions that should lead to slashing?\n\n- How will malicious/incorrect behavior be detected off-chain?
- How will the slashing transaction be triggered?

## 3. Reputation and QoS Scoring

- Detail the current QoS scoring system based on latency and response validity (`qos/monitor.go`).
- How will uptime be incorporated into the scoring?
- How will the QoS score impact provider selection (Load Balancer) and potential future reward distribution?

## 4. Integration with Gateway and Smart Contracts

- How do the gateway components (QoS monitor, potential detection logic) interact with the smart contracts for staking, slashing, and QoS updates?

---

**Instructions for Designer:** Please fill in the details under each section based on the decided design. Add new sections or points as needed.
