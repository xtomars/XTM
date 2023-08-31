/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: retry.go
 * @Date: 2020/05/09
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package common

import "time"

// FuncWithRetry wrap func with retry
func FuncWithRetry(retryCount int, fn func() error) error {
	var err error
	for i := 0; i < retryCount; i++ {
		if err = fn(); err == nil {
			return nil
		}

		time.Sleep(200 * time.Millisecond)
	}

	return err
}
