# CCTP Relayer for dYdX Deposits

**CCTP Relayer** is a Go package that demonstrates how to deposit USDC to dYdX from other chains using Circle's Cross-Chain Transfer Protocol (CCTP). This example specifically shows how to deposit USDC from x-chain to dYdX. For more information on how it works, please refer [here](https://dydx.exchange/blog/cctp).

## Features

- Initiate cross-chain USDC transfers to dYdX
- Supports Arbitrum, Optimism and Avalanche
- Token balance and allowance checks

## Installation

To use this package in your Go project, run:

```bash
github.com/franco-bianco/dydx-cctp-relayer
```

## Usage

1. Create a new CCTP client:

   ```go
   client := relayer.NewClient()
   ```

2. Request a CCTP transfer:

   ```go
   err := client.RequestTransfer(...)
   ```

## Configuration

The package uses environment variables for configuration. Create a `.env` file in your project root with the following variables:

- `PRIVATE_KEY`
- `NOBLE_MINT_RECIPIENT`

### Noble Mint Recipient Address

Currently not able to obtain the Noble Mint Recipieint dynamically from the dYdX on-chain address, so an initial manual deposit will have to be done first to obtain it. The **NOBLE_MINT_RECIPIENT** address is unique to your dYdX account and is crucial for depositing to dYdX. To obtain this address:

1. Make an initial manual deposit to your dYdX address from the dYdX frontend UI
2. The mint recipient address can be found in the input data of the transaction

**Note:** The mint recipient address appears to be consistent across all chains for a particular dYdX account. Once you've identified it, you can use the same address for deposits from different chains.

## Supported Chains

Currently, the package supports:

- Arbitrum
- Avalanche
- Optimism

More chains can be added by extending the `Chain` type and `ChainParametersMap` in the package.

## Disclaimer

This package is provided as-is and serves as an example of how to deposit USDC to dYdX using the CCTP relayer. Always test thoroughly before using in a production environment. The authors are not responsible for any losses incurred through the use of this software.
