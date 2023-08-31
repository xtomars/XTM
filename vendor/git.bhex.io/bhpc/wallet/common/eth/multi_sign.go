/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: multi_sign.go
 * @Date: 2018/06/04
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package eth

import (
	"encoding/hex"
	"errors"
	"math/big"

	"git.bhex.io/bhpc/wallet/common/eth/factory"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// MultiSign class
type MultiSign struct {
	Client *ethclient.Client
}

// NewMultiSign return a MultiSign instance
func NewMultiSign(RPCUrl string) (MultilSigner, error) {
	if RPCUrl == "" {
		return &MultiSign{}, nil
	}

	// create a MultiSign with rpc connect
	client, err := ethclient.Dial(RPCUrl)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &MultiSign{
		Client: client,
	}, nil
}

// CreateMutilSigWallet create a mutil sign wallet
func (multi *MultiSign) CreateMutilSigWallet(deployPrikey string, owners []string, required int, nonce, gasPrice uint64) (string, *types.Transaction, error) {
	var walletAddress string

	if gasPrice == 0 {
		return "", nil, errors.New("GasPrice can not be zero")
	}

	// check owners
	if !checkOwner(owners) {
		return walletAddress, nil, errors.New("Invalid owners")
	}

	// convert private string to ecdsa privatekey
	ecdsaPri, err := crypto.HexToECDSA(deployPrikey)
	if err != nil {
		log.Errorf("Parse private key err:%v", err)
		return walletAddress, nil, err
	}

	if len(owners) < required {
		return walletAddress, nil, errors.New("Required large than owners number")
	}

	// crate wallet transaction
	var walletOwners []common.Address
	for _, owner := range owners {
		walletOwners = append(walletOwners, common.HexToAddress(owner))
	}

	opts := bind.NewKeyedTransactor(ecdsaPri)
	opts.GasLimit = DefaultGasLimit
	opts.Nonce = big.NewInt(int64(nonce))
	opts.GasPrice = big.NewInt(int64(gasPrice))

	walletAddr, signedTx, _, err := factory.DeployMultiSigWallet(opts, multi.Client, walletOwners, big.NewInt(int64(required)))
	if err != nil {
		log.Errorf("Deploy mutil sign wallet err:%v", err)
		return walletAddress, nil, err
	}
	walletAddress = walletAddr.String()
	log.Debugf("Mutil sign address: %v", walletAddress)

	return walletAddress, signedTx, nil
}

// SubmitTransaction create a eth or erc20 transaction and sign it
func (multi *MultiSign) SubmitTransaction(walletAddress, tokenAddress, toAddress, fromPriKey string, amount string, nonce, gasPrice, decimals uint64) (*types.Transaction, error) {
	if gasPrice == 0 {
		return nil, errors.New("Nonce can not be zero")
	}

	// check wallet address
	if !common.IsHexAddress(walletAddress) {
		log.Error("Wallet address invalid")
		return nil, errors.New("Wallet address invalid")
	}

	// check to address
	if !common.IsHexAddress(toAddress) {
		log.Error("To address invalid")
		return nil, errors.New("To address invalid")
	}

	// check token and decimals
	if common.IsHexAddress(toAddress) && decimals == 0 {
		return nil, errors.New("Decimals can not be zero")
	}

	// build wallet instance
	walletInstance, opts, err := buildWalletInstance(multi.Client, walletAddress, fromPriKey, nonce, gasPrice)
	if err != nil {
		log.Errorf("buildWalletInstance err:%v\n", err)
		return nil, err
	}

	switch tokenAddress {
	case "":
		// ETH transfer
		return multiSignTransferEth(multi.Client, walletInstance, opts, toAddress, amount)
	default:
		// ERC20 transfer
	}

	return multiSignTransferErc20(multi.Client, walletInstance, opts, tokenAddress, toAddress, amount, decimals)
}

// ConfirmTransaction confirm the transaction with txID
func (multi *MultiSign) ConfirmTransaction(walletAddress, confirmerPriKey string, txID, nonce, gasPrice uint64) (*types.Transaction, error) {
	if gasPrice == 0 {
		return nil, errors.New("Nonce can not be zero")
	}

	walletInstance, opts, err := buildWalletInstance(multi.Client, walletAddress, confirmerPriKey, nonce, gasPrice)
	if err != nil {
		log.Errorf("buildWalletInstance err:%v\n", err)
		return nil, err
	}

	// build confirm ransaction
	signedTx, err := walletInstance.ConfirmTransaction(opts, big.NewInt(int64(txID)))
	if err != nil {
		log.Errorf("ConfirmTransaction err:%v\n", err)
		return nil, err
	}
	return signedTx, nil
}

// RevokeConfirmation revoke a transaction according to txID
func (multi *MultiSign) RevokeConfirmation(walletAddress, revokerPriKey string, txID, nonce, gasPrice uint64) (*types.Transaction, error) {
	if gasPrice == 0 {
		return nil, errors.New("Nonce can not be zero")
	}

	walletInstance, opts, err := buildWalletInstance(multi.Client, walletAddress, revokerPriKey, nonce, gasPrice)
	if err != nil {
		log.Errorf("buildWalletInstance err:%v\n", err)
		return nil, err
	}

	// build revoke ransaction
	signedTx, err := walletInstance.RevokeConfirmation(opts, big.NewInt(int64(txID)))
	if err != nil {
		log.Errorf("RevokeConfirmation err:%v\n", err)
		return nil, err
	}
	return signedTx, nil
}

// BroadcastTransaction broadcast the transaction to network
func (multi *MultiSign) BroadcastTransaction(tx *types.Transaction) (txHash string, costFee uint64, err error) {
	return broadcastTransaction(multi.Client, tx)
}

// Balance return the account balance
func (multi *MultiSign) Balance(account string) (*big.Int, error) {
	return nil, nil
}

// Nonce return the account nonce
func (multi *MultiSign) Nonce(account string) (uint64, error) {
	return nonce(multi.Client, account)
}

// GasPrice return the current gas price
func (multi *MultiSign) GasPrice() (*big.Int, error) {
	return gasPrice(multi.Client)
}

// BlockHeight return the current block height
func (multi *MultiSign) BlockHeight() (uint64, error) {
	return blockHeight(multi.Client)
}

// Decimals return the contract decimals
func (multi *MultiSign) Decimals(tokenAddress string) (uint8, error) {
	return decimals(multi.Client, tokenAddress)
}

// BalanceOf return the account balance of contract
func (multi *MultiSign) BalanceOf(tokenAddress, account string) (*big.Int, error) {
	return balanceOf(multi.Client, tokenAddress, account)
}

func multiSignTransferEth(client *ethclient.Client, wallet *factory.MultiSigWallet, opts *bind.TransactOpts, toAddress string, amount string) (*types.Transaction, error) {
	// calc amount
	weiAmount, err := ToWei(amount, DefaultDecimals)
	log.Debugf("Wei amount=%v", weiAmount)
	if err != nil {
		log.Errorf("ToWei err:%v, amount:%v, decimal:%v", err, amount, decimals)
		return nil, err
	}

	signedTx, err := wallet.SubmitTransaction(opts, common.HexToAddress(toAddress), weiAmount, nil)
	if err != nil {
		log.Errorf("SubmitTransaction err:%v\n", err)
		return nil, err
	}
	return signedTx, nil
}

func multiSignTransferErc20(client *ethclient.Client, wallet *factory.MultiSigWallet, opts *bind.TransactOpts, tokenAddress, toAddress string, amount string, decimals uint64) (*types.Transaction, error) {
	// get token instance
	tokenInstance, err := factory.NewToken(common.HexToAddress(tokenAddress), client)
	if err != nil {
		log.Errorf("NewToken err:%v\n", err)
		return nil, err
	}

	// calc amount
	weiAmount, err := ToWei(amount, int(decimals))
	log.Debugf("Wei amount=%v", weiAmount)
	if err != nil {
		log.Errorf("ToWei err:%v, amount:%v, decimal:%v", err, amount, decimals)
		return nil, err
	}

	// get token transfer data
	tokenTx, err := tokenInstance.Transfer(opts, common.HexToAddress(toAddress), weiAmount)
	if err != nil {
		log.Errorf("Get token tx data err:%v\n", err)
		return nil, err
	}
	log.Debugf("Transfer tx data: %v\n", hex.EncodeToString(tokenTx.Data()))

	// build submit transaction
	signedTx, err := wallet.SubmitTransaction(opts, common.HexToAddress(tokenAddress), big.NewInt(0), tokenTx.Data())
	if err != nil {
		log.Errorf("SubmitTransaction err:%v\n", err)
		return nil, err
	}

	return signedTx, nil
}

func checkOwner(owners []string) bool {
	for _, owner := range owners {
		if !common.IsHexAddress(owner) {
			return false
		}
	}
	return true
}
