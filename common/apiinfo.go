/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: apiinfo.go
 * @Date: 2020/02/11
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package common

import (
	"sync"
)

// globalAPIInfos storage
var globalAPIInfos sync.Map

// SaveAPIInfo save info
func SaveAPIInfo(apiKey, pubKey string) {
	globalAPIInfos.Store(apiKey, pubKey)
}

// ValidAPIKey valid the apikey
func ValidAPIKey(apiKey string) bool {
	_, ok := globalAPIInfos.Load(apiKey)
	return ok
}
