# Feemarket

## Overview
This document specifies the feemarket module which allows to define a global transaction fee for the network.

This module has been designed to support EIP1559 in cosmos-sdk.

The MempoolFeeDecorator in x/auth module needs to be overrided to check the baseFee along with the minimal-gas-prices allowing to implement a global fee mechanism which vary depending on the network activity.

For more reference to EIP1559:

https://github.com/ethereum/EIPs/blob/master/EIPS/eip-1559.md

## Concepts

EIP-1559 (opens new window)describes a pricing mechanism that was proposed on Ethereum to improve to calculation of transaction fees. It includes a fixed-per-block network fee that is burned and dynamically expands/contracts block sizes to deal with peaks of network congestion.

Before EIP-1559 the transaction fee is calculated with

```
fee = gasPrice * gasLimit
```

, where gasPrice is the price per gas and gasLimit describes the amount of gas required to perform the transaction. The more complex operations a transaction requires, the higher the gasLimit (See Executing EVM bytecode (opens new window)). To submit a transaction, the signer needs to specify the gasPrice.


With EIP-1559 enabled, the transaction fee is calculated with

```
fee = (baseFee + priorityTip) * gasLimit
```

, where baseFee is the fixed-per-block network fee per gas and priorityTip is an additional fee per gas that can be set optionally. Note, that both the base fee and the priority tip are a gas prices. To submit a transaction with EIP-1559, the signer needs to specify the gasFeeCap a maximum fee per gas they are willing to pay total and optionally the priorityTip , which covers both the priority fee and the block's network fee per gas (aka: base fee).