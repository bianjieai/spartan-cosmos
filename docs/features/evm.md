# EVM

## Overview

Smart Contract, a computer protocol designed to disseminate, validate, or enforce contracts in an informational manner, is a customizable extension to the functionality of the blockchain. The main popular smart contract implementation method is Ether's `EVM`. spartan has support for `EVM`.

Key features include

- The default `chain-id` for `EVM` is `1223`; to change this port, modify `github.com/tharsis/ethermint/types.EvmChainID=<your_chain_id>` in the makefile before compiling
- The relevant `API` ports for `EVM` are: `8545` and `8546`
- The default is to enable `EVM` related functions. The `namespace` that are enabled by default are: `"eth,net,web3"`
- Compatible with web3 related components. For example: `metamask` and `Remix` and other related development components
- For other configuration items, please refer to `EVM Configuration` in `app.toml`.

### Caution

The `EVM` module only supports accounts generated with the `eth_secp256k1` algorithm. Accounts generated with the `eth_secp256k1` algorithm are generated as follows

```shell
spartan keys add [account_name] --algo eth_secp256k1
```

## Functions

### API-related features

spartan supports all features of ``EVM``. For the related `API` documentation, you can refer to: [EVM API](https://eth.wiki/json-rpc/API)

### Export the `ETH` private key of the account

```shell
spartan keys unsafe-export-eth-key [name]
```

### Import the account's `ETH` private key

```shell
spartan keys unsafe-import-eth-key [name] [pk]
```

### Get the `code` of the smart contract

Allows the user to query the smart contract code at the given address.

```shell
spartan query evm code [address] [flags]
```

### Get the `code` of the smart contract

Allows the user to query the `address` store with the given `key` and `height`.

```shell
spartan query evm storage [address] [key] [flags]
```

### Sending transactions

Allows users to construct `Cosmos` transactions from raw `ETH` transactions.

```shell
spartan tx evm raw [tx-hex] [flags]
```