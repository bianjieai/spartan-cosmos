---
order: 4
---

# Join the Mainnet

:::tip WARNING
Since the Miannet is not prepared, this document provides a limited reference.
:::

## Select a Mainnet
TODO

## Install Spartan
Follow the [Installation](installation.md) document to install the Spartan binary.

## Initialize Node
```shell
# initialize node configurations
spartan testnet -v 1 --chain-id=<spartan-mainnet>
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

## Run a Mainnet Validator
TODO

## Start the Mainnet
```shell
spartan start 
```