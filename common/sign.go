/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: sign.go
 * @Date: 2020/02/12
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package common

import (
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	"golang.org/x/crypto/ed25519"
)

// CreateSignMsg create the common sign msg
func CreateSignMsg(method, url, timestamp string, mapBody map[string]interface{}) string {
	result := strings.Join([]string{method, url, timestamp}, "|")

	if mapBody != nil {
		var keys []string

		for key := range mapBody {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		bodyParam := "|"
		for index, key := range keys {
			if strings.Contains(key, "memo") {
				if key != "memo" {
					log.Errorf("CreateSignMsg invalid memo key:%v", key)
					continue
				}

				if mapBody[key].(string) == "" {
					log.Errorf("CreateSignMsg had memo field, but value is nil")
					continue
				}
			}

			if index == len(keys)-1 {
				bodyParam = bodyParam + fmt.Sprintf("%s=%v", key, mapBody[key])
			} else {
				bodyParam = bodyParam + fmt.Sprintf("%s=%v&", key, mapBody[key])
			}
		}

		result = result + bodyParam
		log.Infof("makeParams result:%v", result)
		return result
	}

	log.Infof("makeParams result:%v", result)
	return result
}

// Sign return ed25519 sign result
func Sign(pri ed25519.PrivateKey, msg []byte) []byte {
	return ed25519.Sign(pri, msg)
}

// VerifySign  verify the ed25519 signature
func VerifySign(apiKey, message, signature string) bool {
	srcSign, err := hex.DecodeString(signature)
	if err != nil {
		log.Errorf("VerifySign err:%v", err)
		return false
	}

	pubKey, ok := globalAPIInfos.Load(apiKey)
	if !ok {
		log.Errorf("not found pub, api_Key:%v", apiKey)
	}

	pub, err := hex.DecodeString(pubKey.(string))
	if err != nil {
		log.Errorf("VerifySign err:%v", err)
		return false
	}

	return ed25519.Verify(pub, []byte(message), srcSign)
}
