---
order: 5
---

# Join the Testnet

:::tip WARNING
Since the Testnet is not prepared, this document provides a limited reference.
:::

## Pick a Testnet
TODO

## Install Spartan
Follow the [Installation](installation.md) document to install the Spartan binary.

## Initialize Node
```shell
# initialize node configurations
spartan init <moniker> --chain-id=<spartan-mainnet>
```

## Genesis & Seeds

In this step you should download public `genesis.json` from official site.
```shell
# download mainnet public config.toml and genesis.json
curl -o ~/.spartan/config/genesis.json <official-genesis-url>
```

To find peers, it's necessary to add healthy seed nodes to `config.toml`. As well, you can download it from official site.
```shell
curl -o ~/.spartan/config/config.toml <official-config-url>
```

## Run a Testnet Validator
TODO

## Start the Testnet
```shell
spartan start 
```
