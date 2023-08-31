/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: eth_address.go
 * @Date: 2018/10/05
 * @Author: Jia Liu
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */
package eth

import (
	"encoding/hex"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func CheckEthAddr(ethAddrStr string) bool {
	return common.IsHexAddress(ethAddrStr)
}

func CheckEthPrivkeyAddr(privKey []byte, addrInput string) (bool, error) {
	key, err := hex.DecodeString(string(privKey))
	if err != nil {
		return false, err
	}

	ecdsaPri, err := crypto.ToECDSA(key)
	if err != nil {
		return false, err
	}

	etcAddr := crypto.PubkeyToAddress(ecdsaPri.PublicKey).String()
	if strings.Compare(etcAddr, addrInput) != 0 {
		return false, errors.New("address not matched with privkey")
	}
	return true, nil
}

func EthPrivKeyToPub(privKey []byte) ([]byte, error) {
	key, err := hex.DecodeString(string(privKey))
	if err != nil {
		return nil, err
	}

	ecdsaPri, err := crypto.ToECDSA(key)
	if err != nil {
		return nil, err
	}

	etcPub := crypto.FromECDSAPub(&ecdsaPri.PublicKey)
	return etcPub, nil
}
