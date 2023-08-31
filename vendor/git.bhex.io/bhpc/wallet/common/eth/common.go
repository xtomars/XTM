/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: common.go
 * @Date: 2018/06/04
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package eth

import (
	"context"
	"errors"
	"math"
	"math/big"

	"git.bhex.io/bhpc/wallet/common/eth/factory"
	"git.bhex.io/bhpc/wallet/common/logger"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

const (
	// DefaultGasLimit for the default gas limit when build a transaction
	DefaultGasLimit = 470000
	// DefaultGasPrice for the default gas price (wei) when build a transaction
	DefaultGasPrice = 10
	// DefaultDecimals is ETH default decimals
	DefaultDecimals = 18
)

var log = logger.New("eth")

// ToWei convert actual amount to wei amount
func ToWei(amount string, decimals int) (*big.Int, error) {
	aa, err := decimal.NewFromString(amount)
	if err != nil {
		return nil, err
	}
	bb := decimal.NewFromFloat(math.Pow10(decimals))
	aa = aa.Mul(bb)

	weiAmount, _ := big.NewInt(0).SetString(aa.String(), 0)

	return weiAmount, nil
}

// ToEther convert  wei amount to actual amount
func ToEther(i *big.Int, decimals int) *big.Float {
	return new(big.Float).Mul(big.NewFloat(math.Pow10(decimals*-1)), big.NewFloat(0).SetInt(i))
}

func broadcastTransaction(client *ethclient.Client, tx *types.Transaction) (txHash string, gasUsed uint64, err error) {
	// broadcast transaction to eth network
	err = client.SendTransaction(context.TODO(), tx)
	if err != nil {
		log.Errorf("BroadcastTransaction tx err:%v", err)
		return
	}
	log.Debugf("BroadcastTransaction tx: 0x%x\n", tx.Hash())

	// wait for the transaction be mined
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Errorf("WaitMined tx err:%v", err)
		return
	}
	txHash, gasUsed = receipt.TxHash.String(), receipt.GasUsed
	log.Debugf("BroadcastTransaction status %v, TxHash %v, Gas used %v, costFee: %v", receipt.Status, txHash, receipt.GasUsed, tx.Cost())
	return
}

func buildWalletInstance(client *ethclient.Client, walletAddress, priKey string, nonce, gasPrice uint64) (*factory.MultiSigWallet, *bind.TransactOpts, error) {
	// convert private string to ecdsa privatekey
	ecdsaPri, err := crypto.HexToECDSA(priKey)
	if err != nil {
		log.Errorf("Parse private key err:%v", err)
		return nil, nil, err
	}
	opts := bind.NewKeyedTransactor(ecdsaPri)
	opts.GasLimit = DefaultGasLimit
	opts.GasPrice = big.NewInt(int64(gasPrice))
	opts.Nonce = big.NewInt(int64(nonce))

	// get wallet instance
	walletInstance, err := factory.NewMultiSigWallet(common.HexToAddress(walletAddress), client)
	if err != nil {
		log.Errorf("Get wallet instance err:%v", err)
		return nil, nil, err
	}

	return walletInstance, opts, nil
}

func blockHeight(client *ethclient.Client) (uint64, error) {
	var height uint64
	if client == nil {
		return height, errors.New("nil client")
	}
	latestBlock, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Errorf("BlockByNumber err:%v", err)
		return height, err
	}

	return latestBlock.NumberU64(), nil
}

func gasPrice(client *ethclient.Client) (*big.Int, error) {
	if client == nil {
		return nil, errors.New("nil client")
	}
	return client.SuggestGasPrice(context.Background())
}

func nonce(client *ethclient.Client, account string) (uint64, error) {
	var nonce uint64
	if client == nil {
		return nonce, errors.New("nil client")
	}
	return client.NonceAt(context.Background(), common.HexToAddress(account), nil)
}

func balance(client *ethclient.Client, account string) (*big.Int, error) {
	if client == nil {
		return nil, errors.New("nil client")
	}
	return client.BalanceAt(context.Background(), common.HexToAddress(account), nil)
}

func balanceOf(client *ethclient.Client, tokenAddress, account string) (*big.Int, error) {
	if client == nil {
		return nil, errors.New("nil client")
	}
	tokenInstance, err := factory.NewToken(common.HexToAddress(tokenAddress), client)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return tokenInstance.BalanceOf(nil, common.HexToAddress(account))
}

func decimals(client *ethclient.Client, tokenAddress string) (uint8, error) {
	var decimal uint8
	if client == nil {
		return decimal, errors.New("nil client")
	}

	tokenInstance, err := factory.NewToken(common.HexToAddress(tokenAddress), client)
	if err != nil {
		log.Error(err)
		return decimal, err
	}

	return tokenInstance.Decimals(nil)
}
