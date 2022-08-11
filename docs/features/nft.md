# NFT

## Overview

The NFT module provides the ability to digitize assets. Through this module, each off-chain asset will be modeled as a unique on-chain nft.

_For NFT commands, refer to [NFT CLI client](../cli-client/nft.md)_

## Concepts

nft on the chain are identified by `ID`. With the help of the secure and non-tamperable features of the blockchain, the ownership of nft will be clarified. The transaction process of nft among members will also be publicly recorded to facilitate traceability and dispute settlement.

nft metadata (`metadata`) can be stored directly on the chain, or the `URI` of its storage source outside the chain can be stored on the chain. nft metadata is organized according to a specific [JSON Schema](https://JSON-Schema.org/).

nft need to be issued before creation to declare their abstract properties:

- **Denom:** the globally unique nft classification name

- **Denom ID:** the globally unique nft classification identifier of Denom

- **Symbol:** the symbol of the token

- **Mint-restricted:** This field indicates whether there are restrictions on the issuance of NFTs under this classification, true means that only Denom owners can issue NFTs under this classification, false means anyone can

- **Update-restricted:** This field indicates whether there are restrictions on updating NFTs under this classification, true means that no one under this classification can update the NFT, false means that only the owner of this NFT can update

- **Metadata Specification:** The JSON Schema that nft metadata should follow

Each specific nft is described by the following elements:

- **Denom:** the classification of the nft

- **ID:** The identifier of the nft, which is unique in this nft denom; this ID is generated off-chain

- **Metadata:** The structure containing the specific data of the nft

- **Metadata URI:** When metadata is stored off-chain, this URI indicates its storage location
