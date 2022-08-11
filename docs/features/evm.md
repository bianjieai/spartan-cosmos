# EVM

## Overview

The growth of EVM-based chains (e.g. Ethereum), however, has uncovered several scalability challenges that are often referred to as the [Trilemma of decentralization, security, and scalability](https://vitalik.ca/general/2021/04/07/sharding.html). Developers are frustrated by high gas fees, slow transaction speed & throughput, and chain-specific governance that can only undergo slow change because of its wide range of deployed applications. A solution is required that eliminates these concerns for developers, who build applications within a familiar EVM environment.

The EVM module provides this EVM familiarity on a scalable, high-throughput Proof-of-Stake blockchain. It is built as a [Cosmos SDK module](https://docs.cosmos.network/master/building-modules/intro.html)which allows for the deployment of smart contracts, interaction with the EVM state machine (state transitions), and the use of EVM tooling.

_For EVM commands, refer to [EVM CLI client](../cli-client/evm.md)_

## Concepts

### EVM

The Ethereum Virtual Machine (EVM) is a computation engine which can be thought of as one single entity maintained by thousands of connected computers (nodes) running an Ethereum client. As a virtual machine ([VM ](https://en.wikipedia.org/wiki/Virtual_machine)), the EVM is responisble for computing changes to the state deterministically regardless of its environment (hardware and OS). This means that every node has to get the exact same result given an identical starting state and transaction (tx).

The EVM is considered to be the part of the Ethereum protocol that handles the deployment and execution of [smart contracts](https://ethereum.org/en/developers/docs/smart-contracts/). To make a clear distinction:

- The Ethereum protocol describes a blockchain, in which all Ethereum accounts and smart contracts live. It has only one canonical state (a data structure, which keeps all accounts) at any given block in the chain.
- The EVM, however, is the [state machine](https://en.wikipedia.org/wiki/Finite-state_machine)that defines the rules for computing a new valid state from block to block. It is an isolated runtime, which means that code running inside the EVM has no access to network, filesystem, or other processes (not external APIs).

The EVM module implements the EVM as a Cosmos SDK module. It allows users to interact with the EVM by submitting Ethereum txs and executing their containing messages on the given state to evoke a state transition.

### State

The Ethereum state is a data structure, implemented as a [Merkle Patricia Trie](https://en.wikipedia.org/wiki/Merkle_tree), that keeps all accounts on the chain. The EVM makes changes to this data structure resulting in a new state with a different State Root. Ethereum can therefore be seen as a state chain that transitions from one state to another by executing transations in a block using the EVM. A new block of txs can be described through its Block header (parent hash, block number, time stamp, nonce, receipts,...).

### Accounts

There are two types of accounts that can be stored in state at a given address:

- **Externally Owned Account (EOA)**: Has nonce (tx counter) and balance
- **Smart Contract**: Has nonce, balance, (immutable) code hash, storage root (another Merkle Patricia Trie)

Smart contracts are just like regular accounts on the blockchain, which additionally store executable code in an Ethereum-specific binary format, known as **EVM bytecode**. They are typically written in an Ethereum high level language such as Solidity which is compiled down to EVM bytecode and deployed on the blockchain, by submitting a tx using an Ethereum client.

## Functions

### API-related features

Spartan-Cosmos supports all features of `EVM`. For the related `API` documentation, you can refer to: [EVM API](https://eth.wiki/json-rpc/API)

### Export the `ETH` private key of the account

```shell
spartan keys unsafe-export-eth-key [name]
```

### Import the account's `ETH` private key

```shell
spartan keys unsafe-import-eth-key [name] [pk]
```

### spartan tx evm raw

Build cosmos transaction from raw ethereum transaction

```bash
spartan tx evm raw [tx-hex] [flags]
```

### spartan query evm code

Gets code from an account.

```bash
spartan query evm code [tx-hex] [flags]
```

### spartan query evm storage

Gets storage for an account with a given key and height.

```bash
spartan query evm storage [address] [key] [flags]
```
