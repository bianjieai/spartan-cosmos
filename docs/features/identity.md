# Identity

## Overview

The Identity module builds a decentralized identity system (`DID`) that implements and extends the [W3C DID specification](https://www.w3.org/TR/did-core/).

Key features include.

- The DID method is `irita`, and the full identity DID form is represented as: did:irita:id

- The cryptographic material of the identity includes the public key as well as the public key certificate

- The identity can contain additional credential information

An identity consists of the following components.

- _ID_: a globally unique identity identifier

- _public key list_: list of public keys for the identity subject

- _public-key-certificate list_: list of public-key certificates for the identity subject

- _identity credential URI_: URI of the identity subject's off-chain credential information

_For Identity commands, refer to [Identity CLI client](../cli-client/identity.md)_

## Functions

### Create

An identity is created by providing the identity ID, public key, certificate, and credential URI. All the above parameters are optional. If the identity ID is not specified, it will be generated automatically.

Currently supported public key algorithms include.

- `RSA`: `DER` encoded public key
- `DSA`: `DER` encoded public key
- `ECDSA`: 33-byte compressed public key
- `ED25519`: 32-byte compressed public key
- `SM2`: 33-byte compressed public key

All public keys are represented by the `Hex` string.
