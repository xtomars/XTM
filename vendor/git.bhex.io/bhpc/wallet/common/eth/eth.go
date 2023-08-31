/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: eth.go
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

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	ChainIDMain = 1
	ChainIDTest = 3
)

// Eth class
type Eth struct {
	Client  *ethclient.Client
	ChainID int64
}

// NewEth return a Eth instance
func NewEth(RPCUrl string, chainID int64) (Ether, error) {
	if RPCUrl == "" {
		return &Eth{
			ChainID: chainID,
		}, nil
	}

	// create a Eth with rpc connect
	client, err := ethclient.Dial(RPCUrl)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &Eth{
		Client:  client,
		ChainID: chainID,
	}, nil
}

// CreateAndSignTransaction create a eth transaction and sign it
func (eth *Eth) CreateAndSignTransaction(toAddress, fromPriKey string, amount string, nonce, gasPrice, gasLimit uint64) (*types.Transaction, error) {
	if gasPrice == 0 {
		return nil, errors.New("GasPrice can not be zero")
	}

	// check to address
	if !common.IsHexAddress(toAddress) {
		log.Error("To address invalid")
		return nil, errors.New("To address invalid")
	}

	// calc amount
	weiAmount, err := ToWei(amount, DefaultDecimals)
	log.Debugf("Wei amount=%v", weiAmount)
	if err != nil {
		log.Errorf("ToWei err:%v, amount:%v, decimal:%v", err, amount, DefaultDecimals)
		return nil, err
	}

	// make a transaction
	tx := types.NewTransaction(nonce, common.HexToAddress(toAddress), weiAmount, gasLimit, big.NewInt(int64(gasPrice)), nil)
	log.Debugf("send transaction: 0x%x\n", tx.Hash())

	// convert private string to ecdsa privatekey
	ecdsaPri, err := crypto.HexToECDSA(fromPriKey)
	if err != nil {
		log.Errorf("Parse private key err:%v", err)
		return nil, err
	}

	// sign transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(eth.ChainID)), ecdsaPri)
	if err != nil {
		log.Errorf("Sign tx err:%v", err)
		return nil, err
	}

	return signedTx, nil
}

// BroadcastTransaction broadcast the transaction to network
func (eth *Eth) BroadcastTransaction(tx *types.Transaction) (txHash string, gasUsed uint64, err error) {
	return broadcastTransaction(eth.Client, tx)
}

// Balance return the account balance
func (eth *Eth) Balance(account string) (*big.Int, error) {
	return balance(eth.Client, account)
}

// Nonce return the account nonce
func (eth *Eth) Nonce(account string) (uint64, error) {
	return nonce(eth.Client, account)
}

// GasPrice return the current gas price
func (eth *Eth) GasPrice() (*big.Int, error) {
	return gasPrice(eth.Client)
}

// BlockHeight return the current block height
func (eth *Eth) BlockHeight() (uint64, error) {
	return blockHeight(eth.Client)
}
