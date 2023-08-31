/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: interface.go
 * @Date: 2018/06/04
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// Ether define the eth transfer interface
type Ether interface {
	CreateAndSignTransaction(toAddress, fromPriKey string, amount string, nonce, gasPrice, gasLimit uint64) (*types.Transaction, error)
	Queryer
	Broadcaster
}

// Erc20er define the erc20 transfer interface
type Erc20er interface {
	CreateAndSignTransaction(tokenAddress, toAddress, fromPriKey string, amount string, nonce, gasPrice, decimals, gasLimit uint64) (*types.Transaction, error)
	ContractQueryer
	Broadcaster
}

// MultilSigner define the MultilSign transfer interface
type MultilSigner interface {
	CreateMutilSigWallet(deployPrikey string, owners []string, required int, nonce, gasPrice uint64) (string, *types.Transaction, error)
	SubmitTransaction(walletAddress, tokenAddress, toAddress, fromPriKey string, amount string, nonce, gasPrice, decimals uint64) (*types.Transaction, error)
	ConfirmTransaction(walletAddress, confirmerPriKey string, txID, nonce, gasPrice uint64) (*types.Transaction, error)
	RevokeConfirmation(walletAddress, revokerPriKey string, txID, nonce, gasPrice uint64) (*types.Transaction, error)
	ContractQueryer
	Broadcaster
}

// Broadcaster define the broadcast transaction interface
type Broadcaster interface {
	BroadcastTransaction(tx *types.Transaction) (txHash string, costFee uint64, err error)
}

// Queryer define the query interface
type Queryer interface {
	Balance(account string) (*big.Int, error)
	Nonce(account string) (uint64, error)
	GasPrice() (*big.Int, error)
	BlockHeight() (uint64, error)
}

// ContractQueryer define the contract query interface
type ContractQueryer interface {
	Decimals(tokenAddress string) (uint8, error)
	BalanceOf(tokenAddress, account string) (*big.Int, error)
	Queryer
}
