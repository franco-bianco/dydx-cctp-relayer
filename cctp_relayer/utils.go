package cctprelayer

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	usdc_bindings "dydx-cctp-relayer/bindings/usdc"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func hexToBytes32(hexStr string) ([32]byte, error) {
	var result [32]byte

	if len(hexStr) >= 2 && hexStr[:2] == "0x" {
		hexStr = hexStr[2:]
	}

	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return result, err
	}

	if len(bytes) != 32 {
		return result, fmt.Errorf("input hex string must be exactly 32 bytes long")
	}

	copy(result[:], bytes)

	return result, nil
}

func (c *CCTPClient) checkBalanceAndAllowance() error {
	balance, err := c.getUSDCBalance()
	if err != nil {
		return fmt.Errorf("failed to get USDC balance: %w", err)
	}

	if balance.Cmp(c.transferAmount) < 0 {
		return fmt.Errorf("insufficient USDC balance | Required: %s | Available: %s", c.transferAmount, balance)
	}

	allowance, err := c.getUSDCAllowance()
	if err != nil {
		return fmt.Errorf("failed to get USDC allowance: %w", err)
	}

	if allowance.Cmp(c.transferAmount) < 0 {
		c.log.Info("insufficient USDC allowance, approving...")
		if err := c.approveUSDC(); err != nil {
			return fmt.Errorf("failed to approve USDC: %w", err)
		}
	}

	return nil
}

func (c *CCTPClient) approveUSDC() error {
	usdcABI, err := abi.JSON(strings.NewReader(usdc_bindings.BindingsABI))
	if err != nil {
		return fmt.Errorf("failed to parse USDC ABI: %w", err)
	}

	data, err := usdcABI.Pack("approve", ChainParametersMap[c.chain].CCTPRelayerContract, c.transferAmount)
	if err != nil {
		return fmt.Errorf("failed to pack approve function call: %w", err)
	}

	tx, err := c.createAndSendTransaction(ChainParametersMap[c.chain].USDCContract, data)
	if err != nil {
		return fmt.Errorf("failed to create and send transaction: %w", err)
	}

	receipt, err := bind.WaitMined(c.ctx, c.client, tx)
	if err != nil {
		return fmt.Errorf("failed to get transaction receipt: %w", err)
	}

	if receipt.Status == types.ReceiptStatusSuccessful {
		c.log.Infof("USDC approved with tx hash: %s", tx.Hash().Hex())
		return nil
	}

	return fmt.Errorf("USDC approval failed")
}

func (c *CCTPClient) createAndSendTransaction(to common.Address, data []byte) (*types.Transaction, error) {
	nonce, err := c.client.PendingNonceAt(c.ctx, crypto.PubkeyToAddress(c.privateKey.PublicKey))
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	gasLimit, err := c.client.EstimateGas(c.ctx, ethereum.CallMsg{
		From: crypto.PubkeyToAddress(c.privateKey.PublicKey),
		To:   &to,
		Data: data,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to estimate gas: %w", err)
	}
	gasLimit = uint64(float64(gasLimit) * 1.1) // Add 10% buffer

	gasPrice, err := c.client.SuggestGasPrice(c.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to suggest gas price: %w", err)
	}

	tx := types.NewTransaction(nonce, to, big.NewInt(0), gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(ChainParametersMap[c.chain].ChainID), c.privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %w", err)
	}

	err = c.client.SendTransaction(c.ctx, signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}

	return signedTx, nil
}

func (c *CCTPClient) getUSDCBalance() (*big.Int, error) {
	usdcContract, err := usdc_bindings.NewBindings(ChainParametersMap[c.chain].USDCContract, c.client)
	if err != nil {
		return nil, err
	}
	return usdcContract.BalanceOf(&bind.CallOpts{}, crypto.PubkeyToAddress(c.privateKey.PublicKey))
}

func (c *CCTPClient) getUSDCAllowance() (*big.Int, error) {
	usdcContract, err := usdc_bindings.NewBindings(ChainParametersMap[c.chain].USDCContract, c.client)
	if err != nil {
		return nil, err
	}
	return usdcContract.Allowance(&bind.CallOpts{}, crypto.PubkeyToAddress(c.privateKey.PublicKey), ChainParametersMap[c.chain].CCTPRelayerContract)
}
