/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: tokeninfo.go
 * @Date: 2020/02/12
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package common

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"git.bhex.io/bhpc/wallet/chainnode/baasnode/config"
	"git.bhex.io/bhpc/wallet/chainnode/util/walletclient"
	"git.bhex.io/bhpc/wallet/common/models"
	"git.bhex.io/bhpc/wallet/common/protos/common"
	"git.bhex.io/bhpc/wallet/common/protos/wallet"
)

var (
	walletClient wallet.WalletClient

	// globalTokenInfos storage
	globalTokenInfos sync.Map

	globalChainInfo sync.Map
)

// TokenInfo define
type TokenInfo struct {
	TokenID          string `json:"token_id"`
	APIKey           string `json:"api_key"`
	Chain            string `json:"chain"`
	Decimals         int    `json:"decimals"`
	EnableDeposit    bool   `json:"enable_deposit"`
	EnableWithdrawal bool   `json:"enable_withdrawal"`
	NeedMemo         bool   `json:"need_memo"`
	AddressRegexp    string `json:"address_regexp"`
}

func init() {
	var err error
	// wallet client
	walletClient, err = walletclient.New(config.Config.Wallet.Host, config.Config.Wallet.Port)
	if err != nil {
		log.Error("Connect wallet server err:", err)
		panic(err)
	}
}

// SyncAllTokens async
func SyncAllTokens() {
	loadTokens()

	ticker := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-ticker.C:
			loadTokens()
		}
	}
}

// CheckToken check token
func CheckToken(tokenID string) bool {
	_, exist := globalTokenInfos.Load(tokenID)
	return exist
}

// CheckBlockTime format
func CheckBlockTime(blockTime string) bool {
	now := time.Now().Unix()
	if len(fmt.Sprintf("%v", now)) != len(blockTime) {
		log.Infof("block time:%v format error", blockTime)
		return false
	}
	return true
}

// GetTokenInfoByTokenID get token info
func GetTokenInfoByTokenID(tokenID string) *TokenInfo {
	info, exist := globalTokenInfos.Load(tokenID)
	if !exist {
		return nil
	}

	return info.(*TokenInfo)
}

// GetTokenInfoByChain get token info
func GetTokenInfoByChain(chain string) *TokenInfo {
	tokenInfo := new(TokenInfo)

	var exist bool

	globalTokenInfos.Range(func(k, v interface{}) bool {
		tokenInfo = v.(*TokenInfo)
		if tokenInfo.Chain == chain {
			exist = true
			return false
		}
		return true
	})

	if !exist {
		return nil
	}

	return tokenInfo
}

func loadTokens() {
	log.Infof("loadTokens start")

	// 后期币种增多，需要做分页查询
	reply, err := walletClient.QueryWalletDB(context.Background(), &wallet.QueryWalletDBRequest{
		Method: "GetBaasTokens",
	})

	if err != nil {
		log.Warn("Wallet service not avaliable")
		return
	}

	if reply.Code != common.ReturnCode_SUCCESS {
		log.Errorf("loadTokens err:%v", reply.Msg)
		return
	}

	var result []models.Token
	err = json.Unmarshal(reply.ResData, &result)
	if err != nil {
		log.Errorf("loadTokens json Unmarshal err:%v", err)
		return
	}

	 //log.Infof("loadTokens get baas %v tokens: %v", len(result), result)

	for _, info := range result {
		exts := parseExtension(info.Extension)
		chain, apiKey, pubKey, addressRegexp := exts["chain"].(string), exts["api_key"].(string), exts["pub_key"].(string), exts["address_regexp"].(string)

		if chain == "" || apiKey == "" || pubKey == "" {
			continue
		}

		globalTokenInfos.Store(string(info.TokenID), &TokenInfo{
			TokenID:          string(info.TokenID),
			Chain:            chain,
			APIKey:           apiKey,
			EnableDeposit:    info.EnableNotify,
			EnableWithdrawal: info.EnableWithdrawal,
			NeedMemo:         info.NeedMemo,
			Decimals:         int(info.Decimals),
			AddressRegexp:    addressRegexp,
		})
		globalAPIInfos.Store(apiKey, pubKey)
	}
}

func parseExtension(ext string) map[string]interface{} {
	result := make(map[string]interface{})
	json.Unmarshal([]byte(ext), &result)
	return result
}
