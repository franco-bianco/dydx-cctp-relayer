package cctprelayer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	cctp_bindings "dydx-cctp-relayer/bindings/cctp_relayer"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

type CCTPClient struct {
	ctx                context.Context
	log                *logrus.Entry
	chain              Chain
	privateKey         *ecdsa.PrivateKey
	client             *ethclient.Client
	nobleMintRecipient string
	transferAmount     *big.Int
}

func NewCCTPClient(ctx context.Context, log *logrus.Logger, hexPrivateKey string, chain Chain, mintRecipient string, transferAmountFloat float64) (*CCTPClient, error) {
	privateKey, err := crypto.HexToECDSA(hexPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	client, err := ethclient.Dial(ChainParametersMap[chain].RPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the client: %w", err)
	}

	transferAmount := new(big.Int).SetInt64(int64(transferAmountFloat * 1e6))

	return &CCTPClient{
		ctx:                ctx,
		log:                log.WithField("component", "cctp_client"),
		chain:              chain,
		privateKey:         privateKey,
		client:             client,
		nobleMintRecipient: mintRecipient,
		transferAmount:     transferAmount,
	}, nil
}

func (c *CCTPClient) RequestCCTPTransfer() error {
	c.log.Info("initiating CCTP transfer request...")

	if err := c.checkBalanceAndAllowance(); err != nil {
		return err
	}

	relayerABI, err := abi.JSON(strings.NewReader(cctp_bindings.BindingsABI))
	if err != nil {
		return fmt.Errorf("failed to parse CCTPRelayer ABI: %w", err)
	}

	mintRecipientBytes, err := hexToBytes32(c.nobleMintRecipient)
	if err != nil {
		return fmt.Errorf("failed to convert noble mint recipient to bytes32: %w", err)
	}

	destinationCaller, err := hexToBytes32(ChainParametersMap[c.chain].DestinationCaller)
	if err != nil {
		return fmt.Errorf("failed to convert destination caller to bytes32: %w", err)
	}

	data, err := relayerABI.Pack("requestCCTPTransferWithCaller",
		c.transferAmount,
		uint32(DestinationDomain),
		mintRecipientBytes,
		ChainParametersMap[c.chain].USDCContract,
		ChainParametersMap[c.chain].FeeAmount,
		destinationCaller,
	)
	if err != nil {
		return fmt.Errorf("failed to pack requestCCTPTransferWithCaller function call: %w", err)
	}

	tx, err := c.createAndSendTransaction(ChainParametersMap[c.chain].CCTPRelayerContract, data)
	if err != nil {
		return fmt.Errorf("failed to create and send transaction: %w", err)
	}

	receipt, err := bind.WaitMined(c.ctx, c.client, tx)
	if err != nil {
		return fmt.Errorf("failed to get transaction receipt: %w", err)
	}

	if receipt.Status == types.ReceiptStatusSuccessful {
		c.log.Infof("CCTP transfer successful | Tx: %s | Gas Used: %d | Block Number: %d", tx.Hash().String(), receipt.GasUsed, receipt.BlockNumber)
		return nil
	}

	return fmt.Errorf("CCTP transfer failed")
}
