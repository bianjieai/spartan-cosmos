# NFT

## Overview

The NFT module provides the ability to digitize assets. Through this module, each off-chain asset will be modeled as a unique on-chain nft.

_For NFT commands, refer to [NFT CLI client](../cli-client/nft.md)_

## Concepts

nft on the chain are identified by `ID`. With the help of the secure and non-tamperable features of the blockchain, the ownership of nft will be clarified. The transaction process of nft among members will also be publicly recorded to facilitate traceability and dispute settlement.

nft metadata (`metadata`) can be stored directly on the chain, or the `URI` of its storage source outside the chain can be stored on the chain. nft metadata is organized according to a specific [JSON Schema](https://JSON-Schema.org/).

nft need to be issued before creation to declare their abstract properties:

- _Denom_: the globally unique nft classification name

- _Denom ID_: the globally unique nft classification identifier of Denom

- _Symbol_: the symbol of the token

- _Mint-restricted_: This field indicates whether there are restrictions on the issuance of NFTs under this classification, true means that only Denom owners can issue NFTs under this classification, false means anyone can

- _Update-restricted_: This field indicates whether there are restrictions on updating NFTs under this classification, true means that no one under this classification can update the NFT, false means that only the owner of this NFT can update

- _Metadata Specification_: The JSON Schema that nft metadata should follow

Each specific nft is described by the following elements:

- _Denom_: the classification of the nft

- _ID_: The identifier of the nft, which is unique in this nft denom; this ID is generated off-chain

- _Metadata_: The structure containing the specific data of the nft

- _Metadata URI_: When metadata is stored off-chain, this URI indicates its storage location
