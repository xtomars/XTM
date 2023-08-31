/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: env.go
 * @Date 2018/05/16
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package env

import (
	"os"
)

const (
	envKey  = "BHEX_WALLET_ENV"
	envTest = "BhexWalletTestEnv"
	envProd = "BhexWalletProdEnv"
)

var actualEnv string

func init() {
	actualEnv = os.Getenv(envKey)
}

// IsTestEnv return the env is test or not
func IsTestEnv() bool {
	return actualEnv == envTest
}

// IsProdEnv return the env is test or not
func IsProdEnv() bool {
	return actualEnv == envProd
}

// SetTestEnv set test env
func SetTestEnv() {
	actualEnv = envTest
}
