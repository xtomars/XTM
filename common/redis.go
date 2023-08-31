/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: redis.go
 * @Date: 2020/02/11
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package common

import (
	"fmt"

	"git.bhex.io/bhpc/wallet/chainnode/baasnode/config"
	"git.bhex.io/bhpc/wallet/chainnode/util/redis"
	"git.bhex.io/bhpc/wallet/common/logger"
)

var (
	// Redis global
	Redis *redis.Cluster

	log = logger.New("common")

	// reids keys

	// UnusedAddressList unused addr list key
	UnusedAddressList = "baas_%s_unused_address_list"

	// UnusedAddressSet unused addr set key
	UnusedAddressSet = "baas_%s_unused_address_set"

	// UsedAddressSet used addr set key
	UsedAddressSet = "baas_%s_used_address_set"

	// notifyFailedLock global notify failed lock
	notifyFailedLock = "baas_notify_failed_lock"
	// WithdrawalOrderList store unhandled withdrawal order
	// WithdrawalOrderList = "baas_%s_withdrawal_order_list"
)

func init() {
	var err error

	redisURL := fmt.Sprintf("%s:%s",
		config.Config.Redis.Host,
		config.Config.Redis.Port)

	// redis
	Redis, err = redis.NewCluster(
		redisURL,
		config.Config.Redis.Pass,
		config.Config.Redis.MaxConn)

	if err != nil {
		log.Error("[ALERT] Connect redis err:", err)
		panic(err)
	}
}
