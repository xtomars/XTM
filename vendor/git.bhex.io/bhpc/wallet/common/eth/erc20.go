/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: erc20.go
 * @Date: 2018/06/04
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package eth

import (
	"errors"
	"math/big"

	"git.bhex.io/bhpc/wallet/common/eth/factory"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Erc20 class
type Erc20 struct {
	Client  *ethclient.Client
	ChainID int64
}

// NewErc20 return a Erc20 instance
func NewErc20(RPCUrl string, chainID int64) (Erc20er, error) {
	if RPCUrl == "" {
		return &Erc20{
			ChainID: chainID,
		}, nil
	}
	// create a Erc20 with rpc connect
	client, err := ethclient.Dial(RPCUrl)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &Erc20{
		Client:  client,
		ChainID: chainID,
	}, nil
}

// CreateAndSignTransaction create a erc20 transaction and sign it
func (erc *Erc20) CreateAndSignTransaction(tokenAddress, toAddress, fromPriKey string, amount string, nonce, gasPrice, decimals, gasLimit uint64) (*types.Transaction, error) {
	if tokenAddress == "" {
		return nil, errors.New("Token address can not be nil")
	}

	if gasPrice == 0 {
		return nil, errors.New("GasPrice can not be zero")
	}

	if decimals == 0 {
		return nil, errors.New("Decimals can not be zero")
	}

	// get token instance
	tokenInstance, err := factory.NewToken(common.HexToAddress(tokenAddress), erc.Client)
	if err != nil {
		log.Error("Token address invalid")
		return nil, errors.New("Token address invalid")
	}

	// check to address
	if !common.IsHexAddress(toAddress) {
		log.Error("To address invalid")
		return nil, errors.New("To address invalid")
	}

	// convert to wei amount
	weiAmount, err := ToWei(amount, int(decimals))
	log.Debugf("Wei amount=%v\n", weiAmount)
	if err != nil {
		log.Errorf("ToWei err:%v, amount:%v, decimal:%v", err, amount, decimals)
		return nil, err
	}

	// convert private string to ecdsa privatekey
	ecdsaPri, err := crypto.HexToECDSA(fromPriKey)
	if err != nil {
		log.Errorf("Parse private key err:%v", err)
		return nil, err
	}

	// sign tx
	opts := bind.NewKeyedTransactor(ecdsaPri)
	opts.GasLimit = gasLimit
	opts.GasPrice = big.NewInt(int64(gasPrice))
	opts.Nonce = big.NewInt(int64(nonce))

	signedTx, err := tokenInstance.Transfer(opts, common.HexToAddress(toAddress), weiAmount)
	if err != nil {
		log.Errorf("Token: %v, to address: %s, amount: %v, transfer failed: %v", tokenAddress, toAddress, amount, err)
		return nil, err
	}

	return signedTx, nil
}

// BroadcastTransaction broadcast the transaction to network
func (erc *Erc20) BroadcastTransaction(tx *types.Transaction) (txHash string, costFee uint64, err error) {
	return broadcastTransaction(erc.Client, tx)
}

// Balance return the account balance
func (erc *Erc20) Balance(account string) (*big.Int, error) {
	return nil, nil
}

// Nonce return the account nonce
func (erc *Erc20) Nonce(account string) (uint64, error) {
	return nonce(erc.Client, account)
}

// GasPrice return the current gas price
func (erc *Erc20) GasPrice() (*big.Int, error) {
	return gasPrice(erc.Client)
}

// BlockHeight return the current block height
func (erc *Erc20) BlockHeight() (uint64, error) {
	return blockHeight(erc.Client)
}

// Decimals return the contract decimals
func (erc *Erc20) Decimals(tokenAddress string) (uint8, error) {
	return decimals(erc.Client, tokenAddress)
}

// BalanceOf return the account balance of contract
func (erc *Erc20) BalanceOf(tokenAddress, account string) (*big.Int, error) {
	return balanceOf(erc.Client, tokenAddress, account)
}
