/*
 * *******************************************************************
 * @项目名称: controller
 * @文件名称: address.go
 * @Date: 2020/02/12
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package controller

import (
	"errors"
	"fmt"
	"strings"

	"git.bhex.io/bhpc/wallet/chainnode/baasnode/common"
)

func createAddress(chain string) (string, error) {
	log.Info("createAddress chain:%v", chain)

	var address string
	var err error

	unUsedAddressKey := fmt.Sprintf(common.UnusedAddressSet, chain)
	usedAddressKey := fmt.Sprintf(common.UsedAddressSet, chain)

	for {
		listLen := common.Redis.Client.SCard(unUsedAddressKey).Val()
		if listLen == 0 {
			return "", errors.New(chain + " empty address")
		}

		// get from unused address set
		address = common.Redis.Client.SPop(unUsedAddressKey).Val()
		if strings.Contains(address, "baas_") {
			continue
		}

		// check in usedaddress set
		exist, err := common.Redis.IsSetMember(usedAddressKey, address)
		if err != nil {
			log.Errorf("createAddress check address err:%v", err)
			return "", err
		}

		if !exist && address != "" {
			break
		}
	}

	// set new address to set
	err = common.Redis.AddToSet(usedAddressKey, address)
	if err != nil {
		log.Errorf("createAddress add address err:%v", err)
		return "", err
	}

	log.Infof("get chain %v, address %v", chain, address)
	return address, nil
}
