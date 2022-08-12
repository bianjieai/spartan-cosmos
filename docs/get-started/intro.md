---
order: 1
---

# Introduction

## Spartan-Cosmos

The **Spartan-Cosmos** chain is a Non-Crypto Public Chain based on **Cosmos-SDK**, **IRISnet**, and **IRITA**, which forms **the BSN Spartan Network**, together with other blockchains conforming to the BSN Non-Crypto Public Chain Transformation Instructions.

### Node Management

In Spartan Network exist two types of nodes: Accounting Node and Consensus Node. An Accounting Node can join and exit the network freely, while a Consensus Node is only allowed to join if voted on by the Spartan Governance System. 

Spartan-Cosmos achieves this feature by introducing [the Gov module](../features/gov.md). A node will be selected as a Consensus Node, also known as a Validator in Cosmos-SDK based chain, only after it passes the voting process of an on-chain proposal for creating validators.

### Resource Management

Non-Crypto Public Chains need to support the billing of computing and storage resources consumed by transaction processing on the chain based on a [Gas](../concepts/basics.md#gas) metering mechanism. Users can pay for Gas through [Energy Credits](../concepts/basics.md#energy-credit).

Spartan-Cosmos integrates [Cosmos GasMeter](https://docs.cosmos.network/main/basics/gas-fees.html#gas-meter) to manage the consumption of resources during execution. In the Cosmos-SDK based chain, Gas is a special unit used to track resource consumption during execution. Gas is typically consumed whenever read and writes are made to the store, but it can also be consumed if expensive computation needs to be done. 

### Cryptocurrency Transfer Control Management

Unlike Crypto Public Chains, Energy Credit cannot be traded or transferred between Chain Accounts Non-Crypto Public Chains must disable peer-to-peer Energy Credit transfers between Chain Accounts (both wallet address or contract address) and only support recharging from the corresponding data center’s NTT account, which are used to pay for Gas. The spent Energy Credits will be sent to Admin Accounts and deleted.

Spartan-Cosmos uses [the OPB module](../features/opb.md) to place restrictions on account transactions and transfers. 

:::tip
Refer to **[Features](../features/overview.md)** to discover more about Spartan-Cosmos.  
:::