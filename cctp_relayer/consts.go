package cctprelayer

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Chain string

const (
	Arbitrum  Chain = "arbitrum"
	Avalanche Chain = "avalanche"
	Optimism  Chain = "optimism"

	DestinationDomain = 4                                                                    // Noble Domain
	DestinationCaller = "0x000000000000000000000000691cf4641d5608f085b2c1921172120bb603d074" // Noble CCTP Relayer
)

type ChainParameters struct {
	RPCURL              string
	USDCContract        common.Address
	CCTPRelayerContract common.Address
	ChainID             *big.Int
	FeeAmount           *big.Int
	DestinationCaller   string // TODO: convert to common.Hash
}

var ChainParametersMap = map[Chain]ChainParameters{
	Arbitrum: {
		RPCURL:              "https://arbitrum.llamarpc.com",
		USDCContract:        common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831"),
		CCTPRelayerContract: common.HexToAddress("0xe298b93ffB5eA1FB628e0C0D55A43aeaC268e347"),
		ChainID:             big.NewInt(42161),
		FeeAmount:           big.NewInt(20000),
		DestinationCaller:   DestinationCaller,
	},
	Avalanche: {
		RPCURL:              "https://avalanche-c-chain-rpc.publicnode.com",
		USDCContract:        common.HexToAddress("0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E"),
		CCTPRelayerContract: common.HexToAddress("0x1e457c5c024707f4192Fa4B4471679882c127F60"),
		ChainID:             big.NewInt(43114),
		FeeAmount:           big.NewInt(20000),
		DestinationCaller:   DestinationCaller,
	},
	Optimism: {
		RPCURL:              "https://optimism.llamarpc.com",
		USDCContract:        common.HexToAddress("0x0b2c639c533813f4aa9d7837caf62653d097ff85"),
		CCTPRelayerContract: common.HexToAddress("0x48C5417ED570928eC85D5e3AD4e7E0EeD7dB1E2A"),
		ChainID:             big.NewInt(10),
		FeeAmount:           big.NewInt(20000),
		DestinationCaller:   DestinationCaller,
	},
}
